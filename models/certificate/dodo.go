package certificate

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils/gormutils"
	"github.com/pkg/errors"
)

type DodoCertificate struct {
	models.BaseModel
	DodoSourceId          string `json:"dodo_source_id"`
	CertificateStrategyID uint   `json:"certificate_strategy_id"`
}

type DodoCertiOperator struct {
	Strategy *CertificateStrategy
}

func (a *DodoCertiOperator) CheckQualified(userAddress string) (bool, error) {
	if a.Strategy == nil || a.Strategy.CertificateType != enums.CERTIFICATE_DODO {
		return false, nil
	}

	su, err := models.FindSocialUserByAddress(userAddress)
	if err != nil {
		if gormutils.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	var count int64
	if err := models.GetDB().Model(&DodoCertificate{}).Find("dodo_source_id = ? and certificate_strategy_id=?", su.UserId, a.Strategy.ID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a *DodoCertiOperator) GetCertificates(offset int, limit int) (*CertificatesQueryResult[any], error) {
	certificates := CertificatesQueryResult[*DodoCertificate]{
		CertificateType: enums.CERTIFICATE_DODO,
	}

	err := models.GetDB().Model(&DodoCertificate{}).
		Where("certificate_strategy_id=?", a.Strategy.ID).
		Count(&certificates.Count).Offset(offset).Limit(limit).Find(&certificates.Items).Error
	if err != nil {
		return nil, err
	}
	return certificates.ToAny(), nil
}

func (a *DodoCertiOperator) InsertCertificates(items []any) error {
	results, err := new(CertificateConverter[*DodoCertificate]).ConvertSlice(items)
	if err != nil {
		return err
	}

	for i, c := range results {
		if c.DodoSourceId == "" {
			return errors.Errorf("item %v missing address (index from 1)", i+1)
		}
		c.CertificateStrategyID = a.Strategy.ID
	}

	return models.GetDB().Save(&results).Error
}
