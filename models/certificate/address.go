package certificate

import "github.com/nft-rainbow/rainbow-app-service/models"

type AddressCertificate struct {
	models.BaseModel
	Address               string `json:"address"`
	CertificateStrategyID uint
}
