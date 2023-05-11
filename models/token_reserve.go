package models

type TokenReserve struct {
	BaseModel
	TokenIdStart int64 `gorm:"index"`
	TokenIdEnd   int64 `gorm:"index"`
	ActivityId   uint  `gorm:"index"`
}

func IsTokenReserved(activityId uint, tokenId string) (bool, error) {
	var count int64
	err := GetDB().Debug().Where("activity_id=?", activityId).Where("token_id_start>=?", tokenId).Where("token_id_end<=?", tokenId).Count(&count).Error
	return count > 0, err
}

func GetActivityReserveAmount(activityId uint) {
}
