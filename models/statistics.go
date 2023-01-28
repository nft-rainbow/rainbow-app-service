package models

import (
	"gorm.io/gorm"
)

type Statistic struct {
	BaseModel
	Method string `gorm:"varchar(16);index" json:"method"`
	Path   string `gorm:"type:varchar(64);index" json:"path"`
	Ip     string `gorm:"type:varchar(64);index" json:"ip"`
	Count  uint   `gorm:"default:0" json:"count"`
	Date   string `gorm:"type:varchar(10);index" json:"date"`
}

func IncreaseStatistic(method string, path string, ip string, date string) error {
	s := Statistic{
		Method: method,
		Path:   path,
		Ip:     ip,
		Date:   date,
	}

	res := db.Model(&s).Where(&s).Update("count", gorm.Expr("count + 1"))
	if res.Error == nil && res.RowsAffected == 0 {
		s.Count = 1
		return db.Create(&s).Error
	}

	return res.Error
}
