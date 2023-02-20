package services

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/bwmarrin/discordgo"
	dodoModel "github.com/dodo-open/dodo-open-go/model"
	"github.com/fogleman/gg"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"image"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"net/http"
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
	tokenId, hash, status, err := getTokenInfo(id, middlewares.PrefixToken(token))
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

	for *resp.Status != models.STATUS_SUCCESS && *resp.Hash == "" {
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

func generateActivityURLById(activityId string) string {
	return viper.GetString("url.activity") + "?activity_id=" + activityId
}

func generateActivityUrlByFileUrl(file, activity string) string {
	tmp := strings.Split(file, "/")
	return "https://" + viper.GetString("oss.bucketName") + "." + viper.GetString("oss.endpoint") + "/" + path.Join(viper.GetString("imagesDir.minted"), activity, tmp[len(tmp)-1])
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

func drawTimeStringWithColor(dc *gg.Context, sep, text string, x, y float64, color string) {
	parts := strings.Split(text, sep)
	for i, part := range parts {
		if i != 1 {
			dc.SetHexColor("#696679")
		} else {
			dc.SetHexColor(color)
		}
		dc.DrawString(part, x, y)

		w, _ := dc.MeasureString(part)
		x += w
		if i == 0 {
			dc.DrawString(sep, x, y)
		}
		w, _ = dc.MeasureString(sep)
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

func SyncPOAPResultStatus() {
	logrus.Info("start task for syncing nft mint status")
	for {
		var results []*models.POAPResult
		models.GetDB().Where("status = ?", models.STATUS_INIT).Limit(100).Find(&results)
		if len(results) == 0 {
			time.Sleep(time.Second * 5)
			continue
		}
		for _, v := range results {
			config, _ := models.FindPOAPActivityConfigById(v.ActivityID)
			token, err := middlewares.GeneratePOAPOpenJWT(v.ProjectorId, v.AppId)
			if err != nil {
				fmt.Printf("Failed to generate open JWT for %v:%v \n", v.ConfigID, err.Error())
			}
			tokenId, hash, status, _ := getTokenInfo(v.TxID, middlewares.PrefixToken(token))
			if status == models.STATUS_FAIL {
				models.GetDB().Delete(&v)
				continue
			}
			v.TokenID = tokenId
			v.Hash = hash
			v.Status = status
			group := new(errgroup.Group)
			group.Go(func() error {
				err := generateResultPoster(v, config.Name)
				if err != nil {
					fmt.Printf("Failed to generate poap result poster in activity %v for %v:%v \n", v.ActivityID, v.Address, err.Error())
				}
				return err
			})
			models.GetDB().Save(&v)
		}
	}

}
