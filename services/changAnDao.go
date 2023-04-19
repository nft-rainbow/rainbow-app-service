package services

import (
	"sync/atomic"

	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/spf13/viper"
)

var changAnDaoNum uint64

func InitChangAnDaoNum() {
	var count int64
	models.GetDB().Model(&models.POAPResult{}).Where("activity_id = ? and status = ?", viper.GetString("changAnDao.activityId"), models.STATUS_SUCCESS).Count(&count)
	atomic.StoreUint64(&changAnDaoNum, uint64(count))
}

func IncreaseChangAnDaoNum() {
	atomic.AddUint64(&changAnDaoNum, 1)
}

func GetChangAnDaoNum() uint64 {
	return atomic.LoadUint64(&changAnDaoNum)
}
