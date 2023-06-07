package certificate

import "github.com/nft-rainbow/rainbow-app-service/models/enums"

type CertiOperator interface {
	CheckQualified(userAddress string) (bool, error)
	GetCertificates(offset int, limit int) (*CertificatesQueryResult[any], error)
	InsertCertificates([]any) error
	DeleteCertificates([]uint) error
}

func GetCertiOperator(cs *CertificateStrategy) CertiOperator {
	switch cs.CertificateType {
	case enums.CERTIFICATE_ADDRESS:
		return &AddressCertiOperator{cs}
	case enums.CERTIFICATE_PHONE:
		return &PhoneCertiOperator{cs}
	case enums.CERTIFICATE_DODO:
		return &DodoCertiOperator{cs}
	case enums.CERTIFICATE_CONTRACT:
		return &ContractCertiOperator{cs}
	case enums.CERTIFICATE_GASLESS:
		return &GaslessCertiOperator{ContractCertiOperator{cs}}
	}
	return nil
}
