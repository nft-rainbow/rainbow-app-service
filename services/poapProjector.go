package services

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"image"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type POAPRequest struct {
	ActivityID  string `json:"activity_id" binding:"required"`
	UserAddress string `json:"user_address" binding:"required"`
	Command     string `json:"command"`
}

func POAPActivityConfig(config *models.POAPActivityConfig, id uint) (*models.POAPActivityConfig, error) {
	config.RainbowUserId = int32(id)

	poapId, err := getPOAPId(config.ContractAddress, config.Name)
	if err != nil {
		return nil, err
	}

	config.ActivityID = poapId

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return nil, res.Error
	}
	group := new(errgroup.Group)
	group.Go(func() error {
		err := generateActivityPoster(config)
		if err != nil {
			fmt.Printf("Failed to gen poster for activity %v:%v \n", config.ActivityID, err.Error())
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
	if oldConfig.ContractID == 0 || config.ContractID != oldConfig.ContractID {
		token, err := middlewares.GenPOAPOpenJWTByRainbowUserId(*oldConfig)
		if err != nil {
			return nil, err
		}
		info, err := GetContractInfo(config.ContractID, "Bearer "+token)
		if err != nil {
			return nil, err
		}
		oldConfig.ContractType = *info.Type
		oldConfig.ChainId = *info.ChainId
		oldConfig.ChainType = *info.ChainType
		oldConfig.AppId = *info.AppId
		oldConfig.ContractAddress = *info.Address
		oldConfig.ContractID = config.ContractID
	}

	oldConfig.NFTConfigs = config.NFTConfigs
	oldConfig.AppName = config.AppName
	oldConfig.ActivityType = config.ActivityType
	oldConfig.Command = config.Command
	oldConfig.StartedTime = config.StartedTime
	oldConfig.EndedTime = config.EndedTime
	oldConfig.Amount = config.Amount
	oldConfig.ActivityPictureURL = config.ActivityPictureURL
	oldConfig.Name = config.Name
	oldConfig.Description = config.Description
	oldConfig.WhiteListInfos = config.WhiteListInfos

	if oldConfig.NFTConfigs != nil {
		deleteObjects := make([]string, 0)
		for _, v := range oldConfig.NFTConfigs {
			tmp := strings.Split(v.ImageURL, "/")
			deleteObjects = append(deleteObjects, path.Join(viper.GetString("imagesDir.minted"), oldConfig.ActivityID, tmp[len(tmp)-1]))
		}
		bucket, err := getOSSBucket(viper.GetString("oss.bucketName"))
		group := new(errgroup.Group)

		group.Go(func() error {
			_, err = bucket.DeleteObjects(deleteObjects)
			if err != nil {
				fmt.Printf("Failed to delete old NFTConfigs for %v: %v \n", config.ActivityID, err.Error())
			}
			return err
		})

		for _, v := range config.NFTConfigs {
			tmp := strings.Split(v.ImageURL, "/")
			group.Go(func() error {
				err = AddLogoAndUpload(v.ImageURL, tmp[len(tmp)-1], oldConfig.ActivityID)
				if err != nil {
					fmt.Printf("Failed to add logo and upload new NFTConfigs for %v: %v \n", config.ActivityID, err.Error())
				}
				return err
			})

		}
	}

	models.GetDB().Save(&oldConfig)

	return oldConfig, nil
}

func HandlePOAPCSVMint(req *POAPRequest) (*models.POAPResult, error) {
	config, err := models.FindPOAPActivityConfigById(req.ActivityID)
	if err != nil {
		return nil, err
	}

	if len(config.WhiteListInfos) == 0 {
		return nil, fmt.Errorf("The activity has not opened the white list")
	}

	token, err := middlewares.GeneratePOAPOpenJWT(config.RainbowUserId, config.AppId)
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
	err = checkPersonalAmount(config.ActivityID, req.UserAddress, config.MaxMintCount)
	if err != nil {
		return nil, err
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

	resp, err := sendCustomMintRequest("Bearer "+token, openapiclient.ServicesCustomMintDto{
		Chain:           chain,
		ContractAddress: config.ContractAddress,
		MintToAddress:   req.UserAddress,
		MetadataUri:     metadataURI,
	})
	if err != nil {
		return nil, err
	}

	item := &models.POAPResult{
		ConfigID:    int32(config.ID),
		Address:     req.UserAddress,
		ContractID:  config.ContractID,
		TxID:        *resp.Id,
		ActivityID:  config.ActivityID,
		ProjectorId: config.RainbowUserId,
		AppId:       config.AppId,
	}

	res := models.GetDB().Create(&item)

	cache, err := models.InitCache(item)
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
		return nil, fmt.Errorf("The activity has opened the white list")
	}

	token, err := middlewares.GeneratePOAPOpenJWT(config.RainbowUserId, config.AppId)
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
	var index int
	if config.ActivityType == utils.SINGLE {
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
	}

	resp, err := sendCustomMintRequest("Bearer "+token, openapiclient.ServicesCustomMintDto{
		Chain:           chain,
		ContractAddress: config.ContractAddress,
		MintToAddress:   req.UserAddress,
		MetadataUri:     metadataURI,
	})
	if err != nil {
		return nil, err
	}

	item := &models.POAPResult{
		ConfigID:    int32(config.ID),
		Address:     req.UserAddress,
		ContractID:  config.ContractID,
		TxID:        *resp.Id,
		ActivityID:  config.ActivityID,
		FileURL:     config.NFTConfigs[index].ImageURL,
		ProjectorId: config.RainbowUserId,
		AppId:       config.AppId,
	}

	res := models.GetDB().Create(&item)
	cache, err := models.InitCache(item)
	if err != nil {
		return nil, err
	}
	cache.Lock()
	cache.Count += 1
	cache.Unlock()

	group := new(errgroup.Group)
	group.Go(func() error {
		err := generateResultPoster(item, config.Name)
		if err != nil {
			fmt.Printf("Failed to generate poap result poster in activity %v for %v:%v \n", config.ActivityID, req.UserAddress, err.Error())
		}
		return err
	})
	
	return item, res.Error
}

func generateActivityPoster(config *models.POAPActivityConfig) error {
	templateImg, err := gg.LoadImage("./assets/images/activityPoster.png")

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
	resp, err := models.CountPOAPResultByAddress(address, activityID)
	if err != nil {
		return nil, err
	}
	var count int32
	remainedMinted := int32(int64(config.MaxMintCount) - resp)

	if config.Amount == -1 {
		count = remainedMinted
	} else {
		res, err := models.CountPOAPResult(address)
		if err != nil {
			return nil, err
		}
		if config.Amount-int32(res) < remainedMinted {
			count = config.Amount - int32(res)
		} else {
			count = remainedMinted
		}
	}
	return &count, nil
}

func commonCheck(config *models.POAPActivityConfig, req *POAPRequest) error {
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
			AttributeName: &v.Name,
			DisplayType:   &v.DisplayType,
			TraitType:     &v.TraitType,
			Value:         &v.Value,
		})
	}

	now := time.Now().Format("2006-01-02 15:04:05 MST Mon")
	name := "mint_time"
	trait := "time"
	display := "date"
	attributes = append(attributes, openapiclient.ModelsExposedMetadataAttribute{
		AttributeName: &name,
		Value:         &now,
		TraitType:     &trait,
		DisplayType:   &display,
	})

	resp, err := sendCreateMetadataRequest("Bearer "+token, openapiclient.ServicesMetadataDto{
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

	newYearId := hex.EncodeToString(sum)
	return newYearId[:8], nil
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
