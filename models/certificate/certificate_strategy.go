package certificate

import (
	"fmt"

	"github.com/mcuadros/go-defaults"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
)

type CertificateStrategy struct {
	models.BaseModel
	CertificateStrategyCore
}

type CertificateStrategyCore struct {
	CertificateType enums.CertificateType `json:"certificate_type" swaggertype:"string"`
	Name            string                `gorm:"type:varchar(255)" json:"name"`
	Description     string                `json:"description"`
	UserId          uint                  `json:"user_id"`
}

type CertiStrategyFilter struct {
	NameLike        string `form:"name_like"`
	CertificateType string `form:"certificate_type"`
	models.Pagination
}

func FindCertificateStrategies(filter CertiStrategyFilter, userId uint) (*models.ItemsWithCount[CertificateStrategy], error) {
	defaults.SetDefaults(&filter)

	sql := models.GetDB().Debug().Model(&CertificateStrategy{}).Where("user_id=?", userId)
	if filter.NameLike != "" {
		sql = sql.Where("name like ?", fmt.Sprintf("%%%s%%", filter.NameLike))
	}
	if filter.CertificateType != "" {
		certiType, err := enums.ParseCertificateType(filter.CertificateType)
		if err != nil {
			return nil, err
		}
		sql = sql.Where("certificate_type = ?", certiType)
	}

	var result models.ItemsWithCount[CertificateStrategy]
	var count int64
	if err := sql.Count(&count).Order("id desc").Offset(filter.Offset()).Limit(filter.Limit).Find(&result.Items).Error; err != nil {
		return nil, err
	}

	result.Count = int(count)
	return &result, nil
}

func FindCertificateStrategyById(id uint) (*CertificateStrategy, error) {
	var cs CertificateStrategy
	err := models.GetDB().Where("id=?", id).First(&cs).Error
	return &cs, err
}

func (cs *CertificateStrategy) CheckQualified(userAddress string) (bool, error) {
	return GetCertiOperator(cs).CheckQualified(userAddress)
}

func (cs *CertificateStrategy) GetCertificates(offset int, limit int) (*CertificatesQueryResult[any], error) {
	return GetCertiOperator(cs).GetCertificates(offset, limit)
}

func (cs *CertificateStrategy) InsertCertificates(items []any) error {
	return GetCertiOperator(cs).InsertCertificates(items)
}

func (cs *CertificateStrategy) DeleteCertificates(ids []uint) error {
	return GetCertiOperator(cs).DeleteCertificates(ids)
}
