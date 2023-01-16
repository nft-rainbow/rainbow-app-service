package services

import (
	"log"

	"math/big"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nft-rainbow/rainbow-app-service/contracts"
	"github.com/spf13/viper"
)

var cfxTestClient *sdk.Client
var cfxMainClient *sdk.Client

func InitConfluxChainClient() {
	option := sdk.ClientOption{}
	var err error
	cfxTestClient, err = sdk.NewClient(viper.GetString("rpcs.testnet"), option)
	if err != nil {
		log.Fatalln(err)
	}
	cfxMainClient, err = sdk.NewClient(viper.GetString("rpcs.mainnet"), option)
	if err != nil {
		log.Fatalln(err)
	}
}

func ERC1155BalanceOfBatch(address *cfxaddress.Address, accounts []*cfxaddress.Address, ids []*big.Int) ([]*big.Int, error) {
	var err error
	var nftCaller *contracts.ERC1155NFTCaller
	if address.GetNetworkID() == 1029 {
		nftCaller, err = contracts.NewERC1155NFTCaller(*address, cfxMainClient)
	} else {
		nftCaller, err = contracts.NewERC1155NFTCaller(*address, cfxTestClient)
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
