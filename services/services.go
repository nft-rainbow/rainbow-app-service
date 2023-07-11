package services

import (
	"log"
	"os"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/spf13/viper"
)

var (
	cfxMainClient *sdk.Client
	cfxTestClient *sdk.Client
)

func Init() {
	InitConfluxChainClient()
	RunBatchMintTaskOnInit()
}

func InitConfluxChainClient() {
	option := sdk.ClientOption{Logger: os.Stdout}
	var err error
	cfxTestClient, err = sdk.NewClient(viper.GetString("rpcs.testnet"), option)
	if err != nil {
		log.Panic("Init Conflux testnet client failed", err)
	}
	cfxMainClient, err = sdk.NewClient(viper.GetString("rpcs.mainnet"), option)
	if err != nil {
		log.Panic("Init Conflux mainnet client failed", err)
	}
}

func GetConfluxClinet(chain enums.Chain) *sdk.Client {
	switch chain {
	case enums.CHAIN_CONFLUX:
		return cfxMainClient
	case enums.CHAIN_CONFLUX_TEST:
		return cfxTestClient
	}
	return nil
}
