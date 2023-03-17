package models

type PhoneWhiteList struct {
	BaseModel
	ActivityId string `gorm:"type:varchar(256);index" json:"activity_id"`
	Phone      string `gorm:"type:varchar(256);index" json:"phone"`
}

func IsPhoneInWhiteList(activityId, phone string) bool {
	var count int64
	db.Model(&PhoneWhiteList{}).Where("activity_id = ? AND phone = ?", activityId, phone).Count(&count)
	return count > 0
}
