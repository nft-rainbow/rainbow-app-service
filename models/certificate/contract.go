package certificate

import (
	"errors"

	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
)

type ContractCertificate struct {
	models.BaseModel
	ContractAddress       string `json:"contract_address"`
	SnapshotEpochNumber   uint64 `json:"snapshot_epoch_number"`
	SnapshotTaskStatus    uint   `json:"snapshot_task_status"`
	CertificateStrategyID uint   `json:"certificate_strategy_id"`
}

func FindContractCertificatesByStrategyId(id uint) (certis []*ContractCertificate, err error) {
	err = models.GetDB().Where("certificate_strategy_id=?", id).Find(&certis).Error
	return
}

// TODO:
// 1. gen snapshot task, 首次自动同步，
// 2. gen snapshot api, 如果失败使用api重新获取
// 3. query snapshot list

type ContractCertiChecker struct {
	Strategy *CertificateStrategy
}

func (a *ContractCertiChecker) CheckQualified(userAddress string) (bool, error) {
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
		snapShots, err := FindContractSnapshotByOwner(cert.ID, userAddress)
		if err != nil {
			return false, err
		}
		if len(snapShots) > 0 {
			return true, nil
		}
	}
	return false, nil
}

func (a *ContractCertiChecker) GetCertificates(offset int, limit int) (*Certificates, error) {
	var certificates Certificates
	err := models.GetDB().Model(&ContractCertificate{}).
		Where("certificate_strategy_id=?", a.Strategy.ID).
		Count(&certificates.Count).Offset(offset).Limit(limit).Find(&certificates.Items).Error
	if err != nil {
		return nil, err
	}
	return &certificates, nil
}
