package certificate

import (
	"github.com/nft-rainbow/rainbow-app-service/models"
)

func Init() {
	err := models.GetDB().AutoMigrate(
		&CertificateStrategy{},
		&AddressCertificate{},
		&PhoneCertificate{},
		&DodoCertificate{},
		&ContractCertificate{},
		&ContractSnapshot{},
	)
	if err != nil {
		panic(err)
	}
}
