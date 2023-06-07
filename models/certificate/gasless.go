package certificate

import (
	"github.com/nft-rainbow/rainbow-app-service/config"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/pkg/errors"
)

type GaslessCertificate struct {
	ContractCertificate
}

type GaslessCertiOperator struct {
	ContractCertiOperator
}

func (a *GaslessCertiOperator) GetCertificates(offset int, limit int) (*CertificatesQueryResult[any], error) {
	certificates, err := a.ContractCertiOperator.GetCertificates(offset, limit)
	if err != nil {
		return nil, err
	}
	certificates.CertificateType = enums.CERTIFICATE_GASLESS
	return certificates, nil
}

func (a *GaslessCertiOperator) InsertCertificates(items []any) error {
	converter := new(CertificateConverter[*ContractCertificate])
	results, err := converter.ConvertSlice(items)
	if err != nil {
		return err
	}

	for i, c := range results {
		if c.ActivityCode == "" {
			return errors.Errorf("item %v missing activity_code", i+1)
		}

		gaslessContractRawId := config.GetGaslessContractIdByChain(c.Chain)
		contract, err := models.FindContractByRawId(gaslessContractRawId)
		if err != nil {
			return err
		}
		c.ContractAddress = contract.ContractAddress
		c.ContractType = enums.ContractType_ERC721
	}

	return a.ContractCertiOperator.InsertCertificates(converter.ConvertSliceBack(results))
}
