package certificate

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
)

type ContractSnapshot struct {
	models.BaseModel
	Contract              string `json:"contract"`
	ActivityCode          string `json:"activity_code"`
	SnapshotEpochNumber   uint64 `json:"snapshot_epoch_number"`
	TokenId               string `json:"token_id"`
	Owner                 string `json:"owner"`
	ContractCertificateId uint   `json:"contract_certificate_id"`
}

func FindContractSnapshotsByOwner(certiId uint, owner string) (snapshots []*ContractSnapshot, err error) {
	err = models.GetDB().Where("contract_certificate_id=? and owner=?", certiId, owner).Find(&snapshots).Error
	return
}

func FindContractSnapshots(certiId uint, offset, limit int) (snapshotStatus enums.SnapshotStatus, count int64, snapshots []*ContractSnapshot, err error) {
	contractCerti, err := FindContractCertificateById(certiId)
	if err != nil {
		return enums.SNAPSHOT_STATUS_INIT, 0, nil, err
	}
	snapshotStatus = contractCerti.SnapshotTaskStatus
	err = models.GetDB().Model(&ContractSnapshot{}).Where("contract_certificate_id=? ", certiId).Count(&count).Offset(offset).Limit(limit).Find(&snapshots).Error
	return
}
