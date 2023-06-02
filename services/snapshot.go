package services

import (
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/nft-rainbow/rainbow-app-service/contracts"
	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/certificate"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/pkg/errors"
)

type SnapshotService struct {
}

func (s *SnapshotService) Start(certi *certificate.ContractCertificate) {
	certi.SnapshotTaskStatus = enums.SNAPSHOT_STATUS_PROCESSING

	var err error
	switch certi.ContractType {
	case enums.ContractType_ERC721:
		err = s.snapshotErc721(certi)
	case enums.ContractType_ERC1155:
		err = s.snapshotErc1155(certi)
	}

	if err != nil {
		certi.SnapshotTaskStatus = enums.SNAPSHOT_STATUS_SUCCESS
		certi.SnapshotProcessError = err.Error()
	} else {
		certi.SnapshotTaskStatus = enums.SNAPSHOT_STATUS_FAILED
	}
}

func (s *SnapshotService) snapshotErc721(certi *certificate.ContractCertificate) error {
	option := &bind.CallOpts{
		EpochNumber: types.NewEpochNumberUint64(certi.SnapshotEpochNumber),
	}
	erc721, err := contracts.NewERC721NFTCustom(cfxaddress.MustNew(certi.ContractAddress), GetConfluxClinet(certi.Chain))
	if err != nil {
		return err
	}
	total, err := erc721.TotalSupply(option)
	if err != nil {
		return errors.WithMessage(err, "failed to get totalSupply")
	}

	delta := 100
	for i := certi.SnapshotProcessingIndex; i < int(total.Int64()); i = i + delta {
		results, err := erc721.Tokens(option, big.NewInt(int64(i)), big.NewInt(int64(delta)))
		if err != nil {
			return errors.WithMessagef(err, "failed to get tokens %v ~ %v", i, i+delta)
		}
		// get owner
		for j, t := range results.TokenIds {
			owner, err := erc721.OwnerOf(option, t)
			if err != nil {
				return errors.WithMessagef(err, "failed to get owner of %v", t)
			}
			if err := models.GetDB().Save(&certificate.ContractSnapshot{
				ContractCertificateId: certi.ID,
				TokenId:               t.String(),
				Owner:                 owner.String(),
			}).Error; err != nil {
				return errors.WithMessage(err, "failed to save snapshot")
			}

			certi.SnapshotProcessingIndex = i + j
			if err := models.GetDB().Save(certi).Error; err != nil {
				return errors.WithMessage(err, "failed to update snapshot processing index")
			}
		}
	}
	return nil
}

func (s *SnapshotService) snapshotErc1155(certi *certificate.ContractCertificate) error {
	return errors.New("not implemented")
}
