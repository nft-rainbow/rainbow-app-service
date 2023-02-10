package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/nft-rainbow/rainbow-app-service/middlewares"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/utils"
	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ShareRequest struct {
	Sharer     string `json:"sharer" binding:"required"`
	Receiver   string `json:"receiver" binding:"required"`
	ActivityID string `json:"activity_id" binding:"required"`
}

type MintCountResponse struct {
	Address    string `json:"address"`
	ActivityID string `json:"activity_id"`
	Count      int32  `json:"count"`
}

var clock time.Time

func SetNewYearConfig(config *models.NewYearConfig, id uint) (*models.NewYearConfig, error) {
	config.RainbowUserId = int32(id)
	token, err := middlewares.GenOpenJWTByRainbowUserId(config.RainbowUserId, config.AppId)
	if err != nil {
		return nil, err
	}
	info, err := GetContractInfo(config.ContractID, "Bearer "+token)
	if err != nil {
		return nil, err
	}
	config.ContractType = *info.Type
	config.ContractAddress = *info.Address
	config.ChainType = *info.ChainType
	config.ChainId = *info.ChainId
	config.AppId = *info.AppId
	newYearId, err := getPoAPId(config.ContractAddress, config.Name)
	if err != nil {
		return nil, err
	}

	config.ActivityID = newYearId

	res := models.GetDB().Create(&config)
	if res.Error != nil {
		return nil, res.Error
	}

	return config, nil
}

func HandleSpecialNFTMint(req *POAPRequest) (*models.POAPResult, error) {
	config, err := models.FindNewYearConfigById(req.ActivityID)
	if err != nil {
		return nil, err
	}

	commonConfig, err := models.FindNewYearConfigById(viper.GetString("newYearEvent.newYearCommonId"))
	if err != nil {
		return nil, err
	}

	token, err := middlewares.GenOpenJWTByRainbowUserId(config.RainbowUserId, config.AppId)
	if err != nil {
		return nil, err
	}

	err = newYearCommonCheck(config.StartedTime, config.EndedTime, config.ActivityID, config.Amount)
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

	item, err := burnNFTs(commonConfig, req.UserAddress, token, chain)
	if err != nil {
		return nil, err
	}

	go SyncNFTBurnTaskAndMint(token, req.UserAddress, chain, item, config)

	return &models.POAPResult{
		Address:    req.UserAddress,
		ActivityID: req.ActivityID,
		ContractID: int32(config.ID),
	}, nil
}

func HandleCommonNFTMint(req *POAPRequest) (*models.POAPResult, error) {
	err := checkMintCount(req.UserAddress, req.ActivityID)
	if err != nil {
		return nil, err
	}

	config, err := models.FindNewYearConfigById(req.ActivityID)
	if err != nil {
		return nil, err
	}

	token, err := middlewares.GenOpenJWTByRainbowUserId(config.RainbowUserId, config.AppId)
	if err != nil {
		return nil, err
	}
	err = newYearCommonCheck(config.StartedTime, config.EndedTime, config.ActivityID, config.Amount)
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

	resp, index, err := randomMint(config, token, req.UserAddress, chain)
	if err != nil {
		return nil, err
	}

	item := &models.POAPResult{
		ConfigID:   int32(config.ID),
		Address:    req.UserAddress,
		ContractID: config.ContractID,
		TxID:       *resp.Id,
		TokenID:    config.ContractInfos[index].TokenID,
		ActivityID: config.ActivityID,
	}

	res := models.GetDB().Create(&item)
	cache := models.Cache[config.ActivityID]
	cache.Lock()
	cache.Count += 1
	cache.Unlock()

	everyDay, err := models.FindEveryDayMintCount(req.UserAddress, req.ActivityID)
	if err != nil {
		return nil, err
	}
	if everyDay.Count != 0 {
		_, err = models.UpdateEveryDayMintCount(req.UserAddress, req.ActivityID, -1)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = models.UpdateMintCount(req.UserAddress, req.ActivityID, -1)
		if err != nil {
			return nil, err
		}
	}

	go SyncNFTMintTaskStatus(token, item)

	return item, res.Error
}

func burnNFTs(config *models.NewYearConfig, address, token, chain string) (*models.BatchBurnResult, error) {
	err := checkEnough(config, address)
	if err != nil {
		return nil, err
	}
	var amount = int32(1)
	items := make([]openapiclient.ServicesBurnItemDto, 0)
	contractType, err := utils.ContractTypeByTypeId(uint(config.ContractType))
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(config.ContractInfos); i++ {
		items = append(items, openapiclient.ServicesBurnItemDto{
			Amount:  &amount,
			TokenId: strconv.Itoa(i + 1),
		})

	}
	dto := &openapiclient.ServicesBurnBatchDto{
		Chain:           chain,
		ContractAddress: config.ContractAddress,
		ContractType:    contractType,
		Items:           items,
		User:            &address,
	}

	resp, err := sendBatchBurnNFTRequest("Bearer "+token, *dto)
	if err != nil {
		return nil, err
	}

	item := &models.BatchBurnResult{
		Address:    address,
		BurnID:     *resp[0].Id,
		Status:     0,
		ActivityID: config.ActivityID,
	}
	models.GetDB().Create(&item)

	return item, nil
}

func UpdateBySharing(req ShareRequest) error {
	if req.Sharer == req.Receiver {
		return fmt.Errorf("Can not share to yourself")
	}
	err := utils.IsCfxAddress(req.Sharer)
	if err != nil {
		return err
	}
	err = utils.IsCfxAddress(req.Receiver)
	if err != nil {
		return err
	}
	err = checkAndCreateNewAccount(req.Receiver, req.ActivityID)
	if err != nil {
		return err
	}
	count, err := models.CountTodaySharerInfo(req.Sharer, req.ActivityID, clock)
	if err != nil {
		return err
	}
	if count < viper.GetInt64("newYearEvent.everyDaySharerLimit") {
		resp, err := models.FindSharingInfo(req.Sharer, req.Receiver, req.ActivityID)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			if viper.GetString("env") == "dev" {
				if resp.UpdatedAt.Unix() > clock.Unix() &&
					resp.UpdatedAt.Unix() < clock.Add(viper.GetDuration("testMinuteDuration")*time.Minute).Unix() {
					return fmt.Errorf("The sharer has shared the link to receiver")
				}
			} else if viper.GetString("env") == "prod" {
				if resp.UpdatedAt.Unix() > clock.Unix() &&
					resp.UpdatedAt.Unix() < clock.Add(24*time.Hour).Unix() {
					return fmt.Errorf("The sharer has shared the link to receiver")
				}
			}
		} else {
			item := models.ShareInfo{
				Sharer:     req.Sharer,
				Receiver:   req.Receiver,
				ActivityID: req.ActivityID,
			}
			models.GetDB().Create(&item)
		}
		_, err = models.UpdateMintCount(req.Sharer, req.ActivityID, 1)
		if err != nil {
			return err
		}
	}

	item := models.ShareInfo{
		Sharer:     req.Sharer,
		Receiver:   req.Receiver,
		ActivityID: req.ActivityID,
	}
	res := models.GetDB().Model(&item).Where(&item).Update("updated_at", time.Now())

	return res.Error
}

func GetSpecialMintCount(address, poapId string) (*int32, error) {
	config, err := models.FindNewYearConfigById(poapId)
	if err != nil {
		return nil, err
	}
	cur := int64(math.MaxInt64)

	resp, err := CommonBalanceOfBatch(config.ContractAddress, address)
	if err != nil {
		return nil, err
	}
	for i := range resp {
		if resp[i].Int64() < cur {
			cur = resp[i].Int64()
		}
	}

	tmp, err := models.CountUnhandledBurnResult(viper.GetString("newYearEvent.newYearCommonId"), address)
	if err != nil {
		return nil, err
	}

	res := int32(cur) - int32(tmp)

	return &res, nil
}

func GetCommonMintCount(address, poapId string) (*int32, error) {
	err := checkAndCreateNewAccount(address, poapId)
	if err != nil {
		return nil, err
	}
	resp, err := models.FindShareMintCount(address, poapId)
	if err != nil {
		return nil, err
	}

	tmp, err := models.FindEveryDayMintCount(address, poapId)
	if err != nil {
		return nil, err
	}
	res := resp.Count + tmp.Count
	return &res, nil
}

func UpdateEveryday() {
	resp, err := models.GetClock(viper.GetString("newYearEvent.newYearCommonId"))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		models.GetDB().Create(&models.ClockTime{
			Time:       viper.GetTime("startTime"),
			ActivityID: viper.GetString("newYearEvent.newYearCommonId"),
		})
		clock = viper.GetTime("startTime")
	} else {
		clock = resp.Time
	}
	var c <-chan time.Time
	if viper.GetString("env") == "dev" {
		target := resp.Time.Add(viper.GetDuration("testMinuteDuration") * time.Minute)

		for target.Unix() < time.Now().Unix() {
			target = target.Add(viper.GetDuration("testMinuteDuration") * time.Minute)
		}

		models.GetDB().Model(&models.ClockTime{}).
			Where("activity_id = ?",
				viper.GetString("newYearEvent.newYearCommonId")).
			Update("time", target.Add(-viper.GetDuration("testMinuteDuration")*time.Minute))

		c = time.Tick(target.Sub(time.Now()))
	} else if viper.GetString("env") == "prod" {
		target := resp.Time.Add(24 * time.Hour)
		for target.Unix() < time.Now().Unix() {
			target = target.Add(24 * time.Hour)
		}
		models.GetDB().Model(&models.ClockTime{}).
			Where("activity_id = ?",
				viper.GetString("newYearEvent.newYearCommonId")).
			Update("time", target.Add(-24*time.Hour))

		c = time.Tick(target.Sub(time.Now()))
	}

	go func() {
		for {
			<-c
			updateVal := time.Now()
			models.GetDB().Model(&models.ClockTime{}).
				Where("activity_id = ?", viper.GetString("newYearEvent.newYearCommonId")).Update("time", updateVal)

			clock = updateVal
			var cond models.EveryDayMintCount
			cond.ActivityID = viper.GetString("newYearEvent.newYearCommonId")
			models.GetDB().Model(models.EveryDayMintCount{}).Where(&cond).Not("count > ?", 0).Update("count", gorm.Expr("count + ?", 1))
			if viper.GetString("env") == "prod" {
				c = time.Tick(24 * time.Hour)
			} else if viper.GetString("env") == "dev" {
				c = time.Tick(viper.GetDuration("testMinuteDuration") * time.Minute)
			}
		}
	}()
}

func checkEnough(config *models.NewYearConfig, address string) error {
	resp, err := CommonBalanceOfBatch(config.ContractAddress, address)
	if err != nil {
		return err
	}
	for i := range resp {
		if resp[i].Int64() <= 0 {
			return fmt.Errorf("The common NFTs are not enough")
		}
	}

	return nil
}

func randomMint(config *models.NewYearConfig, token, address, chain string) (*openapiclient.ModelsMintTask, int32, error) {
	var index int
	probabilities := make([]float32, 0)
	for i := 0; i < len(config.ContractInfos); i++ {
		probabilities = append(probabilities, config.ContractInfos[i].Probability)
	}
	index = weightedRandomIndex(probabilities)

	resp, err := sendCustomMintRequest("Bearer "+token, openapiclient.ServicesCustomMintDto{
		Chain:           chain,
		ContractAddress: config.ContractAddress,
		MintToAddress:   address,
		MetadataUri:     &(config.ContractInfos[index].MetadataURI),
		TokenId:         &(config.ContractInfos[index].TokenID),
	})
	if err != nil {
		return nil, 0, err
	}

	return resp, int32(index), nil
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

func newYearCommonCheck(startTime, endTime int64, poapId string, amount int32) error {
	if startTime != -1 && time.Now().Unix() < startTime {
		return fmt.Errorf("The activity has not been started")
	}

	if endTime != -1 && time.Now().Unix() > endTime {
		return fmt.Errorf("The activity has been expired")
	}

	err := checkAmount(poapId, amount)
	if err != nil {
		return err
	}
	return nil
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

func checkMintCount(address, poapId string) error {
	err := checkAndCreateNewAccount(address, poapId)
	if err != nil {
		return err
	}
	resp, err := models.FindShareMintCount(address, poapId)
	if err != nil {
		return err
	}
	res, err := models.FindEveryDayMintCount(address, poapId)
	if err != nil {
		return err
	}

	if res.Count+resp.Count <= 0 {
		return fmt.Errorf("The mint count is not enough")
	}
	return nil
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

func checkAndCreateNewAccount(address string, poapId string) error {
	_, err := models.FindShareMintCount(address, poapId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		item := &models.ShareMintCount{
			Address:    address,
			ActivityID: poapId,
			Count:      0,
		}

		models.GetDB().Create(item)

		everyDay := &models.EveryDayMintCount{
			Address:    address,
			ActivityID: poapId,
			Count:      1,
		}
		models.GetDB().Create(everyDay)
	}

	return nil
}

func getPoAPId(address string, name string) (string, error) {
	hash := sha256.New()

	_, err := hash.Write([]byte(address + name + strconv.FormatInt(time.Now().UnixNano(), 10)))
	if err != nil {
		return "", err
	}
	sum := hash.Sum(nil)

	newYearId := hex.EncodeToString(sum)
	return newYearId[:8], nil
}
