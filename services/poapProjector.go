package services

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type POAPRequest struct {
	ActivityID  string `json:"activity_id" binding:"required"`
	UserAddress string `json:"user_address" binding:"required"`
	Command     string `json:"command"`
}

func POAPActivityConfig(config *models.POAPActivityConfig, id uint) (*models.POAPActivityConfig, error) {
	config.RainbowUserId = id
	config.ActivityID = utils.GenerateIDByTimeHash("", 8)
	config.IsCommand = config.Command != ""

	if config.StartedTime == 0 {
		config.StartedTime = -1
	}
	if config.EndedTime == 0 {
		config.EndedTime = -1
	}

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return nil, res.Error
	}

	// generate event poster
	group := new(errgroup.Group)
	group.Go(func() error {
		err := generateActivityPoster(config)
		if err != nil {
			logrus.Errorf("Failed to generate poster for activity %v:%v \n", config.ActivityID, err.Error())
		}
		return err
	})

	return config, nil
}

func POAPH5Config(config *models.H5Config) (*models.H5Config, error) {
	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return nil, res.Error
	}

	return config, nil
}

func UpdatePOAPActivityConfig(config *models.POAPActivityConfig, activityId string) (*models.POAPActivityConfig, error) {
	oldConfig, err := models.FindPOAPActivityConfigById(activityId)
	if err != nil {
		return nil, err
	}

	if config.ContractID != oldConfig.ContractID {
		token, err := middlewares.GenPOAPOpenJWTByRainbowUserId(*oldConfig)
		if err != nil {
			return nil, err
		}
		if config.ContractID != nil {
			info, err := GetContractInfo(*config.ContractID, middlewares.PrefixToken(token))
			if err != nil {
				return nil, err
			}
			oldConfig.ContractType = *info.Type
			oldConfig.ChainId = *info.ChainId
			oldConfig.ChainType = *info.ChainType
			oldConfig.AppId = uint(*info.AppId)
			oldConfig.ContractAddress = info.Address
			oldConfig.ContractID = config.ContractID
		}
	}

	// Create a map of oldConfig.NFTConfigs for fast searching
	oldNFTConfigsMap := make(map[uint]*models.NFTConfig)
	newNFTConfigsMap := make(map[uint]*models.NFTConfig)

	for i, nftConfig := range oldConfig.NFTConfigs {
		oldNFTConfigsMap[nftConfig.ID] = &oldConfig.NFTConfigs[i]
	}

	for i, nftConfig := range config.NFTConfigs {
		if nftConfig.ID != 0 {
			newNFTConfigsMap[nftConfig.ID] = &config.NFTConfigs[i]
		}
	}

	// Update NFTConfigs
	for _, newNFTConfig := range config.NFTConfigs {
		if oldNFTConfig, ok := oldNFTConfigsMap[newNFTConfig.ID]; ok {
			// Update existing NFTConfig
			oldNFTConfig.Probability = newNFTConfig.Probability
			oldNFTConfig.Name = newNFTConfig.Name
			oldNFTConfig.ImageURL = newNFTConfig.ImageURL

			// Update MetadataAttributes
			oldMetadataAttributesMap := make(map[uint]*models.MetadataAttribute)
			newMetadataAttributesMap := make(map[uint]*models.MetadataAttribute)

			for j, metadataAttribute := range oldNFTConfig.MetadataAttributes {
				oldMetadataAttributesMap[metadataAttribute.ID] = oldNFTConfig.MetadataAttributes[j]
			}
			for j, metadataAttribute := range newNFTConfig.MetadataAttributes {
				if metadataAttribute.ID != 0 {
					newMetadataAttributesMap[metadataAttribute.ID] = newNFTConfig.MetadataAttributes[j]
				}
			}

			if len(newNFTConfig.MetadataAttributes) > 0 {
				for _, newMetadataAttribute := range newNFTConfig.MetadataAttributes {
					if oldMetadataAttribute, ok := oldMetadataAttributesMap[newMetadataAttribute.ID]; ok {
						// Update existing MetadataAttribute
						oldMetadataAttribute.TraitType = newMetadataAttribute.TraitType
						oldMetadataAttribute.Value = newMetadataAttribute.Value
						oldMetadataAttribute.DisplayType = newMetadataAttribute.DisplayType
						models.GetDB().Save(&oldMetadataAttribute)
					} else {
						// Create new MetadataAttribute
						newMetadataAttribute.NFTConfigID = newNFTConfig.ID
						oldNFTConfig.MetadataAttributes = append(oldNFTConfig.MetadataAttributes, newMetadataAttribute)
					}
				}
				for j := len(oldNFTConfig.MetadataAttributes) - 1; j >= 0; j-- {
					if oldNFTConfig.MetadataAttributes[j].ID == 0 {
						models.GetDB().Save(&oldNFTConfig.MetadataAttributes[j])
						continue
					}
					if _, ok := newMetadataAttributesMap[oldNFTConfig.MetadataAttributes[j].ID]; !ok {
						// Delete MetadataAttribute
						models.GetDB().Delete(&oldNFTConfig.MetadataAttributes[j])
						oldNFTConfig.MetadataAttributes = append(oldNFTConfig.MetadataAttributes[:j], oldNFTConfig.MetadataAttributes[j+1:]...)
					}
				}
			} else {
				for j, attribute := range oldNFTConfig.MetadataAttributes {
					models.GetDB().Delete(&attribute)
					oldNFTConfig.MetadataAttributes = append(oldNFTConfig.MetadataAttributes[:j], oldNFTConfig.MetadataAttributes[j+1:]...)
				}
			}
			models.GetDB().Save(&oldNFTConfig)
		} else {
			// Create new NFTConfig
			newNFTConfig.POAPActivityConfigID = oldConfig.ID
			oldConfig.NFTConfigs = append(oldConfig.NFTConfigs, newNFTConfig)
		}
	}

	// Delete NFTConfigs
	for i := len(oldConfig.NFTConfigs) - 1; i >= 0; i-- {
		if oldConfig.NFTConfigs[i].ID == 0 {
			models.GetDB().Save(&oldConfig.NFTConfigs[i])
			continue
		}
		if _, ok := newNFTConfigsMap[oldConfig.NFTConfigs[i].ID]; !ok {
			// Delete NFTConfig
			models.GetDB().Delete(&oldConfig.NFTConfigs[i])
			oldConfig.NFTConfigs = append(oldConfig.NFTConfigs[:i], oldConfig.NFTConfigs[i+1:]...)
		}
	}

	oldConfig.AppName = config.AppName
	oldConfig.MaxMintCount = config.MaxMintCount
	if !(config.Command == "" && config.IsCommand == true) {
		oldConfig.Command = config.Command
		if oldConfig.Command == "" {
			oldConfig.IsCommand = false
		} else {
			oldConfig.IsCommand = true
		}
	}

	oldConfig.StartedTime = config.StartedTime
	oldConfig.EndedTime = config.EndedTime
	oldConfig.Amount = config.Amount
	oldConfig.Name = config.Name
	oldConfig.Description = config.Description
	if len(config.WhiteListInfos) != 0 {
		models.GetDB().Delete(&oldConfig.WhiteListInfos)
	}
	oldConfig.WhiteListInfos = config.WhiteListInfos

	if oldConfig.ActivityPictureURL != config.ActivityPictureURL {
		oldConfig.ActivityPictureURL = config.ActivityPictureURL
		group := new(errgroup.Group)
		group.Go(func() error {
			err := generateActivityPoster(config)
			if err != nil {
				logrus.Errorf("Failed to generate poster for activity %v:%v \n", config.ActivityID, err.Error())
			}
			return err
		})
	}

	res := models.GetDB().Save(&oldConfig)

	return oldConfig, res.Error
}

func HandlePOAPCSVMint(req *POAPRequest) (*models.POAPResult, error) {
	config, err := models.FindPOAPActivityConfigById(req.ActivityID)
	if err != nil {
		return nil, err
	}

	if len(config.WhiteListInfos) == 0 {
		return nil, fmt.Errorf("The activity has not opened the white list")
	}

	token, err := middlewares.GenerateRainbowConsoleJWT(config.RainbowUserId, config.AppId)
	if err != nil {
		return nil, err
	}

	err = commonCheck(config, req)
	if err != nil {
		return nil, err
	}

	if !checkWhiteList(config.WhiteListInfos, req.UserAddress) {
		return nil, fmt.Errorf("The address is not listed in the white list")
	}

	err = checkWhiteListLimit(config, req.UserAddress)
	if err != nil {
		return nil, err
	}

	if config.ActivityID != "" {
		err = checkPersonalAmount(config.ActivityID, req.UserAddress, config.MaxMintCount)
		if err != nil {
			return nil, err
		}
	}

	chain, err := utils.ChainById(uint(config.ChainId))
	if err != nil {
		return nil, err
	}

	var metadataURI *string
	if config.ActivityType == utils.SINGLE {
		metadataURI, err = createMetadata(config, token, 0)
		if err != nil {
			return nil, err
		}
	} else if config.ActivityType == utils.BLIND_BOX {
		var index int
		probabilities := make([]float32, 0)
		for i := 0; i < len(config.NFTConfigs); i++ {
			probabilities = append(probabilities, config.NFTConfigs[i].Probability)
		}
		index = weightedRandomIndex(probabilities)
		metadataURI, err = createMetadata(config, token, index)
		if err != nil {
			return nil, err
		}
	}

	if err := config.CheckContractAndActivityValid(); err != nil {
		return nil, err
	}

	resp, err := sendCustomMintRequest(middlewares.PrefixToken(token), openapiclient.ServicesCustomMintDto{
		Chain:           chain,
		ContractAddress: *config.ContractAddress,
		MintToAddress:   req.UserAddress,
		MetadataUri:     metadataURI,
	})
	if err != nil {
		return nil, err
	}

	item := &models.POAPResult{
		ConfigID:    int32(config.ID),
		Address:     req.UserAddress,
		ContractID:  *config.ContractID,
		TxID:        *resp.Id,
		ActivityID:  config.ActivityID,
		ProjectorId: config.RainbowUserId,
		AppId:       config.AppId,
	}
	res := models.GetDB().Create(&item)

	cache, err := models.InitCache(req.ActivityID)
	if err != nil {
		return nil, err
	}
	cache.Lock()
	cache.Count += 1
	cache.Unlock()

	return item, res.Error
}

func HandlePOAPH5Mint(req *POAPRequest) (*models.POAPResult, error) {
	config, err := models.FindPOAPActivityConfigById(req.ActivityID)
	if err != nil {
		return nil, err
	}
	if len(config.WhiteListInfos) != 0 {
		return nil, fmt.Errorf("the activity has opened the white list")
	}
	if err = config.CheckContractAndActivityValid(); err != nil {
		return nil, err
	}

	// phone whiteList logic check
	if config.IsPhoneWhiteListOpened {
		phoneInfo, err := models.FindAnywebUserByAddress(req.UserAddress)
		if err == nil && len(phoneInfo.Phone) > 0 {
			isInWhiteList := models.IsPhoneInWhiteList(req.ActivityID, phoneInfo.Phone)
			if !isInWhiteList { // phone not in whitelist
				return nil, errors.New("无领取资格")
			}
		} else if errors.Is(err, gorm.ErrRecordNotFound) { // not found phone info
			return nil, errors.New("无领取资格")
		}
	}

	token, err := middlewares.GenerateRainbowConsoleJWT(config.RainbowUserId, config.AppId)
	if err != nil {
		return nil, err
	}

	err = commonCheck(config, req)
	if err != nil {
		return nil, err
	}

	err = checkPersonalAmount(config.ActivityID, req.UserAddress, config.MaxMintCount)
	if err != nil {
		return nil, err
	}
	chain, err := utils.ChainById(uint(config.ChainId))
	if err != nil {
		return nil, err
	}

	var metadataURI *string
	var nextTokenId uint64
	var index int
	if req.ActivityID == viper.GetString("changAnDao.activityId") { // TMP code
		nextTokenId = GetChangAnDaoNum() + 1
		metaUri := utils.ChangAnDaoMetadataUriFromId(nextTokenId)
		metadataURI = &metaUri
	} else if config.ActivityType == utils.SINGLE {
		metadataURI, err = createMetadata(config, token, 0)
		if err != nil {
			return nil, err
		}
	} else if config.ActivityType == utils.BLIND_BOX {
		probabilities := make([]float32, 0)
		for i := 0; i < len(config.NFTConfigs); i++ {
			probabilities = append(probabilities, config.NFTConfigs[i].Probability)
		}
		index = weightedRandomIndex(probabilities)
		metadataURI, err = createMetadata(config, token, index)
		if err != nil {
			return nil, err
		}
	} else { // old activity
		metadataURI = &config.MetadataUri
	}

	mintMeta := openapiclient.ServicesCustomMintDto{
		Chain:           chain,
		ContractAddress: *config.ContractAddress,
		MintToAddress:   req.UserAddress,
		MetadataUri:     metadataURI,
	}

	if req.ActivityID == viper.GetString("changAnDao.activityId") { // TMP code
		tokenIdStr := strconv.Itoa(int(nextTokenId))
		mintMeta.TokenId = &tokenIdStr
		IncreaseChangAnDaoNum()
	}

	resp, err := sendCustomMintRequest(middlewares.PrefixToken(token), mintMeta)
	if err != nil {
		return nil, err
	}

	cache, err := models.InitCache(req.ActivityID)
	if err != nil {
		return nil, err
	}

	// compatible with old activity
	fileUrl := ""
	if len(config.NFTConfigs) > index {
		fileUrl = config.NFTConfigs[index].ImageURL
	}

	item := &models.POAPResult{
		ConfigID:    int32(config.ID),
		Address:     req.UserAddress,
		ContractID:  *config.ContractID,
		TxID:        *resp.Id,
		ActivityID:  config.ActivityID,
		FileURL:     fileUrl,
		ProjectorId: config.RainbowUserId,
		AppId:       config.AppId,
	}
	res := models.GetDB().Create(&item)

	cache.Lock()
	cache.Count += 1
	cache.Unlock()

	return item, res.Error
}

func generateActivityPoster(config *models.POAPActivityConfig) error {
	if err := config.CheckActivityValid(); err != nil {
		return err
	}

	templateImg, err := gg.LoadImage("./assets/images/activityPoster.png")
	if err != nil {
		return err
	}

	dc := gg.NewContext(templateImg.Bounds().Dx(), templateImg.Bounds().Dy())
	dc.DrawImage(templateImg, 0, 0)

	resp, err := http.Get(config.ActivityPictureURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	imgData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	img, err := imaging.Decode(bytes.NewReader(imgData))
	if err != nil {
		return err
	}
	img = imaging.Fit(img, 1260, 1260, imaging.Lanczos)
	dc.DrawImage(img, 120, 200)

	// QR Code Generate
	targetUrl := generateActivityURLById(config.ActivityID)
	qrCode, _ := qrcode.New(targetUrl, qrcode.Low)
	qrImg := qrCode.Image(268)
	dc.DrawImage(qrImg, 1112, 2212)

	// 增加文字
	err = dc.LoadFontFace("./assets/fonts/PingFang.ttf", 88)
	if err != nil {
		panic(err)
	}
	dc.SetHexColor("#05001F")
	dc.DrawStringAnchored(config.Name, 120, 1580, 0, 0)

	err = dc.LoadFontFace("./assets/fonts/PingFang.ttf", 64)
	if err != nil {
		panic(err)
	}
	lines := []string{"", ""}
	curLine := 0
	var lineLen float64
	for _, r := range config.Description {
		w, _ := dc.MeasureString(string(r))
		if lineLen+w > 1260 {
			curLine++
			lineLen = 0
		}
		lines[curLine] += string(r)
		lineLen += w
	}

	dc.SetHexColor("#696679")
	for i, line := range lines {
		dc.DrawString(line, 120, float64(1732+i*96))
	}
	var start, end string
	if config.StartedTime == -1 {
		start = "不限时"
	} else {
		start = time.Unix(config.StartedTime, 0).Format("2006-01-02")
	}
	if config.EndedTime == -1 {
		end = "不限时"
	} else {
		end = time.Unix(config.EndedTime, 0).Format("2006-01-02")
	}
	dc.DrawStringAnchored(fmt.Sprintf("开始时间：%v", start), 120, 1988, 0, 0)
	dc.DrawStringAnchored(fmt.Sprintf("结束时间：%v", end), 120, 2100, 0, 0)
	buf := new(bytes.Buffer)
	dc.EncodePNG(buf)

	bucket, err := getOSSBucket(viper.GetString("oss.bucketName"))
	if err != nil {
		return err
	}
	if err := bucket.PutObject(path.Join(viper.GetString("posterDir.activity"), config.ActivityID+".png"), buf); err != nil {
		return err
	}

	return nil
}

func generateResultPoster(result *models.POAPResult, name string) error {
	templateImg, err := gg.LoadImage("./assets/images/resultPoster.png")

	dc := gg.NewContext(templateImg.Bounds().Dx(), templateImg.Bounds().Dy())
	dc.DrawImage(templateImg, 0, 0)

	resp, err := http.Get(generateActivityUrlByFileUrl(result.FileURL, result.ActivityID))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	imgData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	img, err := imaging.Decode(bytes.NewReader(imgData))
	if err != nil {
		return err
	}
	img = imaging.Fit(img, 1260, 1260, imaging.Lanczos)
	dc.DrawImage(img, 120, 200)

	// QR Code Generate
	targetUrl := generateActivityURLById(result.ActivityID)
	qrCode, _ := qrcode.New(targetUrl, qrcode.Low)
	qrImg := qrCode.Image(268)
	dc.DrawImage(qrImg, 1112, 2212)

	// 增加文字
	err = dc.LoadFontFace("./assets/fonts/PingFang.ttf", 88)
	if err != nil {
		return err
	}
	dc.SetHexColor("#05001F")
	dc.DrawStringAnchored(name, 120, 1708, 0, 0)

	err = dc.LoadFontFace("./assets/fonts/PingFang.ttf", 64)
	if err != nil {
		return err
	}
	dc.SetHexColor("#696679")
	x := 120.00
	dc.DrawString("由「", x, 1580)
	w, _ := dc.MeasureString("由「")
	x += w
	dc.SetHexColor("#6953EF")

	dc.DrawString(fmt.Sprintf("%v", utils.SimpleAddress(result.Address)), x, 1580)
	w, _ = dc.MeasureString(fmt.Sprintf("%v", utils.SimpleAddress(result.Address)))
	x += w
	dc.SetHexColor("#696679")
	dc.DrawString("」拥有", x, 1580)

	drawTimeStringWithColor(dc, "：", fmt.Sprintf("徽章编号：%v", result.TokenID), 120, 1908, "#6953EF")
	drawTimeStringWithColor(dc, "：", fmt.Sprintf("领取时间：%v", result.CreatedAt.Format("2006-01-02")), 120, 2036, "#05001F")
	buf := new(bytes.Buffer)
	dc.EncodePNG(buf)

	bucket, err := getOSSBucket(viper.GetString("oss.bucketName"))
	if err := bucket.PutObject(path.Join(viper.GetString("posterDir.result"), result.ActivityID, result.Address, strconv.Itoa(int(result.ID))+".png"), buf); err != nil {
		return err
	}

	return nil
}

func AddLogoAndUpload(url, name, activity string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	imgData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return err
	}

	logoFile, err := os.Open("./assets/images/logo.png")
	if err != nil {
		return err
	}
	defer logoFile.Close()

	logo, _, err := image.Decode(logoFile)
	if err != nil {
		return err
	}

	withLogo, err := addLogo(img, logo)
	if err != nil {
		return err
	}

	bucket, err := getOSSBucket(viper.GetString("oss.bucketName"))
	if err := bucket.PutObject(path.Join(viper.GetString("imagesDir.minted"), activity, name), bytes.NewReader(withLogo)); err != nil {
		return err
	}

	return nil
}

func GetMintCount(activityID, address string) (*int32, error) {
	config, err := models.FindPOAPActivityConfigById(activityID)
	if err != nil {
		return nil, err
	}
	var count int32

	// phone white list logic: if whiteList config opened and user not in whiteList then the mint count is 0
	if config.IsPhoneWhiteListOpened {
		phoneInfo, err := models.FindAnywebUserByAddress(address)
		if err == nil && len(phoneInfo.Phone) > 0 { // TODO check the phone not found case
			isInWhiteList := models.IsPhoneInWhiteList(activityID, phoneInfo.Phone)
			if !isInWhiteList {
				count = 0
				return &count, nil
			}
		}
	}

	mintedCount, err := models.CountPOAPResultByAddress(address, activityID)
	if err != nil {
		return nil, err
	}

	var remainedMinted int32
	if config.MaxMintCount == -1 {
		remainedMinted = -1
	} else {
		remainedMinted = int32(int64(config.MaxMintCount) - mintedCount)
	}

	if config.Amount == -1 {
		count = remainedMinted
	} else {
		if remainedMinted == -1 {
			cache, err := models.InitCache(activityID)
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

func commonCheck(config *models.POAPActivityConfig, req *POAPRequest) error {
	if err := config.CheckActivityValid(); err != nil {
		return err
	}
	if req.Command != config.Command {
		return fmt.Errorf("The command is wrong")
	}
	if config.StartedTime != -1 && time.Now().Unix() < config.StartedTime {
		return fmt.Errorf("The activity has not been started")
	}

	if config.EndedTime != -1 && time.Now().Unix() > config.EndedTime {
		return fmt.Errorf("The activity has been expired")
	}

	err := checkAmount(config.ActivityID, config.Amount)
	if err != nil {
		return err
	}
	return nil
}

func checkWhiteList(whiteList []models.WhiteListInfo, address string) bool {
	for _, v := range whiteList {
		if address == v.User {
			return true
		}
	}
	return false
}

func checkWhiteListLimit(config *models.POAPActivityConfig, address string) error {
	if err := config.CheckActivityValid(); err != nil {
		return err
	}
	resp, err := models.CountPOAPResultByAddress(address, config.ActivityID)
	if err != nil {
		return err
	}
	for _, v := range config.WhiteListInfos {
		if v.User == address && resp >= int64(v.Count) {
			return fmt.Errorf("The NFT minted by the account has exceeded the mint limit")
		}
	}
	return nil
}

func createMetadata(config *models.POAPActivityConfig, token string, index int) (*string, error) {
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

	resp, err := sendCreateMetadataRequest(middlewares.PrefixToken(token), openapiclient.ServicesMetadataDto{
		Description: config.Description,
		Image:       config.NFTConfigs[index].ImageURL,
		Name:        config.NFTConfigs[index].Name,
		Attributes:  attributes,
	})
	if err != nil {
		return nil, err
	}

	return resp.Uri, nil
}

func getPOAPId(address string, name string) (string, error) {
	hash := sha256.New()

	_, err := hash.Write([]byte(address + name + strconv.FormatInt(time.Now().UnixNano(), 10)))
	if err != nil {
		return "", err
	}
	sum := hash.Sum(nil)

	pId := hex.EncodeToString(sum)
	return pId[:8], nil
}

func checkAmount(poapId string, amount int32) error {
	if amount != -1 {
		resp, err := models.CountPOAPResult(poapId)
		if err != nil {
			return err
		}
		if int32(resp) >= amount {
			return fmt.Errorf("The mint amount has exceeded the limit")
		}
	}
	return nil
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

func checkPersonalAmount(activityId, user string, max int32) error {
	if max == -1 {
		return nil
	}

	count, err := models.CountPOAPResultByAddress(user, activityId)
	if err != nil {
		return err
	}

	if int32(count) >= max {
		return fmt.Errorf("The mint amount has exceeded the personal limit")
	}
	return nil
}
