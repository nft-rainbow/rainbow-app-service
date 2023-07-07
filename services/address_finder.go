package services

import (
	"errors"

	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
)

type AddressFinder struct {
	SourceType enums.SourceType
}

func (a *AddressFinder) Find(sources []string) (exists map[string]string, unexists []string, err error) {
	switch a.SourceType {
	case enums.SOURCE_TYPE_ADDRESS:
		result := make(map[string]string)
		for _, source := range sources {
			result[source] = source
		}
		return result, nil, nil

	case enums.SOURCE_TYPE_PHONE:
		exists, err = models.FindTopWalletUsersByPhones(enums.WALLET_CELLAR, sources)
		if err != nil {
			return nil, nil, err
		}

		for _, phone := range sources {
			if _, ok := exists[phone]; !ok {
				unexists = append(unexists, phone)
			}
		}
		return exists, unexists, nil
	case enums.SOURCE_TYPE_DODO:
		return nil, nil, errors.New("unsupported dodo source type")
	}
	return nil, nil, errors.New("unsupported source type")
}
