package utils

import "fmt"

type ChainType uint
type ChainID uint
type ContractType uint

// chain types
const (
	CHAIN_TYPE_CFX ChainType = iota + 1
	CHAIN_TYPE_ETH
)

const (
	CONFLUX_MAINNET_ID ChainID = 1029
	CONFLUX_TEST_ID    ChainID = 1
)

// contract types
const (
	CONTRACT_TYPE_ERC721 ContractType = iota + 1
	CONTRACT_TYPE_ERC1155
)

// contract type names
const ERC721 = "erc721"
const ERC1155 = "erc1155"

const CONFLUX_TEST = "conflux_test"
const CONFLUX = "conflux"

func ChainInfoByName(name string) (ChainType, ChainID, error) {
	switch name {
	case CONFLUX_TEST:
		return CHAIN_TYPE_CFX, 1, nil
	case CONFLUX:
		return CHAIN_TYPE_CFX, 1029, nil
	default:
		return 0, 0, fmt.Errorf("unknown chain name: %s", name)
	}
}

func ContractTypeByName(name string) (ContractType, error) {
	switch name {
	case ERC721:
		return CONTRACT_TYPE_ERC721, nil
	case ERC1155:
		return CONTRACT_TYPE_ERC1155, nil
	default:
		return 0, fmt.Errorf("unknown contract type: %s", name)
	}
}
