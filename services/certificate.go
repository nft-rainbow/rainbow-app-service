package services

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/certificate"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
)

type InsertCertificateStrategyReq[T any] struct {
	CertificateType enums.CertificateType `json:"certificate_type" swaggertype:"string"`
	Items           []T                   `json:"items"`
}

func InsertCertificateStrategy(req *InsertCertificateStrategyReq[any]) (*certificate.CertificateStrategy, error) {
	// err := models.GetDB().Transaction(func(tx *gorm.DB) error {
	cs := certificate.CertificateStrategy{
		CertificateType: req.CertificateType,
	}
	if err := models.GetDB().Save(&cs).Error; err != nil {
		return nil, err
	}
	if err := certificate.GetCertiOperator(&cs).InsertCertificates(req.Items); err != nil {
		models.GetDB().Delete(&cs)
		return nil, err
	}
	return &cs, nil
}

func GetCertificates(certificateStratageId uint, offset, limit int) (*certificate.CertificatesQueryResult[any], error) {
	cs, err := certificate.FindCertificateStrategyById(certificateStratageId)
	if err != nil {
		return nil, err
	}
	return cs.GetCertificates(offset, limit)
}
