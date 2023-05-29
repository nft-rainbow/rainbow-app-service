package services

import (
	"math"
	"sort"
	"strings"
	"sync"

	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"strconv"
	"time"

	. "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	. "github.com/nft-rainbow/rainbow-app-service/config"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	randutils "github.com/nft-rainbow/rainbow-app-service/utils/rand"

	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	activityService     *ActivityService
	activityServiceOnce sync.Once
)

type MintReq struct {
	ActivityID  string `json:"activity_id" binding:"required"`
	UserAddress string `json:"user_address" binding:"required"`
	Command     string `json:"command"`
}

type ActivityService struct {
}

func GetActivityService() *ActivityService {
	if activityService != nil {
		activityServiceOnce.Do(func() {
			activityService = &ActivityService{}
		})
	}
	return activityService
}

func (a *ActivityService) InsertActivity(activityReq *models.ActivityInsertPart, userId uint) (*models.Activity, error) {
	if err := activityReq.SetDefaults(); err != nil {
		return nil, err
	}

	activityId := utils.GenerateIDByTimeHash("", 8)
	posterUrl, err := generateActivityPoster(&activityReq.ActivityUpdateBasePart, activityId)
	if err != nil {
		logrus.Errorf("Failed to generate poster for activity %v:%v \n", activityId, err.Error())
		return nil, errors.WithStack(err)
	}

	activity := models.Activity{
		ActivityInsertPart: *activityReq,
		RainbowUserId:      userId,
		ActivityCode:       activityId,
		ActivityPosterURL:  posterUrl,
	}

	if activityReq.ActivityType == enums.ACTIVITY_GASLESS {
		gasLess := GetConfig().Gasless

		if activityReq.Amount > int32(gasLess.MaxAmount) {
			return nil, errors.Errorf("exceed gasless max amount: %v", gasLess.MaxAmount)
		}

		contractRawId := gasLess.ContractRawID.Testnet
		if activityReq.ChainOfGasless == enums.CHAIN_CONFLUX {
			contractRawId = gasLess.ContractRawID.Mainnet
		}
		_contractRawId := int32(contractRawId)
		activity.ContractRawID = &_contractRawId

		if err := a.UpdateOrCreateContract(gasLess.UserID, gasLess.AppID, contractRawId); err != nil {
			return nil, errors.WithStack(err)
		}
	} else if activityReq.ContractRawID != nil {
		if err := a.UpdateOrCreateContract(userId, activityReq.AppId, uint(*activityReq.ContractRawID)); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	res := models.GetDB().Create(&activity)
	if res.Error != nil {
		return nil, res.Error
	}

	if err := activity.LoadBindedContract(); err != nil {
		return nil, errors.WithStack(err)
	}
	return &activity, nil
}

func (a *ActivityService) AddActivityNftConfigs(activityCode string, nftConfigUpdateParts []*models.NftConfigUpdatePart, userId uint) ([]*models.NFTConfig, error) {
	activity, err := models.FindActivityByCode(activityCode)
	if err != nil {
		return nil, err
	}

	if activity.RainbowUserId != userId {
		return nil, errors.New("no permission")
	}

	var nftConfigs []*models.NFTConfig
	for _, p := range nftConfigUpdateParts {
		nftConfigs = append(nftConfigs, &models.NFTConfig{
			NftConfigUpdatePart: *p,
			ActivityID:          activity.ID,
		})
	}

	if err := models.GetDB().Save(&nftConfigs).Error; err != nil {
		return nil, err
	}

	return nftConfigs, nil
}

func (a *ActivityService) POAPH5Config(config *models.H5Config) (*models.H5Config, error) {
	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return nil, res.Error
	}

	return config, nil
}

func (a *ActivityService) UpdateOrCreateContract(userId uint, appId uint, contractId uint) error {
	token, err := middlewares.GenerateRainbowOpenJWT(userId, appId)
	if err != nil {
		return err
	}

	info, err := utils.GetContractInfo(int32(contractId), token)
	if err != nil {
		return err
	}

	if uint(*info.AppId) != appId {
		return errors.New("contract not belongs to app")
	}

	if info.Address == nil || *info.Address == "" {
		return errors.New("contract not deployed")
	}

	_, err = models.UpdateOrCreateContract(uint(contractId), uint(*info.Type), uint(*info.ChainId), uint(*info.ChainType), *info.Address)
	if err != nil {
		return err
	}
	return nil
}

func (a *ActivityService) UpdateActivityBase(userId uint, activityCode string, req *models.ActivityUpdateBasePart) (*models.Activity, error) {
	activity, err := models.FindActivityByCode(activityCode)
	if err != nil {
		return nil, err
	}

	if activity.RainbowUserId != userId {
		return nil, errors.New("no permission")
	}

	if req.ContractRawID != nil {
		if err := a.UpdateOrCreateContract(activity.RainbowUserId, activity.AppId, uint(*req.ContractRawID)); err != nil {
			return nil, err
		}
		activity.ContractRawID = req.ContractRawID
	}

	// for _, nftConfig := range req.NFTConfigs {
	// 	nftConfig.ActivityID = activity.ID
	// }

	req.SetDefaults()
	activity.ActivityUpdateBasePart = *req
	if err := models.GetDB().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&activity).Error; err != nil {
		return nil, err
	}

	if err := activity.LoadBindedContract(); err != nil {
		return nil, err
	}
	return activity, nil
}

func (a *ActivityService) UpdateNftConfig(userId uint, nftConfigId uint, updateNftConfig *models.NftConfigUpdatePart) (*models.NFTConfig, error) {
	nftConfig, err := models.FindNftConfigById(nftConfigId)
	if err != nil {
		return nil, err
	}

	if err := models.CheckNftConfigBelongToUser(nftConfig, userId); err != nil {
		return nil, err
	}

	nftConfig.NftConfigUpdatePart = *updateNftConfig
	if err := models.GetDB().Save(nftConfig).Error; err != nil {
		return nil, err
	}

	return nftConfig, nil
}

func (a *ActivityService) DeleteActivityNftConfig(userId uint, nftConfigId uint) error {
	nftConfig, err := models.FindNftConfigById(nftConfigId)
	if err != nil {
		return err
	}

	if err := models.CheckNftConfigBelongToUser(nftConfig, userId); err != nil {
		return err
	}

	return models.GetDB().Delete(&nftConfig).Error
}

// func (a *ActivityService) UpdateActivity2(activityId string, req *models.UpdateActivityReq) (*models.Activity, error) {
// 	oldConfig, err := models.FindActivityByCode(activityId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if req.ContractRawID != nil {
// 		if oldConfig.ContractRawID == nil || uint(*req.ContractRawID) != uint(*oldConfig.ContractRawID) {
// 			if err := a.UpdateOrCreateContract(oldConfig.RainbowUserId, oldConfig.AppId, uint(*req.ContractRawID)); err != nil {
// 				return nil, err
// 			}
// 			oldConfig.ContractRawID = req.ContractRawID
// 		}
// 	}

// 	// Create a map of oldConfig.NFTConfigs for fast searching
// 	oldNFTConfigsMap := make(map[uint]*models.NFTConfig)
// 	newNFTConfigsMap := make(map[uint]*models.NFTConfig)

// 	for i, nftConfig := range oldConfig.NFTConfigs {
// 		oldNFTConfigsMap[nftConfig.ID] = &oldConfig.NFTConfigs[i]
// 	}

// 	for i, nftConfig := range req.NFTConfigs {
// 		if nftConfig.ID != 0 {
// 			newNFTConfigsMap[nftConfig.ID] = &req.NFTConfigs[i]
// 		}
// 	}

// 	// Update NFTConfigs
// 	for _, newNFTConfig := range req.NFTConfigs {
// 		if oldNFTConfig, ok := oldNFTConfigsMap[newNFTConfig.ID]; ok {
// 			// Update existing NFTConfig
// 			oldNFTConfig.Probability = newNFTConfig.Probability
// 			oldNFTConfig.Name = newNFTConfig.Name
// 			oldNFTConfig.ImageURL = newNFTConfig.ImageURL

// 			// Update MetadataAttributes
// 			oldMetadataAttributesMap := make(map[uint]*models.MetadataAttribute)
// 			newMetadataAttributesMap := make(map[uint]*models.MetadataAttribute)

// 			for j, metadataAttribute := range oldNFTConfig.MetadataAttributes {
// 				oldMetadataAttributesMap[metadataAttribute.ID] = oldNFTConfig.MetadataAttributes[j]
// 			}
// 			for j, metadataAttribute := range newNFTConfig.MetadataAttributes {
// 				if metadataAttribute.ID != 0 {
// 					newMetadataAttributesMap[metadataAttribute.ID] = newNFTConfig.MetadataAttributes[j]
// 				}
// 			}

// 			if len(newNFTConfig.MetadataAttributes) > 0 {
// 				for _, newMetadataAttribute := range newNFTConfig.MetadataAttributes {
// 					if oldMetadataAttribute, ok := oldMetadataAttributesMap[newMetadataAttribute.ID]; ok {
// 						// Update existing MetadataAttribute
// 						oldMetadataAttribute.TraitType = newMetadataAttribute.TraitType
// 						oldMetadataAttribute.Value = newMetadataAttribute.Value
// 						oldMetadataAttribute.DisplayType = newMetadataAttribute.DisplayType
// 						models.GetDB().Save(&oldMetadataAttribute)
// 					} else {
// 						// Create new MetadataAttribute
// 						newMetadataAttribute.NFTConfigID = newNFTConfig.ID
// 						oldNFTConfig.MetadataAttributes = append(oldNFTConfig.MetadataAttributes, newMetadataAttribute)
// 					}
// 				}
// 				for j := len(oldNFTConfig.MetadataAttributes) - 1; j >= 0; j-- {
// 					if oldNFTConfig.MetadataAttributes[j].ID == 0 {
// 						models.GetDB().Save(&oldNFTConfig.MetadataAttributes[j])
// 						continue
// 					}
// 					if _, ok := newMetadataAttributesMap[oldNFTConfig.MetadataAttributes[j].ID]; !ok {
// 						// Delete MetadataAttribute
// 						models.GetDB().Delete(&oldNFTConfig.MetadataAttributes[j])
// 						oldNFTConfig.MetadataAttributes = append(oldNFTConfig.MetadataAttributes[:j], oldNFTConfig.MetadataAttributes[j+1:]...)
// 					}
// 				}
// 			} else {
// 				for j, attribute := range oldNFTConfig.MetadataAttributes {
// 					models.GetDB().Delete(&attribute)
// 					oldNFTConfig.MetadataAttributes = append(oldNFTConfig.MetadataAttributes[:j], oldNFTConfig.MetadataAttributes[j+1:]...)
// 				}
// 			}
// 			models.GetDB().Save(&oldNFTConfig)
// 		} else {
// 			// Create new NFTConfig
// 			newNFTConfig.ActivityID = oldConfig.ID
// 			oldConfig.NFTConfigs = append(oldConfig.NFTConfigs, newNFTConfig)
// 		}
// 	}

// 	// Delete NFTConfigs
// 	for i := len(oldConfig.NFTConfigs) - 1; i >= 0; i-- {
// 		if oldConfig.NFTConfigs[i].ID == 0 {
// 			models.GetDB().Save(&oldConfig.NFTConfigs[i])
// 			continue
// 		}
// 		if _, ok := newNFTConfigsMap[oldConfig.NFTConfigs[i].ID]; !ok {
// 			// Delete NFTConfig
// 			models.GetDB().Delete(&oldConfig.NFTConfigs[i])
// 			oldConfig.NFTConfigs = append(oldConfig.NFTConfigs[:i], oldConfig.NFTConfigs[i+1:]...)
// 		}
// 	}

// 	// oldConfig.AppName = req.AppName
// 	oldConfig.MaxMintCount = req.MaxMintCount
// 	oldConfig.Command = req.Command
// 	oldConfig.StartedTime = req.StartedTime
// 	oldConfig.EndedTime = req.EndedTime
// 	oldConfig.Amount = req.Amount
// 	oldConfig.Name = req.Name
// 	oldConfig.Description = req.Description
// 	if len(req.WhiteListInfos) != 0 {
// 		models.GetDB().Delete(&oldConfig.WhiteListInfos)
// 	}
// 	oldConfig.WhiteListInfos = req.WhiteListInfos

// 	if oldConfig.ActivityPictureURL != req.ActivityPictureURL {
// 		oldConfig.ActivityPictureURL = req.ActivityPictureURL
// 		posterUrl, err := generateActivityPoster(req, activityId)
// 		if err != nil {
// 			logrus.Errorf("Failed to generate poster for activity %v:%v \n", activityId, err.Error())
// 			return nil, err
// 		}
// 		oldConfig.ActivityPosterURL = posterUrl
// 	}

// 	if err := models.GetDB().Save(&oldConfig).Error; err != nil {
// 		return nil, err
// 	}

// 	if err := oldConfig.LoadBindedContract(); err != nil {
// 		return nil, err
// 	}
// 	return oldConfig, nil
// }

func (a *ActivityService) H5Mint(req *MintReq) (*models.POAPResult, error) {
	activity, err := models.FindActivityByCode(req.ActivityID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = a.CheckMintable(activity, req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	token, err := getOpenApiToken(activity)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	nftConfig, err := calcNftConfig(activity)
	if err != nil {
		return nil, err
	}

	resp, err := mint(activity, nftConfig, req.UserAddress, token)
	if err != nil {
		return nil, err
	}

	models.GetMintCountCache(req.ActivityID).Increase()

	result, err := saveMintResult(activity, nftConfig, resp)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *ActivityService) getMintableCount(activity *models.Activity, address string) (int32, error) {

	if address == "" {
		return 0, nil
	}

	//TODO: is certi qulified

	var err error
	addrsOfPhone := []string{address}
	if activity.IsPhoneWhiteListOpened {
		addrsOfPhone, err = models.FindRelatedAddressWithSamePhone(address)
		if err != nil {
			return 0, err
		}
	}

	mintedCount, err := models.GetMintSumByAddresses(activity.ActivityCode, addrsOfPhone...)
	if err != nil {
		return 0, err
	}

	if activity.Amount == -1 {
		if activity.MaxMintCount == -1 {
			return -1, nil
		}
		return activity.MaxMintCount - int32(mintedCount), nil
	}

	totalRemain := int64(activity.Amount) - models.GetMintCountCache(activity.ActivityCode).GetCount()
	if totalRemain <= 0 {
		return 0, err
	}
	userRemain := math.Min(float64(totalRemain), float64(activity.Amount)-float64(mintedCount))
	return int32(userRemain), nil
}

func (a *ActivityService) GetMintCount(activityID, address string) (int32, error) {
	activity, err := models.FindActivityByCode(activityID)
	if err != nil {
		return 0, err
	}
	return a.getMintableCount(activity, address)

	// config, err := models.FindActivityByCode(activityID)
	// if err != nil {
	// 	return nil, err
	// }
	// var count int32

	// // phone white list logic: if whiteList config opened and user not in whiteList then the mint count is 0
	// if config.IsPhoneWhiteListOpened {
	// 	users, err := models.FindWalletUserByAddress(address)
	// 	if err == nil { // TODO check the phone not found case
	// 		var isInWhiteList bool
	// 		for _, u := range users {
	// 			isInWhiteList = models.IsPhoneInWhiteList(activityID, u.Phone)
	// 			if isInWhiteList {
	// 				break
	// 			}
	// 		}

	// 		if !isInWhiteList {
	// 			count = 0
	// 			return &count, nil
	// 		}
	// 	}
	// }

	// mintedCount, err := models.GetMintSumByAddresses(activityID, address)
	// if err != nil {
	// 	return nil, err
	// }

	// var remainedMinted int32
	// if config.MaxMintCount == -1 {
	// 	remainedMinted = -1
	// } else {
	// 	remainedMinted = int32(int64(config.MaxMintCount) - mintedCount)
	// }
	// logrus.WithField("remain", remainedMinted).Info("get remain mint count")

	// if config.Amount == -1 {
	// 	count = remainedMinted
	// } else {
	// 	if remainedMinted == -1 {
	// 		count = config.Amount - int32(models.GetMintCountCache(activityID).GetCount()) // Amount - total minted count
	// 	} else {
	// 		if config.Amount-int32(mintedCount) < remainedMinted {
	// 			count = config.Amount - int32(mintedCount)
	// 		} else {
	// 			count = remainedMinted
	// 		}
	// 	}
	// }
	// return &count, nil
}

func (a *ActivityService) CheckMintable(activity *models.Activity, req *MintReq) error {
	if err := activity.VerifyMintable(); err != nil {
		return errors.WithStack(err)
	}

	addrsOfPhone := []string{req.UserAddress}
	if activity.IsPhoneWhiteListOpened {
		users, err := models.FindWalletUserByAddress(req.UserAddress)

		if err != nil {
			return err
		}

		var isInWhiteList bool
		for _, u := range users {
			isInWhiteList = models.IsPhoneInWhiteList(req.ActivityID, u.Phone)
			if isInWhiteList {
				break
			}
		}

		if !isInWhiteList {
			return ERR_BUSINESS_NO_MINT_PERMISSIION
		}

		addrs, err := models.FindAllUserAddrsOfPhone(users[0].Phone)
		if err != nil {
			return err
		}
		addrsOfPhone = append(addrsOfPhone, addrs...)
	}

	if err := checkUserMintQuota(activity.ActivityCode, addrsOfPhone, activity.MaxMintCount); err != nil {
		return err
	}

	if req.Command != activity.Command {
		return ERR_BUSINESS_VISPER_WRONG
	}

	return nil
}

func calcNftConfig(activity *models.Activity) (*models.NFTConfig, error) {
	if len(activity.NFTConfigs) == 0 {
		return nil, ERR_BUSNISS_ACTIVITY_CONFIG_WRONG
	}

	var nftConfigIndex int
	switch activity.ActivityType {
	case enums.ACTIVITY_SINGLE:
	case enums.ACTIVITY_BLINDBOX:
		probabilities := make([]float32, 0)
		for i := 0; i < len(activity.NFTConfigs); i++ {
			probabilities = append(probabilities, activity.NFTConfigs[i].Probability)
		}

		nftConfigIndex = weightedRandomIndex(probabilities)
	}

	nftConfig := activity.NFTConfigs[nftConfigIndex]
	return &nftConfig, nil
}

func crateMetadata(activity *models.Activity, nftConfig *models.NFTConfig, token string) (string, error) {
	metadataUri := activity.MetadataUri
	if metadataUri != "" {
		return metadataUri, nil
	}

	return createMetadataByNftConfig(nftConfig, activity.Description, token)
}

func createMetadataByNftConfig(nftConfig *models.NFTConfig, description string, token string) (string, error) {
	attributes := make([]openapiclient.ModelsExposedMetadataAttribute, 0)
	for _, v := range nftConfig.MetadataAttributes {
		attributes = append(attributes, openapiclient.ModelsExposedMetadataAttribute{
			TraitType:   &v.TraitType,
			Value:       &v.Value,
			DisplayType: &v.DisplayType,
		})
	}

	now := strconv.FormatInt(time.Now().Unix(), 10)

	trait := "mint_time"
	display := "date"
	attributes = append(attributes, openapiclient.ModelsExposedMetadataAttribute{
		Value:       &now,
		TraitType:   &trait,
		DisplayType: &display,
	})

	resp, err := utils.SendCreateMetadataRequest(token, openapiclient.ServicesMetadataDto{
		Description: description,
		Image:       nftConfig.ImageURL,
		Name:        nftConfig.Name,
		Attributes:  attributes,
	})
	if err != nil {
		return "", err
	}
	return *resp.Uri, nil
}

func weightedRandomIndex(weights []float32) int {
	if len(weights) == 1 {
		return 0
	}
	var sum float32 = 0.0
	for _, w := range weights {
		sum += w
	}
	r := rand.Float32() * sum
	var t float32 = 0.0
	for i, w := range weights {
		t += w
		if t > r {
			return i
		}
	}
	return len(weights) - 1
}

func checkUserMintQuota(activityId string, userAddrs []string, max int32) error {
	if max == -1 {
		return nil
	}

	count, err := models.GetMintSumByAddresses(activityId, userAddrs...)
	if err != nil {
		return err
	}

	if int32(count) >= max {
		return ERR_BUSINESS_PERSONAL_MAX_AMOUNT_ARRIVED
	}
	return nil
}

func mint(activity *models.Activity, nftConfig *models.NFTConfig, to string, token string) (*openapiclient.ModelsMintTask, error) {
	metadataURI, err := crateMetadata(activity, nftConfig, token)
	if err != nil {
		return nil, err
	}

	chain, err := utils.ChainById(uint(activity.Contract.ChainId))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	nextTokenId := randutils.NumString(10)
	if activity.IsTokenIdOrdered != nil && *activity.IsTokenIdOrdered {
		ignoreTokenIds, err := models.GetActivityResrverTokenIds(activity.ID)
		if err != nil {
			return nil, err
		}
		profile, err := utils.GetContractProfile(activity.Contract.ContractAddress, ignoreTokenIds, token)
		if err != nil {
			return nil, err
		}
		_nextTokenId := calcNextTokenId(uint(*profile.MaxTokenId), ignoreTokenIds)
		nextTokenId = fmt.Sprintf("%d", _nextTokenId)
	}

	// create mint meta
	mintMeta := openapiclient.ServicesCustomMintDto{
		Chain:           chain,
		ContractAddress: activity.Contract.ContractAddress,
		MintToAddress:   to,
		MetadataUri:     utils.PtrString(strings.Replace(metadataURI, "{id}", nextTokenId, -1)),
		TokenId:         &nextTokenId,
	}

	resp, err := utils.SendCustomMintRequest(token, mintMeta)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return resp, nil
}

func calcNextTokenId(currentMax uint, ignoreTokenIds [][2]uint) uint {
	sortIgnoreTokenIds(ignoreTokenIds)
	nextTokenId := uint(currentMax + 1)

	fmt.Println(ignoreTokenIds)

	for _, r := range ignoreTokenIds {
		if nextTokenId >= r[0] && nextTokenId <= r[1] {
			nextTokenId = r[1] + 1
		}
	}
	return nextTokenId
}

func sortIgnoreTokenIds(ignoreTokenIds [][2]uint) {
	sort.SliceStable(ignoreTokenIds, func(i, j int) bool {
		return (ignoreTokenIds[i][0] <= ignoreTokenIds[j][0]) && (ignoreTokenIds[i][1] <= ignoreTokenIds[j][1])
	})
}

func saveMintResult(activity *models.Activity, nftConfig *models.NFTConfig, resp *openapiclient.ModelsMintTask) (*models.POAPResult, error) {
	item := &models.POAPResult{
		ConfigID:      int32(activity.ID),
		Address:       *resp.MintTo,
		ContractRawID: *activity.ContractRawID,
		TxID:          *resp.Id,
		TokenID:       *resp.TokenId,
		ActivityCode:  activity.ActivityCode,
		FileURL:       nftConfig.ImageURL,
		ProjectorId:   activity.RainbowUserId,
		AppId:         activity.AppId,
	}
	err := models.GetDB().Create(&item).Error
	return item, err
}

func getOpenApiToken(activity *models.Activity) (string, error) {
	userId, appId := activity.RainbowUserId, activity.AppId
	if activity.ActivityType == enums.ACTIVITY_GASLESS {
		gasless := GetConfig().Gasless
		userId, appId = gasless.UserID, gasless.AppID
	}
	token, err := middlewares.GenerateRainbowOpenJWT(userId, appId)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return token, err
}

// func getPOAPId(address string, name string) (string, error) {
// 	hash := sha256.New()

// 	_, err := hash.Write([]byte(address + name + strconv.FormatInt(time.Now().UnixNano(), 10)))
// 	if err != nil {
// 		return "", err
// 	}
// 	sum := hash.Sum(nil)

// 	pId := hex.EncodeToString(sum)
// 	return pId[:8], nil
// }

// func checkWhiteList(whiteList []models.WhiteListInfo, address string) bool {
// 	for _, v := range whiteList {
// 		if address == v.User {
// 			return true
// 		}
// 	}
// 	return false
// }

// func checkWhiteListLimit(config *models.Activity, address string) error {
// 	// if err := config.CheckActivityValid(); err != nil {
// 	// 	return err
// 	// }
// 	resp, err := models.CountPOAPResultByAddress(address, config.ActivityCode)
// 	if err != nil {
// 		return err
// 	}
// 	for _, v := range config.WhiteListInfos {
// 		if v.User == address && resp >= int64(v.Count) {
// 			return fmt.Errorf("The NFT minted by the account has exceeded the mint limit")
// 		}
// 	}
// 	return nil
// }
