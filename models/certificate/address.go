package certificate

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
)

type AddressCertificate struct {
	models.BaseModel
	Address               string `json:"address"`
	CertificateStrategyID uint   `json:"certificate_strategy_id"`
}

type AddressCertiChecker struct {
	Strategy *CertificateStrategy
}

func (a *AddressCertiChecker) CheckQualified(userAddress string) (bool, error) {
	if a.Strategy == nil || a.Strategy.CertificateType != enums.CERTIFICATE_ADDRESS {
		return false, nil
	}
	var count int64
	if err := models.GetDB().Find("address=? and certificate_strategy_id=?", userAddress, a.Strategy.ID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a *AddressCertiChecker) GetCertificates(offset int, limit int) (*Certificates, error) {
	var certificates Certificates
	err := models.GetDB().Model(&AddressCertificate{}).
		Where("certificate_strategy_id=?", a.Strategy.ID).
		Count(&certificates.Count).Offset(offset).Limit(limit).Find(&certificates.Items).Error
	if err != nil {
		return nil, err
	}
	return &certificates, nil
}
