package services

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/certificate"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/pkg/errors"
)

type InsertCertificateStrategyReq[T any] struct {
	CertificateType enums.CertificateType `json:"certificate_type" swaggertype:"string" swaggerignore:"true"`
	Name            string                `json:"name"`
	Items           []T                   `json:"items"`
}

func InsertCertificateStrategy(req *InsertCertificateStrategyReq[any], userId uint) (*certificate.CertificateStrategy, error) {
	// check name exists

	var count int64
	if err := models.GetDB().Model(&certificate.CertificateStrategy{}).
		Where("name=? and user_id=?", req.Name, userId).Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, errors.New("same name exists")
	}

	cs := certificate.CertificateStrategy{
		CertificateStrategyCore: certificate.CertificateStrategyCore{
			CertificateType: req.CertificateType,
			UserId:          userId,
			Name:            req.Name,
		},
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

func GetCertificates(certificateStrategyId uint, offset, limit int) (*certificate.CertificatesQueryResult[any], error) {
	cs, err := certificate.FindCertificateStrategyById(certificateStrategyId)
	if err != nil {
		return nil, err
	}
	return cs.GetCertificates(offset, limit)
}
