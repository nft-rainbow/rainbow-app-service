package certificate

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/nft-rainbow/rainbow-app-service/utils/gormutils"
)

type PhoneCertificate struct {
	models.BaseModel
	Phone                 string `json:"phone"`
	CertificateStrategyID uint   `json:"certificate_strategy_id"`
}

type PhoneCertiChecker struct {
	Strategy *CertificateStrategy
}

func (a *PhoneCertiChecker) CheckQualified(userAddress string) (bool, error) {
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

func (a *PhoneCertiChecker) GetCertificates(offset int, limit int) (*Certificates, error) {
	var certificates Certificates
	err := models.GetDB().Model(&PhoneCertificate{}).
		Where("certificate_strategy_id=?", a.Strategy.ID).
		Count(&certificates.Count).Offset(offset).Limit(limit).Find(&certificates.Items).Error
	if err != nil {
		return nil, err
	}
	return &certificates, nil
}
