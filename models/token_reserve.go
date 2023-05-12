package models

type TokenReserve struct {
	BaseModel
	TokenIdStart uint `gorm:"index" json:"token_id_start" binding:"required"`
	TokenIdEnd   uint `gorm:"index" json:"token_id_end" binding:"required"`
	ActivityId   uint `gorm:"index" json:"activity" binding:"required"`
}

func IsTokenReserved(activityId uint, tokenId string) (bool, error) {
	var count int64
	err := GetDB().Debug().Where("activity_id=?", activityId).Where("token_id_start>=?", tokenId).Where("token_id_end<=?", tokenId).Count(&count).Error
	return count > 0, err
}

func GetActivityResrverTokenIds(activityId uint) ([][2]uint, error) {
	var items []*TokenReserve
	err := GetDB().Where("activity_id=?", activityId).Find(&items).Error

	var result [][2]uint
	for _, item := range items {
		result = append(result, [2]uint{item.TokenIdStart, item.TokenIdEnd})
	}
	return result, err
}

func GetResrverTokenIdsByActivityCode(activityCode string) ([]*TokenReserve, error) {
	a, err := FindActivityByCode(activityCode)
	if err != nil {
		return nil, err
	}
	var items []*TokenReserve
	err = GetDB().Where("activity_id=?", a.ID).Find(&items).Error
	return items, err
}
