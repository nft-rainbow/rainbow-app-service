package services

import (
	"bytes"

	"fmt"
	"image"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"path"
	"strings"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/bwmarrin/discordgo"
	"github.com/fogleman/gg"
	. "github.com/nft-rainbow/rainbow-app-service/appService-errors"
	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func BindCfxAddress(userId, userAddress string, socialTool enums.SocialToolType) (string, string, error) {
	config := &models.SocialUserConfig{
		UserId:     userId,
		CFXAddress: userAddress,
		SocialTool: socialTool,
	}

	err := func() error {
		addr, err := cfxaddress.New(config.CFXAddress, uint32(utils.CONFLUX_MAINNET_ID))
		if err != nil {
			return errors.Wrap(ERR_BIND_ADDRESS_WRONG_FORMAT, err.Error())
		}

		result, err := models.FindSocialUserConfig(config.UserId, config.SocialTool)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			config.CFXAddress = addr.String()
			err = models.GetDB().Create(&config).Error
		} else {
			result.CFXAddress = addr.String()
			err = models.GetDB().Save(result).Error
		}
		if err != nil {
			return errors.Wrap(ERR_BIND_ADDRESS_OTHER, err.Error())
		}
		return nil
	}()
	if err != nil {
		return "", "", err
	}

	return GetBindAddress(config.UserId, config.SocialTool)
}

func GetBindAddress(userDodoSourceId string, socialTool enums.SocialToolType) (string, string, error) {
	userConfig, err := models.FindSocialUserConfig(userDodoSourceId, socialTool)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", "", ERR_BUSINESS_NOT_BIND_WALLET
		}
		return "", "", err
	}

	testAddr, err := cfxaddress.New(userConfig.CFXAddress, uint32(utils.CONFLUX_TEST_ID))
	if err != nil {
		return "", "", err
	}

	return userConfig.CFXAddress, testAddr.String(), nil
}

func GetDiscordChannelInfo(guildId string) ([]*discordgo.Channel, error) {
	st, err := GetSession().GuildChannels(guildId)

	if err != nil {
		return nil, err
	}
	return st, err
}

// func GetDoDoChannelInfo(islandId string) ([]*dodoModel.ChannelElement, error) {
// 	st, err := (*GetInstance()).GetChannelList(context.Background(), &dodoModel.GetChannelListReq{
// 		IslandSourceId: islandId,
// 	})

// 	if err != nil {
// 		return nil, err
// 	}
// 	return st, err
// }

func GetDiscordGuildInfo(guildId string) (st *discordgo.Guild, err error) {
	st, err = GetSession().Guild(guildId)
	if err != nil {
		return nil, err
	}
	return st, err
}

// func GetDoDoIslandInfo(islandId string) (st *dodoModel.GetIslandInfoRsp, err error) {
// 	info, err := (*GetInstance()).GetIslandInfo(context.Background(), &dodoModel.GetIslandInfoReq{
// 		IslandSourceId: islandId,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return info, err
// }

// func FindAuthUserServers(offset, limit int, userId uint, bot uint) (*models.UserServerQueryResult, error) {
// 	var items []*models.BotServer
// 	var cond models.BotServer
// 	cond.RainbowUserId = userId
// 	cond.SocialTool = models.SocialToolType(bot)

// 	var count int64
// 	if err := models.GetDB().Find(&items).Where(cond).Count(&count).Error; err != nil {
// 		return nil, err
// 	}

// 	if err := models.GetDB().Find(&items).Where(cond).Offset(offset).Limit(limit).Error; err != nil {
// 		return nil, err
// 	}

// 	if bot == utils.DoDo {
// 		for i := range items {
// 			if !CheckIslandIsActive(GetInstance(), items[i].RawServerId) {
// 				models.GetDB().Delete(&items[i])
// 				items = append(items[:i], items[i+1:]...)
// 			}
// 		}
// 	} else if bot == utils.Discord {
// 		for i := range items {
// 			if !CheckGuildIsActive(GetSession(), items[i].RawServerId) {
// 				models.GetDB().Delete(&items[i])
// 				items = append(items[:i], items[i+1:]...)
// 			}
// 		}
// 	}
// 	return &models.UserServerQueryResult{count, items}, nil
// }

// func GenDiscordMintRes(token, createTime, contractAddress, userAddress, userID, channelID string, id, contractId int32) (*models.CustomMintResp, error) {
// 	tokenId, hash, status, err := getTokenInfo(id, middlewares.PrefixToken(token))
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := &models.CustomMintResp{
// 		UserAddress: userAddress,
// 		NFTAddress:  viper.GetString("customMint.mintRespPrefix") + contractAddress + "/" + tokenId,
// 		Contract:    contractAddress,
// 		TokenID:     tokenId,
// 		Time:        createTime,
// 	}
// 	_, err = models.UpdateDiscordCustomCount(userID, channelID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = models.StoreCustomMintResult(models.CustomMintResult{
// 		UserID:     userID,
// 		ContractID: contractId,
// 		TokenID:    tokenId,
// 		Hash:       hash,
// 		Status:     status,
// 	})
// 	return res, nil
// }

// func GenDoDoMintRes(token, createTime, contractAddress, userAddress, userID, channelID string, id, contractId int32) (*models.CustomMintResp, error) {
// 	tokenId, hash, status, err := getTokenInfo(id, middlewares.PrefixToken(token))
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := &models.CustomMintResp{
// 		UserAddress: userAddress,
// 		NFTAddress:  viper.GetString("customMint.mintRespPrefix") + contractAddress + "/" + tokenId,
// 		Contract:    contractAddress,
// 		TokenID:     tokenId,
// 		Time:        createTime,
// 	}
// 	_, err = models.UpdateDoDoCustomCount(userID, channelID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = models.StoreCustomMintResult(models.CustomMintResult{
// 		UserID:     userID,
// 		ContractID: contractId,
// 		TokenID:    tokenId,
// 		Hash:       hash,
// 		Status:     status,
// 	})
// 	return res, nil
// }

func generateActivityURLById(activityId string) string {
	return viper.GetString("url.activity") + "?activity_id=" + activityId
}

func generateActivityUrlByFileUrl(file, activity string) string {
	tmp := strings.Split(file, "/")
	return "https://" + viper.GetString("oss.bucketName") + "." + viper.GetString("oss.endpoint") + "/" + path.Join(viper.GetString("imagesDir.minted"), activity, tmp[len(tmp)-1])
}

func generateAcvitivyPosterUrl(activityId string) string {
	return fmt.Sprintf("https://%s.%s/%s%s", viper.GetString("oss.bucketName"), viper.GetString("oss.endpoint"), viper.GetString("posterDir.activity"), activityId+".png")
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

func SyncPOAPResultStatus() {
	logrus.Info("start task for syncing poap result status")
	for {
		var results []*models.POAPResult = make([]*models.POAPResult, 0)
		models.GetDB().Where("status = ?", enums.TRANSACTION_STATUS_INIT).Limit(100).Find(&results)
		if len(results) == 0 {
			time.Sleep(time.Second * 5)
			continue
		}
		for _, v := range results {
			jwtToken, err := middlewares.GenerateRainbowOpenJWT(v.ProjectorId, v.AppId)
			if err != nil {
				logrus.Errorf("Failed to generate open JWT for %v:%v \n", v.ConfigID, err.Error())
				continue
			}
			tokenId, hash, status, err := utils.GetMintDetail(v.TxID, jwtToken)
			if status == int32(enums.TRANSACTION_STATUS_INIT) || err != nil {
				continue
			}
			v.TokenID = tokenId
			v.Hash = hash
			v.Status = enums.TransactionStatus(status)
			models.GetDB().Save(&v)
			//config, _ := models.FindPOAPActivityConfigById(v.ActivityID)
			//group := new(errgroup.Group)
			//group.Go(func() error {
			//	err := generateResultPoster(v, config.Name)
			//	if err != nil {
			//		logrus.Errorf("Failed to generate poap result poster in activity %v for %v:%v \n", v.ActivityID, v.Address, err.Error())
			//	}
			//	return err
			//})
		}
		time.Sleep(time.Second * 2)
	}

}
