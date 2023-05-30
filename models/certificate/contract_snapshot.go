package certificate

import "github.com/nft-rainbow/rainbow-app-service/models"

type ContractSnapshot struct {
	models.BaseModel
	Contract              string `json:"contract"`
	TokenId               string `json:"token_id"`
	Owner                 string `json:"owner"`
	ContractCertificateId uint   `json:"contract_certificate_id"`
}

func FindContractSnapshotByOwner(certiId uint, owner string) (snapshots []*ContractSnapshot, err error) {
	err = models.GetDB().Where("contract_certificate_id=? and owner=?", certiId, owner).Find(&snapshots).Error
	return
}
