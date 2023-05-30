package certificate

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils/gormutils"
)

type DodoCertificate struct {
	models.BaseModel
	DodoSourceId          string `json:"dodo_source_id"`
	CertificateStrategyID uint   `json:"certificate_strategy_id"`
}

type DodoCertiChecker struct {
	Strategy *CertificateStrategy
}

func (a *DodoCertiChecker) CheckQualified(userAddress string) (bool, error) {
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

func (a *DodoCertiChecker) GetCertificates(offset int, limit int) (*Certificates, error) {
	var certificates Certificates
	err := models.GetDB().Model(&DodoCertificate{}).
		Where("certificate_strategy_id=?", a.Strategy.ID).
		Count(&certificates.Count).Offset(offset).Limit(limit).Find(&certificates.Items).Error
	if err != nil {
		return nil, err
	}
	return &certificates, nil
}
