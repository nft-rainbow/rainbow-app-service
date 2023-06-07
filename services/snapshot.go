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
	"github.com/sirupsen/logrus"
)

type SnapshotService struct {
}

func (s *SnapshotService) Start(certi *certificate.ContractCertificate) error {
	if certi.SnapshotTaskStatus == enums.SNAPSHOT_STATUS_PROCESSING {
		return errors.New("the contract snapshot is running")
	}

	go func() {
		certi.SnapshotTaskStatus = enums.SNAPSHOT_STATUS_PROCESSING
		certi.SnapshotProcessError = ""
		if err := models.GetDB().Save(certi).Error; err != nil {
			logrus.WithField("certificate", certi).Error("failed to update certificate")
		}

		var err error
		switch certi.ContractType {
		case enums.ContractType_ERC721:
			err = s.snapshotErc721(certi)
		case enums.ContractType_ERC1155:
			err = s.snapshotErc1155(certi)
		}

		if err != nil {
			certi.SnapshotTaskStatus = enums.SNAPSHOT_STATUS_FAILED
			certi.SnapshotProcessError = err.Error()
		} else {
			certi.SnapshotTaskStatus = enums.SNAPSHOT_STATUS_SUCCESS
		}
		logrus.WithError(err).WithField("certificate", certi).Info("get snapshot done")

		if err = models.GetDB().Save(certi).Error; err != nil {
			logrus.WithField("certificate", certi).Error("failed to save certificate")
		}
	}()
	return nil
}

func (s *SnapshotService) snapshotErc721(certi *certificate.ContractCertificate) error {
	option := &bind.CallOpts{
		EpochNumber: types.NewEpochNumberUint64(certi.SnapshotEpochNumber),
	}
	erc721, err := contracts.NewERC721NFTCustom(cfxaddress.MustNew(certi.ContractAddress), GetConfluxClinet(certi.Chain))
	if err != nil {
		return err
	}

	getTokenIdCount := func() (int64, error) {
		if certi.ActivityCode == "" {
			total, err := erc721.TotalSupply(option)
			if err != nil {
				return 0, errors.WithMessage(err, "failed to get totalSupply")
			}
			return total.Int64(), nil
		}
		return models.CountPOAPResult(certi.ActivityCode, &models.POAPResultFilter{Statuses: []enums.TransactionStatus{
			enums.TRANSACTION_STATUS_SUCCESS,
		}})
	}

	getRangeTokens := func(startIdx int, length int) (tokenIds []*big.Int, err error) {
		if certi.ActivityCode == "" {
			results, err := erc721.Tokens(option, big.NewInt(int64(startIdx)), big.NewInt(int64(length)))
			if err != nil {
				return nil, err
			}
			return results.TokenIds, nil
		}

		var poapResults []*models.POAPResult
		if err = models.GetDB().Where("activity_code=? and status=?", certi.ActivityCode, enums.TRANSACTION_STATUS_SUCCESS).
			Order("id DESC").Offset(startIdx).Limit(length).Find(&poapResults).Error; err != nil {
			return nil, err
		}

		for _, v := range poapResults {
			// tokenId, err := strconv.Atoi(v.TokenID)
			tokenId, ok := new(big.Int).SetString(v.TokenID, 0)
			if !ok {
				return nil, errors.Errorf("unkown token id %v", v.TokenID)
			}
			tokenIds = append(tokenIds, tokenId)
		}
		return
	}

	totalCount, err := getTokenIdCount()
	if err != nil {
		return err
	}
	certi.RelatedTokenCount = uint(totalCount)

	delta := 100
	for i := certi.SnapshotProcessingIndex + 1; i < int(totalCount); i = i + delta {
		// results, err := erc721.Tokens(option, big.NewInt(int64(i)), big.NewInt(int64(delta)))
		results, err := getRangeTokens(i, delta)
		if err != nil {
			return errors.WithMessagef(err, "failed to get tokens %v ~ %v", i, i+delta)
		}
		// get owner
		for j, t := range results {
			owner, err := erc721.OwnerOf(option, t)
			if err != nil {
				return errors.WithMessagef(err, "failed to get owner of %v", t)
			}

			if err := models.GetDB().Save(&certificate.ContractSnapshot{
				ContractCertificateId: certi.ID,
				Contract:              certi.ContractAddress,
				ActivityCode:          certi.ActivityCode,
				SnapshotEpochNumber:   certi.SnapshotEpochNumber,
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

type ContractSnapshotResp struct {
	SnapshotStatus enums.SnapshotStatus
	Count          int64
	Snapshots      []*certificate.ContractSnapshot
}

func (s *SnapshotService) GetContractSnapshots(certificateId uint, offsert, limit int) (*ContractSnapshotResp, error) {
	status, count, snapshots, err := certificate.FindContractSnapshots(certificateId, offsert, limit)
	if err != nil {
		return nil, err
	}
	return &ContractSnapshotResp{
		SnapshotStatus: status,
		Count:          count,
		Snapshots:      snapshots,
	}, nil
}
