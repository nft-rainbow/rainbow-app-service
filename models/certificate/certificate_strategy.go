package certificate

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
)

type CertificateStrategy struct {
	models.BaseModel
	CertificateType enums.CertificateType `json:"certificate_type" swaggertype:"string"`
	Addresses       []AddressCertificate  `json:"addresses,omitempty"`
	Phones          []PhoneCertificate    `json:"phones,omitempty"`
	Dodos           []DodoCertificate     `json:"dodos,omitempty"`
	Contracts       []ContractCertificate `json:"contracts,omitempty"`
	Gaslesses       []GaslessCertificate  `json:"gaslesses,omitempty"`
}

type Certificates struct {
	Items           []any                 `json:"items,omitempty"`
	Count           int64                 `json:"count,omitempty"`
	CertificateType enums.CertificateType `json:"certificate_type" swaggertype:"string"`
}

type CertiOperator interface {
	CheckQualified(userAddress string) (bool, error)
	GetCertificates(offset int, limit int) (*Certificates, error)
}

func GetCertiChecker(cs *CertificateStrategy) CertiOperator {
	switch cs.CertificateType {
	case enums.CERTIFICATE_ADDRESS:
		return &AddressCertiChecker{cs}
	}
	return nil
}

func (cs *CertificateStrategy) CheckQualified(userAddress string) (bool, error) {
	return GetCertiChecker(cs).CheckQualified(userAddress)
}

func (cs *CertificateStrategy) GetCertificates(offset int, limit int) (*Certificates, error) {
	return GetCertiChecker(cs).GetCertificates(offset, limit)
}
