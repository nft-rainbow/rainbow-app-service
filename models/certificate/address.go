package certificate

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/pkg/errors"
)

type AddressCertificateInsertPart struct {
	Address string `json:"address"`
}
type AddressCertificate struct {
	models.BaseModel
	AddressCertificateInsertPart
	CertificateStrategyID uint `json:"certificate_strategy_id"`
}

type AddressCertiOperator struct {
	Strategy *CertificateStrategy
}

func (a *AddressCertiOperator) CheckQualified(userAddress string) (bool, error) {
	if a.Strategy == nil || a.Strategy.CertificateType != enums.CERTIFICATE_ADDRESS {
		return false, nil
	}
	var count int64
	if err := models.GetDB().Model(&AddressCertificate{}).
		Where("address=? and certificate_strategy_id=?", userAddress, a.Strategy.ID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a *AddressCertiOperator) GetCertificates(offset int, limit int) (*CertificatesQueryResult[any], error) {
	certificates := CertificatesQueryResult[*AddressCertificate]{
		CertificateType: enums.CERTIFICATE_ADDRESS,
	}

	err := models.GetDB().Model(&AddressCertificate{}).
		Where("certificate_strategy_id=?", a.Strategy.ID).
		Order("id desc").
		Count(&certificates.Count).Offset(offset).Limit(limit).Find(&certificates.Items).Error
	if err != nil {
		return nil, err
	}

	return certificates.ToAny(), nil
}

func (a *AddressCertiOperator) InsertCertificates(items []any) error {
	results, err := new(CertificateConverter[*AddressCertificate]).ConvertSlice(items)
	if err != nil {
		return err
	}

	for i, c := range results {
		if c.Address == "" {
			return errors.Errorf("item %v missing address (index from 1)", i+1)
		}
		c.CertificateStrategyID = a.Strategy.ID
	}

	return models.GetDB().Save(&results).Error
}

func (a *AddressCertiOperator) DeleteCertificates(ids []uint) error {
	return models.GetDB().Where("certificate_strategy_id = ?", a.Strategy.ID).Where("id in (?)", ids).Delete(&AddressCertificate{}).Error
}
