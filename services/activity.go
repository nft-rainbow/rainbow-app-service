package services

import (
	"sync"

	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"strconv"
	"time"

	"github.com/mcuadros/go-defaults"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils"
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

func (a *ActivityService) InsertActivity(activityReq *models.ActivityReq, userId uint) (*models.Activity, error) {
	defaults.SetDefaults(activityReq)

	activityId := utils.GenerateIDByTimeHash("", 8)
	posterUrl, err := generateActivityPoster(&activityReq.UpdateActivityReq, activityId)
	if err != nil {
		logrus.Errorf("Failed to generate poster for activity %v:%v \n", activityId, err.Error())
		return nil, err
	}

	activity := models.Activity{
		ActivityReq:       *activityReq,
		RainbowUserId:     userId,
		ActivityCode:      activityId,
		ActivityPosterURL: posterUrl,
	}

	if activityReq.ContractRawID != nil {
		if err := a.UpdateOrCreateContract(userId, activityReq.AppId, uint(*activityReq.ContractRawID)); err != nil {
			return nil, err
		}
	}

	res := models.GetDB().Create(&activity)
	if res.Error != nil {
		return nil, res.Error
	}

	if err := activity.LoadBindedContract(); err != nil {
		return nil, err
	}
	return &activity, nil
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

func (a *ActivityService) UpdateActivity(activityId string, req *models.UpdateActivityReq) (*models.Activity, error) {
	activity, err := models.FindActivityByCode(activityId)
	if err != nil {
		return nil, err
	}

	if req.ContractRawID != nil {
		if err := a.UpdateOrCreateContract(activity.RainbowUserId, activity.AppId, uint(*req.ContractRawID)); err != nil {
			return nil, err
		}
		activity.ContractRawID = req.ContractRawID
	}

	for _, nftConfig := range req.NFTConfigs {
		nftConfig.ActivityID = activity.ID
	}

	activity.UpdateActivityReq = *req
	if err := models.GetDB().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&activity).Error; err != nil {
		return nil, err
	}

	if err := activity.LoadBindedContract(); err != nil {
		return nil, err
	}
	return activity, nil
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

func (a *ActivityService) HandlePOAPH5Mint(req *MintReq) (*models.POAPResult, error) {
	activity, err := models.FindActivityByCode(req.ActivityID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = a.CheckMintable(activity, req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	token, err := middlewares.GenerateRainbowOpenJWT(activity.RainbowUserId, activity.AppId)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	genMetadataUri := func() (string, *models.NFTConfig, error) {

		// var metadataURI *string
		var index int
		// if req.ActivityID == viper.GetString("changAnDao.activityId") {
		// 	config.ActivityType = enums.ACTIVITY_SINGLE_ID_ORDER
		// } // TMP code

		switch activity.ActivityType {
		// case enums.ACTIVITY_SINGLE_ID_ORDER:
		// 	profile, err := utils.GetContractProfile(config.Contract.ContractAddress, token)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	nextTokenId = uint64(*profile.MaxTokenId) + 1
		// 	// metaUri := utils.ChangAnDaoMetadataUriFromId(nextTokenId)
		// 	metadataURI, err = createMetadata(config, token, 0)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	metaUri := strings.Replace(*metadataURI, "{id}", fmt.Sprintf("%d", nextTokenId), 0)
		// 	metadataURI = &metaUri

		case enums.ACTIVITY_SINGLE:
			index = 0
			// metadataURI, err = createMetadata(config, token, 0)
			// if err != nil {
			// 	return nil, err
			// }
			// if nextTokenId != 0 {
			// 	metaUri := strings.Replace(*metadataURI, "{id}", fmt.Sprintf("%d", nextTokenId), -1)
			// }
			// metadataURI = &metaUri
			// return metadataURI, nil

		case enums.ACTIVITY_BLINDBOX:
			probabilities := make([]float32, 0)
			for i := 0; i < len(activity.NFTConfigs); i++ {
				probabilities = append(probabilities, activity.NFTConfigs[i].Probability)
			}

			index = weightedRandomIndex(probabilities)
			// metadataURI, err = createMetadata(config, token, index)
			// if err != nil {
			// 	return nil, errors.WithStack(err)
			// }
			// default:
			// 	metadataURI = &config.MetadataUri
		}

		nftConfig := activity.NFTConfigs[index]
		metadataUri := activity.MetadataUri

		if metadataUri == "" {
			metadataUri, err = createMetadata(activity, token, index)
			if err != nil {
				return "", nil, errors.WithStack(err)
			}
		}

		return metadataUri, &nftConfig, nil
	}

	metadataURI, nftConfig, err := genMetadataUri()
	if err != nil {
		return nil, err
	}

	mint := func() (*openapiclient.ModelsMintTask, error) {
		chain, err := utils.ChainById(uint(activity.Contract.ChainId))
		if err != nil {
			return nil, errors.WithStack(err)
		}

		// create mint meta
		mintMeta := openapiclient.ServicesCustomMintDto{
			Chain:           chain,
			ContractAddress: activity.Contract.ContractAddress,
			MintToAddress:   req.UserAddress,
			MetadataUri:     &metadataURI,
		}

		if activity.IsTokenIdOrdered != nil && *activity.IsTokenIdOrdered {
			profile, err := utils.GetContractProfile(activity.Contract.ContractAddress, token)
			if err != nil {
				return nil, err
			}
			nextTokenId := uint64(*profile.MaxTokenId) + 1
			tokenIdStr := strconv.Itoa(int(nextTokenId))
			mintMeta.TokenId = &tokenIdStr
			// metaUri := utils.ChangAnDaoMetadataUriFromId(nextTokenId)
		}

		resp, err := utils.SendCustomMintRequest(token, mintMeta)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		return resp, nil
	}
	resp, err := mint()
	if err != nil {
		return nil, err
	}

	saveResult := func() (*models.POAPResult, error) {
		// create mint result
		// compatible with old activity
		// fileUrl := ""
		// if len(activity.NFTConfigs) > index {
		// 	fileUrl = activity.NFTConfigs[index].ImageURL
		// }

		item := &models.POAPResult{
			ConfigID:      int32(activity.ID),
			Address:       req.UserAddress,
			ContractRawID: *activity.ContractRawID,
			TxID:          *resp.Id,
			ActivityCode:  activity.ActivityCode,
			FileURL:       nftConfig.ImageURL,
			ProjectorId:   activity.RainbowUserId,
			AppId:         activity.AppId,
		}
		err := models.GetDB().Create(&item).Error
		return item, err
	}

	result, err := saveResult()
	if err != nil {
		return nil, err
	}

	cache, err := models.GetMintCountCache(req.ActivityID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	cache.Increase()

	return result, nil
}

func (a *ActivityService) GetMintCount(activityID, address string) (*int32, error) {
	config, err := models.FindActivityByCode(activityID)
	if err != nil {
		return nil, err
	}
	var count int32

	// phone white list logic: if whiteList config opened and user not in whiteList then the mint count is 0
	if config.IsPhoneWhiteListOpened {
		users, err := models.FindWalletUserByAddress(address)
		if err == nil { // TODO check the phone not found case
			var isInWhiteList bool
			for _, u := range users {
				isInWhiteList = models.IsPhoneInWhiteList(activityID, u.Phone)
				if isInWhiteList {
					break
				}
			}

			if !isInWhiteList {
				count = 0
				return &count, nil
			}
		}
	}

	mintedCount, err := models.GetMintSumByAddresses(activityID, address)
	if err != nil {
		return nil, err
	}

	var remainedMinted int32
	if config.MaxMintCount == -1 {
		remainedMinted = -1
	} else {
		remainedMinted = int32(int64(config.MaxMintCount) - mintedCount)
	}
	logrus.WithField("remain", remainedMinted).Info("get remain mint count")

	if config.Amount == -1 {
		count = remainedMinted
	} else {
		if remainedMinted == -1 {
			cache, err := models.GetMintCountCache(activityID)
			if err != nil {
				return nil, err
			}
			count = config.Amount - int32(cache.Count) // Amount - total minted count
		} else {
			if config.Amount-int32(mintedCount) < remainedMinted {
				count = config.Amount - int32(mintedCount)
			} else {
				count = remainedMinted
			}
		}
	}
	return &count, nil
}

func (a *ActivityService) CheckMintable(config *models.Activity, req *MintReq) error {
	if err := config.VerifyMintable(); err != nil {
		return errors.WithStack(err)
	}

	addrsOfPhone := []string{req.UserAddress}
	if config.IsPhoneWhiteListOpened {
		users, err := models.FindWalletUserByAddress(req.UserAddress)

		if err != nil {
			return err
		}

		if len(users) == 0 {
			return errors.New("无领取资格")
		}

		var isInWhiteList bool
		for _, u := range users {
			isInWhiteList = models.IsPhoneInWhiteList(req.ActivityID, u.Phone)
			if isInWhiteList {
				break
			}
		}

		if !isInWhiteList {
			return errors.New("无领取资格")
		}

		addrs, err := models.FindAllUserAddrsOfPhone(users[0].Phone)
		if err != nil {
			return err
		}
		addrsOfPhone = append(addrsOfPhone, addrs...)
	}

	if err := checkUserMintQuota(config.ActivityCode, addrsOfPhone, config.MaxMintCount); err != nil {
		return err
	}

	if req.Command != config.Command {
		return errors.New("the command is wrong")
	}

	return nil
}

func createMetadata(config *models.Activity, token string, index int) (string, error) {
	// metadataUri := config.NFTConfigs[index].MetataUri
	// if metadataUri != "" {
	// 	return &metadataUri, nil
	// }

	attributes := make([]openapiclient.ModelsExposedMetadataAttribute, 0)
	for _, v := range config.NFTConfigs[index].MetadataAttributes {
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
		Description: config.Description,
		Image:       config.NFTConfigs[index].ImageURL,
		Name:        config.NFTConfigs[index].Name,
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
		return fmt.Errorf("the mint amount has exceeded the personal limit")
	}
	return nil
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
