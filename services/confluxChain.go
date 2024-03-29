package services

import (
	"math/big"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nft-rainbow/rainbow-app-service/contracts"
)

func ERC1155BalanceOfBatch(contractAddr *cfxaddress.Address, accounts []*cfxaddress.Address, ids []*big.Int) ([]*big.Int, error) {
	var err error
	var nftCaller *contracts.ERC1155NFTCaller
	if contractAddr.GetNetworkID() == 1029 {
		nftCaller, err = contracts.NewERC1155NFTCaller(*contractAddr, cfxMainClient)
	} else {
		nftCaller, err = contracts.NewERC1155NFTCaller(*contractAddr, cfxTestClient)
	}
	if err != nil {
		return nil, err
	}

	commonAddresses := []common.Address{}
	for _, account := range accounts {
		commonAddresses = append(commonAddresses, account.MustGetCommonAddress())
	}

	bigBalance, err := nftCaller.BalanceOfBatch(nil, commonAddresses, ids)
	return bigBalance, err
}

func CommonBalanceOfBatch(contract, user string) ([]*big.Int, error) {
	contractAddress := cfxaddress.MustNewFromBase32(contract)
	userAddress := cfxaddress.MustNewFromBase32(user)
	users := make([]*cfxaddress.Address, 5)

	for i := range users {
		users[i] = &userAddress
	}

	ids := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(5)}

	resp, err := ERC1155BalanceOfBatch(&contractAddress, users, ids)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
