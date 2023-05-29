package certificate

import "github.com/nft-rainbow/rainbow-app-service/models"

type ContractCertificate struct {
	models.BaseModel
	ContractAddress     string `json:"contract_address"`
	SnapshotEpochNumber uint64 `json:"snapshot_epoch_number"`
}
