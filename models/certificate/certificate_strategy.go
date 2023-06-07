package certificate

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
)

type CertificateStrategy struct {
	models.BaseModel
	CertificateType enums.CertificateType `json:"certificate_type" swaggertype:"string"`
	// Addresses       []AddressCertificate  `json:"addresses,omitempty"`
	// Phones          []PhoneCertificate    `json:"phones,omitempty"`
	// Dodos           []DodoCertificate     `json:"dodos,omitempty"`
	// Contracts       []ContractCertificate `json:"contracts,omitempty"`
	// Gaslesses       []GaslessCertificate  `json:"gaslesses,omitempty"`
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
