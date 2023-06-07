package certificate

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/pkg/errors"
)

type ContractCertificateInsertPart struct {
	GaslessCertificateInsertPart
	ContractAddress string             `json:"contract_address"`
	ContractType    enums.ContractType `json:"contract_type"`
}
type ContractCertificate struct {
	models.BaseModel
	ContractCertificateInsertPart
	SnapshotTaskStatus      enums.SnapshotStatus `json:"snapshot_task_status" swaggertype:"string"`
	RelatedTokenCount       uint                 `json:"related_token_count"`
	SnapshotProcessingIndex int                  `gorm:"default:-1" json:"snapshot_processing_index"`
	SnapshotProcessError    string               `json:"snapshot_processing_error"`
	CertificateStrategyID   uint                 `json:"certificate_strategy_id"`
}

func FindContractCertificatesByStrategyId(id uint) (certis []*ContractCertificate, err error) {
	err = models.GetDB().Where("certificate_strategy_id=?", id).Find(&certis).Error
	return
}

func FindContractCertificateById(id uint) (*ContractCertificate, error) {
	var certi ContractCertificate
	err := models.GetDB().First(&certi, id).Error
	if err != nil {
		return nil, errors.WithMessage(err, "failed find contract certificate")
	}
	return &certi, nil
}

func (c *ContractCertificate) FindSnapshots() (snapshots []*ContractSnapshot, err error) {
	err = models.GetDB().Where("contract_certificate_id=?", c.ID).Find(&snapshots).Error
	return
}

type ContractCertiOperator struct {
	Strategy *CertificateStrategy
}

func (a *ContractCertiOperator) CheckQualified(userAddress string) (bool, error) {
	if a.Strategy == nil || a.Strategy.CertificateType != enums.CERTIFICATE_CONTRACT {
		return false, nil
	}

	certis, err := FindContractCertificatesByStrategyId(a.Strategy.ID)
	if err != nil {
		return false, err
	}

	for _, cert := range certis {
		if cert.SnapshotTaskStatus == 0 {
			return false, errors.New("need generate snapshot first")
		}
		if cert.SnapshotTaskStatus == 2 {
			return false, errors.New("snapshot generation failed, please retry")
		}
	}

	for _, cert := range certis {
		snapShots, err := FindContractSnapshotsByOwner(cert.ID, userAddress)
		if err != nil {
			return false, err
		}
		if len(snapShots) > 0 {
			return true, nil
		}
	}
	return false, nil
}

func (a *ContractCertiOperator) GetCertificates(offset int, limit int) (*CertificatesQueryResult[any], error) {
	certificates := CertificatesQueryResult[*ContractCertificate]{
		CertificateType: enums.CERTIFICATE_CONTRACT,
	}
	err := models.GetDB().Model(&ContractCertificate{}).
		Where("certificate_strategy_id=?", a.Strategy.ID).
		Count(&certificates.Count).Offset(offset).Limit(limit).Find(&certificates.Items).Error
	if err != nil {
		return nil, err
	}
	return certificates.ToAny(), nil
}

func (a *ContractCertiOperator) InsertCertificates(items []any) error {
	results, err := new(CertificateConverter[*ContractCertificate]).ConvertSlice(items)
	if err != nil {
		return err
	}

	for i, c := range results {
		if c.ContractAddress == "" {
			return errors.Errorf("item %v missing address (index from 1)", i+1)
		}
		if c.ContractType == enums.ContractType(0) {
			return errors.Errorf("item %v missing contract type (index from 1)", i+1)
		}
		if c.SnapshotEpochNumber == 0 {
			return errors.Errorf("item %v missing snapshot epoch number (index from 1)", i+1)
		}

		// check activity match contract
		if c.ActivityCode != "" {
			activity, err := models.FindActivityByCode(c.ActivityCode)
			if err != nil {
				return errors.WithMessagef(err, " item %v: failed to find activity with code %v", i+1, c.ActivityCode)
			}
			contract, err := models.FindContractByRawId(uint(*activity.ContractRawID))
			if err != nil {
				return errors.WithMessagef(err, " item %v: failed to find contract with contract id %v", i+1, *activity.ContractRawID)
			}
			if contract.ContractAddress != c.ContractAddress {
				return errors.Errorf("the contract address not match the activity")
			}
		}

		c.CertificateStrategyID = a.Strategy.ID
	}

	return models.GetDB().Save(&results).Error
}

func (a *ContractCertiOperator) DeleteCertificates(ids []uint) error {
	return models.GetDB().Where("certificate_strategy_id = ?", a.Strategy.ID).Where("id in (?)", ids).Delete(&ContractCertificate{}).Error
}
