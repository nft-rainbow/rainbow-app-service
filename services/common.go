package services

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/bwmarrin/discordgo"
	"github.com/disintegration/imaging"
	dodoModel "github.com/dodo-open/dodo-open-go/model"
	"github.com/fogleman/gg"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/viper"
	"image"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func bindCFXAddressWithDiscord(req *models.BindCFXWithDiscord) error {
	res := models.GetDB().Create(&req)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func bindCFXAddressWithDoDo(req *models.BindCFXWithDoDo) error {
	res := models.GetDB().Create(&req)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func GetDoDoBindCFXAddress(userID string) (string, error) {
	resp, err := models.FindDoDoBindingCFXAddressById(userID)
	if err != nil {
		return "", err
	}
	return resp.CFXAddress, nil
}

func GetDiscordBindCFXAddress(userID string) (string, error) {
	resp, err := models.FindDiscordBindingCFXAddressById(userID)
	if err != nil {
		return "", err
	}
	return resp.CFXAddress, nil
}

func HandleBindCfxAddress(userId, userAddress, platform string) error {
	var err error
	_, err = utils.CheckCfxAddress(utils.CONFLUX_TEST, userAddress)
	if err != nil {
		return err
	}

	if platform == "discord" {
		err = bindCFXAddressWithDiscord(&models.BindCFXWithDiscord{
			DiscordId:  userId,
			CFXAddress: userAddress,
		})
	} else if platform == "dodo" {
		err = bindCFXAddressWithDoDo(&models.BindCFXWithDoDo{
			DoDoId:     userId,
			CFXAddress: userAddress,
		})
	}
	if err != nil {
		return err
	}

	return nil
}

func GetDiscordChannelInfo(guildId string) ([]*discordgo.Channel, error) {
	st, err := GetSession().GuildChannels(guildId)

	if err != nil {
		return nil, err
	}
	return st, err
}

func GetDoDoChannelInfo(islandId string) ([]*dodoModel.ChannelElement, error) {
	st, err := (*GetInstance()).GetChannelList(context.Background(), &dodoModel.GetChannelListReq{
		IslandId: islandId,
	})

	if err != nil {
		return nil, err
	}
	return st, err
}

func GetDiscordGuildInfo(guildId string) (st *discordgo.Guild, err error) {
	st, err = GetSession().Guild(guildId)
	if err != nil {
		return nil, err
	}
	return st, err
}

func GetDoDoIslandInfo(islandId string) (st *dodoModel.GetIslandInfoRsp, err error) {
	info, err := (*GetInstance()).GetIslandInfo(context.Background(), &dodoModel.GetIslandInfoReq{
		IslandId: islandId,
	})
	if err != nil {
		return nil, err
	}
	return info, err
}

func GenDiscordMintRes(token, createTime, contractAddress, userAddress, userID, channelID string, id, contractId int32) (*models.CustomMintResp, error) {
	tokenId, hash, status, err := getTokenInfo(id, "Bearer "+token)
	if err != nil {
		return nil, err
	}

	res := &models.CustomMintResp{
		UserAddress: userAddress,
		NFTAddress:  viper.GetString("customMint.mintRespPrefix") + contractAddress + "/" + tokenId,
		Contract:    contractAddress,
		TokenID:     tokenId,
		Time:        createTime,
	}
	_, err = models.UpdateDiscordCustomCount(userID, channelID)
	if err != nil {
		return nil, err
	}

	err = models.StoreCustomMintResult(models.CustomMintResult{
		UserID:     userID,
		ContractID: contractId,
		TokenID:    tokenId,
		Hash:       hash,
		Status:     status,
	})
	return res, nil
}

func GenDoDoMintRes(token, createTime, contractAddress, userAddress, userID, channelID string, id, contractId int32) (*models.CustomMintResp, error) {
	tokenId, hash, status, err := getTokenInfo(id, "Bearer "+token)
	if err != nil {
		return nil, err
	}

	res := &models.CustomMintResp{
		UserAddress: userAddress,
		NFTAddress:  viper.GetString("customMint.mintRespPrefix") + contractAddress + "/" + tokenId,
		Contract:    contractAddress,
		TokenID:     tokenId,
		Time:        createTime,
	}
	_, err = models.UpdateDoDoCustomCount(userID, channelID)
	if err != nil {
		return nil, err
	}

	err = models.StoreCustomMintResult(models.CustomMintResult{
		UserID:     userID,
		ContractID: contractId,
		TokenID:    tokenId,
		Hash:       hash,
		Status:     status,
	})
	return res, nil
}

//func sendBurnNFTRequest(token string, dto openapiclient.ServicesBurnDto) (*openapiclient.ModelsBurnTask, error) {
//	fmt.Println("Start to burn")
//	resp, _, err := newClient().BurnsApi.BurnNft(context.Background()).Authorization(token).BurnDto(dto).Execute()
//	if err != nil {
//		return nil, err
//	}
//
//	return resp, nil
//}

func sendBatchBurnNFTRequest(token string, dto openapiclient.ServicesBurnBatchDto) ([]openapiclient.ModelsBurnTask, error) {
	fmt.Println("Start to Batch burn")
	resp, _, err := newClient().BurnsApi.BurnBatch(context.Background()).Authorization(token).BurnBatchDto(dto).Execute()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func sendCustomMintRequest(token string, dto openapiclient.ServicesCustomMintDto) (*openapiclient.ModelsMintTask, error) {
	//configuration := openapiclient.NewConfiguration()
	//apiClient := openapiclient.NewAPIClient(configuration)
	fmt.Println("Start to mint")
	resp, _, err := newClient().MintsApi.CustomMint(context.Background()).Authorization(token).CustomMintDto(dto).Execute()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func sendCreateMetadataRequest(token string, dto openapiclient.ServicesMetadataDto) (*openapiclient.ModelsExposedMetadata, error) {
	resp, _, err := newClient().MetadataApi.CreateMetadata(context.Background()).Authorization(token).MetadataInfo(dto).Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func getTokenInfo(id int32, token string) (string, string, int32, error) {
	//configuration := openapiclient.NewConfiguration()
	//apiClient := openapiclient.NewAPIClient(configuration)
	//fmt.Println("Start to get token Id")
	resp, _, err := newClient().MintsApi.GetMintDetail(context.Background(), id).Authorization(token).Execute()
	if err != nil {
		return "", "", 0, err
	}

	for *resp.Status != 1 && *resp.Hash == "" {
		resp, _, err = newClient().MintsApi.GetMintDetail(context.Background(), id).Authorization(token).Execute()
		if err != nil {
			return "", "", 0, err
		}
		time.Sleep(3 * time.Second)
	}
	return *resp.TokenId, *resp.Hash, resp.GetStatus(), nil
}

func getBurnInfo(id int32, token string) (int32, string, error) {
	//configuration := openapiclient.NewConfiguration()
	//apiClient := openapiclient.NewAPIClient(configuration)
	//fmt.Println("Start to get token Id")
	resp, _, err := newClient().BurnsApi.GetBurnDetail(context.Background(), id).Authorization(token).Execute()
	if err != nil {
		return 0, "", err
	}
	for *resp.Status == 0 && *resp.Hash == "" {
		resp, _, err = newClient().BurnsApi.GetBurnDetail(context.Background(), id).Authorization(token).Execute()
		if err != nil {
			return 0, "", err
		}
		time.Sleep(3 * time.Second)
	}
	return *resp.Status, *resp.Hash, nil
}

func GetContractInfo(id int32, token string) (*openapiclient.ModelsContract, error) {
	//configuration := openapiclient.NewConfiguration()
	//apiClient := openapiclient.NewAPIClient(configuration)
	fmt.Println("Start to get contract information")
	resp, _, err := newClient().ContractApi.GetContractInfo(context.Background(), id).Authorization(token).Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GenerateActivityPoster(url, name, description, start, end, activityUrl string) error {
	templateImg, err := gg.LoadImage("../assets/images/Activity Poster.png")

	dc := gg.NewContext(templateImg.Bounds().Dx(), templateImg.Bounds().Dy())
	dc.DrawImage(templateImg, 0, 0)

	resp, err := http.Get(url)
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
	targetUrl := activityUrl
	qrCode, _ := qrcode.New(targetUrl, qrcode.Low)
	qrImg := qrCode.Image(268)
	dc.DrawImage(qrImg, 1112, 2212)

	// 增加文字
	err = dc.LoadFontFace("../assets/fonts/PingFang SC Bold.ttf", 88)
	if err != nil {
		panic(err)
	}
	dc.SetHexColor("#05001F")
	dc.DrawStringAnchored(name, 120, 1580, 0, 0)

	err = dc.LoadFontFace("../assets/fonts/PingFang SC Bold.ttf", 64)
	if err != nil {
		panic(err)
	}
	lines := []string{"", ""}
	curLine := 0
	var lineLen float64
	for _, r := range description {
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
	dc.DrawStringAnchored(fmt.Sprintf("开始时间：%v", start), 120, 1988, 0, 0)
	dc.DrawStringAnchored(fmt.Sprintf("结束时间：%v", end), 120, 2100, 0, 0)
	f, err := os.Create("output.jpg")

	defer f.Close()

	err = jpeg.Encode(f, dc.Image(), &jpeg.Options{Quality: 95})
	if err != nil {
		panic(err)
	}

	return nil
}

func GenerateResultPoster(result models.POAPResult, name, activityUrl string) error {
	templateImg, err := gg.LoadImage("../assets/images/Result Poster.png")

	dc := gg.NewContext(templateImg.Bounds().Dx(), templateImg.Bounds().Dy())
	dc.DrawImage(templateImg, 0, 0)

	resp, err := http.Get(result.FileURL)
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
	targetUrl := activityUrl
	qrCode, _ := qrcode.New(targetUrl, qrcode.Low)
	qrImg := qrCode.Image(268)
	dc.DrawImage(qrImg, 1112, 2212)

	// 增加文字
	err = dc.LoadFontFace("../assets/fonts/PingFang SC Bold.ttf", 88)
	if err != nil {
		panic(err)
	}
	dc.SetHexColor("#05001F")
	dc.DrawStringAnchored(name, 120, 1708, 0, 0)

	err = dc.LoadFontFace("../assets/fonts/PingFang SC Bold.ttf", 64)
	if err != nil {
		panic(err)
	}

	dc.DrawStringAnchored(fmt.Sprintf("徽章编号：%v", result.TokenID), 120, 1908, 0, 0)
	dc.DrawStringAnchored(fmt.Sprintf("领取时间：%v", result.CreatedAt.Format("2006-01-02")), 120, 2036, 0, 0)
	dc.DrawStringAnchored(fmt.Sprintf("由「%v」拥有", utils.SimpleAddress(result.Address)), 120, 1580, 0, 0)
	f, err := os.Create("output.jpg")

	defer f.Close()

	err = jpeg.Encode(f, dc.Image(), &jpeg.Options{Quality: 95})
	if err != nil {
		panic(err)
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
	if err := bucket.PutObject(path.Join(viper.GetString("imagesDir.nonMinted"), activity, name), bytes.NewReader(imgData)); err != nil {
		return err
	}
	if err := bucket.PutObject(path.Join(viper.GetString("imagesDir.minted"), activity, name), bytes.NewReader(withLogo)); err != nil {
		return err
	}

	return nil
}

func getOSSBucket(bucketName string) (*oss.Bucket, error) {
	endpoint := viper.GetString("oss.endpoint")
	client, err := oss.New("https://"+endpoint, viper.GetString("oss.accessKeyId"), viper.GetString("oss.accessKeySecret"))
	if err != nil {
		return nil, err
	}
	err = checkAndCreateBucket(client, bucketName)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

func addLogo(img image.Image, logo image.Image) ([]byte, error) {
	b := img.Bounds()
	withLogo := image.NewRGBA(b)
	draw.Draw(withLogo, b, img, image.Point{0, 0}, draw.Src)
	draw.Draw(withLogo, image.Rectangle{
		Min: image.Point{b.Max.X - logo.Bounds().Dx(), b.Max.Y - logo.Bounds().Dy()},
		Max: b.Max,
	}, logo, image.Point{0, 0}, draw.Over)

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, withLogo, nil); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func drawStringWithColor(dc *gg.Context, text string, x, y float64, color string) {
	parts := strings.Split(text, " ")
	for i, part := range parts {
		if i%2 == 0 {
			dc.SetRGB(0, 0, 0)
		} else {
			dc.SetHexColor(color)
		}
		dc.DrawString(part, x, y)
		w, _ := dc.MeasureString(part)
		x += w
	}
}

func checkAndCreateBucket(client *oss.Client, bucketName string) error {
	exist, err := client.IsBucketExist(bucketName)
	if err != nil {
		return err
	}
	if !exist {
		err = client.CreateBucket(bucketName)
		if err != nil {
			return err
		}
	}
	return nil
}

func newClient() *openapiclient.APIClient {
	configuration := openapiclient.NewConfiguration()
	configuration.HTTPClient = http.DefaultClient
	configuration.Servers = openapiclient.ServerConfigurations{
		{
			URL: viper.GetString("rainbowOpenApi") + "/v1",
		},
	}
	apiClient := openapiclient.NewAPIClient(configuration)
	return apiClient
}

func SyncNFTMintTaskStatus(token string, res *models.POAPResult) {
	logrus.Info("start task for syncing nft mint status")
	tokenId, hash, status, _ := getTokenInfo(res.TxID, "Bearer "+token)

	res.TokenID = tokenId
	res.Hash = hash
	res.Status = status

	models.GetDB().Save(&res)
}

func SyncNFTBurnTaskAndMint(token, address, chain string, res *models.BatchBurnResult, config *models.NewYearConfig) {
	logrus.Info("start task for syncing nft burn status")
	status, hash, err := getBurnInfo(res.BurnID, "Bearer "+token)
	if err != nil || status != 1 {
		logrus.Info(fmt.Printf("failed to burn NFTs for %v", res.BurnID))
		return
	}
	res.Status = status
	res.Hash = hash

	models.GetDB().Save(&res)

	resp, index, err := randomMint(config, token, address, chain)
	if err != nil {
		logrus.Info(fmt.Printf("failed to mint special NFTs for %v", address))
	}

	item := &models.POAPResult{
		ConfigID:   int32(config.ID),
		Address:    address,
		ContractID: config.ContractID,
		TxID:       *resp.Id,
		TokenID:    config.ContractInfos[index].TokenID,
		ActivityID: config.ActivityID,
	}

	models.GetDB().Create(&item)
	cache := models.Cache[config.ActivityID]
	cache.Lock()
	cache.Count += 1
	cache.Unlock()

	go SyncNFTMintTaskStatus(token, item)
}
