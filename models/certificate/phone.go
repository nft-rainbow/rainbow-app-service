package certificate

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils/gormutils"
	"github.com/pkg/errors"
)

type PhoneCertificateInsertPart struct {
	Phone string `json:"phone"`
}
type PhoneCertificate struct {
	models.BaseModel
	PhoneCertificateInsertPart
	CertificateStrategyID uint `json:"certificate_strategy_id"`
}

type PhoneCertiOperator struct {
	Strategy *CertificateStrategy
}

func (a *PhoneCertiOperator) CheckQualified(userAddress string) (bool, error) {
	if a.Strategy == nil || a.Strategy.CertificateType != enums.CERTIFICATE_PHONE {
		return false, nil
	}

	wu, err := models.FindWalletUserByAddress(userAddress)
	if err != nil {
		if gormutils.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	var count int64
	if err := models.GetDB().Model(&PhoneCertificate{}).Find("phone = ? and certificate_strategy_id=?", wu.Phone, a.Strategy.ID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a *PhoneCertiOperator) GetCertificates(offset int, limit int) (*CertificatesQueryResult[any], error) {
	certificates := CertificatesQueryResult[*PhoneCertificate]{
		CertificateType: enums.CERTIFICATE_PHONE,
	}

	err := models.GetDB().Model(&PhoneCertificate{}).
		Where("certificate_strategy_id=?", a.Strategy.ID).
		Count(&certificates.Count).Offset(offset).Limit(limit).Find(&certificates.Items).Error
	if err != nil {
		return nil, err
	}
	return certificates.ToAny(), nil
}

func (a *PhoneCertiOperator) InsertCertificates(items []any) error {
	results, err := new(CertificateConverter[*PhoneCertificate]).ConvertSlice(items)
	if err != nil {
		return err
	}

	for i, c := range results {
		if c.Phone == "" {
			return errors.Errorf("item %v missing phone (index from 1)", i+1)
		}
		c.CertificateStrategyID = a.Strategy.ID
	}

	return models.GetDB().Save(&results).Error
}

func (a *PhoneCertiOperator) DeleteCertificates(ids []uint) error {
	return models.GetDB().Where("certificate_strategy_id = ?", a.Strategy.ID).Where("id in (?)", ids).Delete(&PhoneCertificate{}).Error
}
