// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/cfxclient/bulk"

	types "github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = ethBind.Bind
	_ = common.Big1
	_ = ethtypes.BloomLookup
	_ = event.NewSubscription
)

// ERC1155NFTCustomABI is the input ABI used to generate the binding from.
const ERC1155NFTCustomABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"PermanentURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PermanentURIGlobal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltiesBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"royaltiesAddress\",\"type\":\"address\"}],\"name\":\"RoyaltyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"transferableByAdmin\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"transferableByUser\",\"type\":\"bool\"}],\"name\":\"TransferableChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"whites\",\"type\":\"address[]\"}],\"name\":\"addSponsorPrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freezeGlobalMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"freezeTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"freezeTokenUris\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"grantAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"grantMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"royaltiesBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"royaltiesAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"tokensTransferableByAdmin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"tokensTransferableByUser\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSetSponsorWhitelistForAllUser\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listSponsorPrivilege\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataUpdatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tos\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"}],\"name\":\"mintToBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"whites\",\"type\":\"address[]\"}],\"name\":\"removeSponsorPrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltiesAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltiesBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_royaltiesBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_royaltiesAddress\",\"type\":\"address\"}],\"name\":\"setRoyalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"transferableByAdmin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"transferableByUser\",\"type\":\"bool\"}],\"name\":\"setTokensTransferable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokenCountOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensBurnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"tokensOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensTransferableByAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensTransferableByUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"transferBatchByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"updateTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC1155NFTCustom is an auto generated Go binding around an Conflux contract.
type ERC1155NFTCustom struct {
	ERC1155NFTCustomCaller     // Read-only binding to the contract
	ERC1155NFTCustomTransactor // Write-only binding to the contract
	ERC1155NFTCustomFilterer   // Log filterer for contract events
}

// ERC1155NFTCustomCaller is an auto generated read-only Go binding around an Conflux contract.
type ERC1155NFTCustomCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155NFTCustomBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type ERC1155NFTCustomBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155NFTCustomTransactor is an auto generated write-only Go binding around an Conflux contract.
type ERC1155NFTCustomTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155NFTCustomBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type ERC1155NFTCustomBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155NFTCustomFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type ERC1155NFTCustomFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155NFTCustomSession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type ERC1155NFTCustomSession struct {
	Contract     *ERC1155NFTCustom // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1155NFTCustomCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type ERC1155NFTCustomCallerSession struct {
	Contract *ERC1155NFTCustomCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC1155NFTCustomTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type ERC1155NFTCustomTransactorSession struct {
	Contract     *ERC1155NFTCustomTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC1155NFTCustomRaw is an auto generated low-level Go binding around an Conflux contract.
type ERC1155NFTCustomRaw struct {
	Contract *ERC1155NFTCustom // Generic contract binding to access the raw methods on
}

// ERC1155NFTCustomCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type ERC1155NFTCustomCallerRaw struct {
	Contract *ERC1155NFTCustomCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1155NFTCustomTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type ERC1155NFTCustomTransactorRaw struct {
	Contract *ERC1155NFTCustomTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155NFTCustom creates a new instance of ERC1155NFTCustom, bound to a specific deployed contract.
func NewERC1155NFTCustom(address types.Address, backend bind.ContractBackend) (*ERC1155NFTCustom, error) {
	contract, err := bindERC1155NFTCustom(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustom{ERC1155NFTCustomCaller: ERC1155NFTCustomCaller{contract: contract}, ERC1155NFTCustomTransactor: ERC1155NFTCustomTransactor{contract: contract}, ERC1155NFTCustomFilterer: ERC1155NFTCustomFilterer{contract: contract}}, nil
}

// NewERC1155NFTCustomCaller creates a new read-only instance of ERC1155NFTCustom, bound to a specific deployed contract.
func NewERC1155NFTCustomCaller(address types.Address, caller bind.ContractCaller) (*ERC1155NFTCustomCaller, error) {
	contract, err := bindERC1155NFTCustom(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomCaller{contract: contract}, nil
}

// NewERC1155NFTCustomTransactor creates a new write-only instance of ERC1155NFTCustom, bound to a specific deployed contract.
func NewERC1155NFTCustomTransactor(address types.Address, transactor bind.ContractTransactor) (*ERC1155NFTCustomTransactor, error) {
	contract, err := bindERC1155NFTCustom(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomTransactor{contract: contract}, nil
}

// NewERC1155NFTCustomFilterer creates a new log filterer instance of ERC1155NFTCustom, bound to a specific deployed contract.
func NewERC1155NFTCustomFilterer(address types.Address, filterer bind.ContractFilterer) (*ERC1155NFTCustomFilterer, error) {
	contract, err := bindERC1155NFTCustom(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomFilterer{contract: contract}, nil
}

// NewERC1155NFTCustomCaller creates a new read-only instance of ERC1155NFTCustom, bound to a specific deployed contract.
func NewERC1155NFTCustomBulkCaller(address types.Address, caller bind.ContractCaller) (*ERC1155NFTCustomBulkCaller, error) {
	contract, err := bindERC1155NFTCustom(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomBulkCaller{contract: contract}, nil
}

// NewERC1155NFTCustomBulkTransactor creates a new write-only instance of ERC1155NFTCustom, bound to a specific deployed contract.
func NewERC1155NFTCustomBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*ERC1155NFTCustomBulkTransactor, error) {
	contract, err := bindERC1155NFTCustom(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomBulkTransactor{contract: contract}, nil
}

// bindERC1155NFTCustom binds a generic wrapper to an already deployed contract.
func bindERC1155NFTCustom(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1155NFTCustomABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155NFTCustom *ERC1155NFTCustomRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155NFTCustom.Contract.ERC1155NFTCustomCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155NFTCustom *ERC1155NFTCustomRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.ERC1155NFTCustomTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155NFTCustom *ERC1155NFTCustomRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.ERC1155NFTCustomTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155NFTCustom.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) DEFAULTADMINROLE(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "DEFAULT_ADMIN_ROLE")

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "DEFAULT_ADMIN_ROLE")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC1155NFTCustom.Contract.DEFAULTADMINROLE(&_ERC1155NFTCustom.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC1155NFTCustom.Contract.DEFAULTADMINROLE(&_ERC1155NFTCustom.CallOpts)
}

// ZERO is a free data retrieval call binding the contract method 0x58fa63ca.
//
// Solidity: function ZERO() view returns(address)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) ZERO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "ZERO")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// ZERO is a free data retrieval call binding the contract method 0x58fa63ca.
//
// Solidity: function ZERO() view returns(address)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) ZERO(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "ZERO")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "ZERO")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// ZERO is a free data retrieval call binding the contract method 0x58fa63ca.
//
// Solidity: function ZERO() view returns(address)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) ZERO() (common.Address, error) {
	return _ERC1155NFTCustom.Contract.ZERO(&_ERC1155NFTCustom.CallOpts)
}

// ZERO is a free data retrieval call binding the contract method 0x58fa63ca.
//
// Solidity: function ZERO() view returns(address)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) ZERO() (common.Address, error) {
	return _ERC1155NFTCustom.Contract.ZERO(&_ERC1155NFTCustom.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "balanceOf", account, id)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) BalanceOf(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, account common.Address, id *big.Int) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "balanceOf", account, id)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "balanceOf")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155NFTCustom.Contract.BalanceOf(&_ERC1155NFTCustom.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155NFTCustom.Contract.BalanceOf(&_ERC1155NFTCustom.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if __err != nil {
		return *new([]*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, __err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) BalanceOfBatch(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) (*[]*big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "balanceOfBatch", accounts, ids)

	out0 := new([]*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "balanceOfBatch")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ERC1155NFTCustom.Contract.BalanceOfBatch(&_ERC1155NFTCustom.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ERC1155NFTCustom.Contract.BalanceOfBatch(&_ERC1155NFTCustom.CallOpts, accounts, ids)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "contractURI")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) ContractURI(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "contractURI")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "contractURI")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) ContractURI() (string, error) {
	return _ERC1155NFTCustom.Contract.ContractURI(&_ERC1155NFTCustom.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) ContractURI() (string, error) {
	return _ERC1155NFTCustom.Contract.ContractURI(&_ERC1155NFTCustom.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) Exists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "exists", tokenId)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) Exists(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, tokenId *big.Int) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "exists", tokenId)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "exists")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) Exists(tokenId *big.Int) (bool, error) {
	return _ERC1155NFTCustom.Contract.Exists(&_ERC1155NFTCustom.CallOpts, tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) Exists(tokenId *big.Int) (bool, error) {
	return _ERC1155NFTCustom.Contract.Exists(&_ERC1155NFTCustom.CallOpts, tokenId)
}

// FreezeTokenUris is a free data retrieval call binding the contract method 0x8d010db3.
//
// Solidity: function freezeTokenUris(uint256 ) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) FreezeTokenUris(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "freezeTokenUris", arg0)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// FreezeTokenUris is a free data retrieval call binding the contract method 0x8d010db3.
//
// Solidity: function freezeTokenUris(uint256 ) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) FreezeTokenUris(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, arg0 *big.Int) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "freezeTokenUris", arg0)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "freezeTokenUris")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// FreezeTokenUris is a free data retrieval call binding the contract method 0x8d010db3.
//
// Solidity: function freezeTokenUris(uint256 ) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) FreezeTokenUris(arg0 *big.Int) (bool, error) {
	return _ERC1155NFTCustom.Contract.FreezeTokenUris(&_ERC1155NFTCustom.CallOpts, arg0)
}

// FreezeTokenUris is a free data retrieval call binding the contract method 0x8d010db3.
//
// Solidity: function freezeTokenUris(uint256 ) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) FreezeTokenUris(arg0 *big.Int) (bool, error) {
	return _ERC1155NFTCustom.Contract.FreezeTokenUris(&_ERC1155NFTCustom.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "getRoleAdmin", role)

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) GetRoleAdmin(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, role [32]byte) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "getRoleAdmin", role)

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "getRoleAdmin")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC1155NFTCustom.Contract.GetRoleAdmin(&_ERC1155NFTCustom.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC1155NFTCustom.Contract.GetRoleAdmin(&_ERC1155NFTCustom.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "hasRole", role, account)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) HasRole(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, role [32]byte, account common.Address) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "hasRole", role, account)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "hasRole")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC1155NFTCustom.Contract.HasRole(&_ERC1155NFTCustom.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC1155NFTCustom.Contract.HasRole(&_ERC1155NFTCustom.CallOpts, role, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0xb6db75a0.
//
// Solidity: function isAdmin() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) IsAdmin(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "isAdmin")

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// IsAdmin is a free data retrieval call binding the contract method 0xb6db75a0.
//
// Solidity: function isAdmin() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) IsAdmin(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "isAdmin")

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "isAdmin")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// IsAdmin is a free data retrieval call binding the contract method 0xb6db75a0.
//
// Solidity: function isAdmin() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) IsAdmin() (bool, error) {
	return _ERC1155NFTCustom.Contract.IsAdmin(&_ERC1155NFTCustom.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0xb6db75a0.
//
// Solidity: function isAdmin() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) IsAdmin() (bool, error) {
	return _ERC1155NFTCustom.Contract.IsAdmin(&_ERC1155NFTCustom.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) IsApprovedForAll(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, account common.Address, operator common.Address) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "isApprovedForAll", account, operator)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "isApprovedForAll")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ERC1155NFTCustom.Contract.IsApprovedForAll(&_ERC1155NFTCustom.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ERC1155NFTCustom.Contract.IsApprovedForAll(&_ERC1155NFTCustom.CallOpts, account, operator)
}

// ListSponsorPrivilege is a free data retrieval call binding the contract method 0x28630e1d.
//
// Solidity: function listSponsorPrivilege() view returns(address[])
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) ListSponsorPrivilege(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "listSponsorPrivilege")

	if __err != nil {
		return *new([]common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, __err

}

// ListSponsorPrivilege is a free data retrieval call binding the contract method 0x28630e1d.
//
// Solidity: function listSponsorPrivilege() view returns(address[])
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) ListSponsorPrivilege(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*[]common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "listSponsorPrivilege")

	out0 := new([]common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "listSponsorPrivilege")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// ListSponsorPrivilege is a free data retrieval call binding the contract method 0x28630e1d.
//
// Solidity: function listSponsorPrivilege() view returns(address[])
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) ListSponsorPrivilege() ([]common.Address, error) {
	return _ERC1155NFTCustom.Contract.ListSponsorPrivilege(&_ERC1155NFTCustom.CallOpts)
}

// ListSponsorPrivilege is a free data retrieval call binding the contract method 0x28630e1d.
//
// Solidity: function listSponsorPrivilege() view returns(address[])
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) ListSponsorPrivilege() ([]common.Address, error) {
	return _ERC1155NFTCustom.Contract.ListSponsorPrivilege(&_ERC1155NFTCustom.CallOpts)
}

// MetadataUpdatable is a free data retrieval call binding the contract method 0x4e6f9dd6.
//
// Solidity: function metadataUpdatable() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) MetadataUpdatable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "metadataUpdatable")

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// MetadataUpdatable is a free data retrieval call binding the contract method 0x4e6f9dd6.
//
// Solidity: function metadataUpdatable() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) MetadataUpdatable(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "metadataUpdatable")

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "metadataUpdatable")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// MetadataUpdatable is a free data retrieval call binding the contract method 0x4e6f9dd6.
//
// Solidity: function metadataUpdatable() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) MetadataUpdatable() (bool, error) {
	return _ERC1155NFTCustom.Contract.MetadataUpdatable(&_ERC1155NFTCustom.CallOpts)
}

// MetadataUpdatable is a free data retrieval call binding the contract method 0x4e6f9dd6.
//
// Solidity: function metadataUpdatable() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) MetadataUpdatable() (bool, error) {
	return _ERC1155NFTCustom.Contract.MetadataUpdatable(&_ERC1155NFTCustom.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "name")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) Name(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "name")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "name")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) Name() (string, error) {
	return _ERC1155NFTCustom.Contract.Name(&_ERC1155NFTCustom.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) Name() (string, error) {
	return _ERC1155NFTCustom.Contract.Name(&_ERC1155NFTCustom.CallOpts)
}

// RoyaltiesAddress is a free data retrieval call binding the contract method 0x32882535.
//
// Solidity: function royaltiesAddress() view returns(address)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) RoyaltiesAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "royaltiesAddress")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// RoyaltiesAddress is a free data retrieval call binding the contract method 0x32882535.
//
// Solidity: function royaltiesAddress() view returns(address)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) RoyaltiesAddress(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "royaltiesAddress")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "royaltiesAddress")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// RoyaltiesAddress is a free data retrieval call binding the contract method 0x32882535.
//
// Solidity: function royaltiesAddress() view returns(address)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) RoyaltiesAddress() (common.Address, error) {
	return _ERC1155NFTCustom.Contract.RoyaltiesAddress(&_ERC1155NFTCustom.CallOpts)
}

// RoyaltiesAddress is a free data retrieval call binding the contract method 0x32882535.
//
// Solidity: function royaltiesAddress() view returns(address)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) RoyaltiesAddress() (common.Address, error) {
	return _ERC1155NFTCustom.Contract.RoyaltiesAddress(&_ERC1155NFTCustom.CallOpts)
}

// RoyaltiesBps is a free data retrieval call binding the contract method 0x99d89f9d.
//
// Solidity: function royaltiesBps() view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) RoyaltiesBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "royaltiesBps")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// RoyaltiesBps is a free data retrieval call binding the contract method 0x99d89f9d.
//
// Solidity: function royaltiesBps() view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) RoyaltiesBps(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "royaltiesBps")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "royaltiesBps")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// RoyaltiesBps is a free data retrieval call binding the contract method 0x99d89f9d.
//
// Solidity: function royaltiesBps() view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) RoyaltiesBps() (*big.Int, error) {
	return _ERC1155NFTCustom.Contract.RoyaltiesBps(&_ERC1155NFTCustom.CallOpts)
}

// RoyaltiesBps is a free data retrieval call binding the contract method 0x99d89f9d.
//
// Solidity: function royaltiesBps() view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) RoyaltiesBps() (*big.Int, error) {
	return _ERC1155NFTCustom.Contract.RoyaltiesBps(&_ERC1155NFTCustom.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	if __err != nil {
		return *new(common.Address), *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, __err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) RoyaltyInfo(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (*common.Address, **big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "royaltyInfo", tokenId, salePrice)

	out0 := new(common.Address)
	out1 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "royaltyInfo")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
		*out1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, out1, __err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _ERC1155NFTCustom.Contract.RoyaltyInfo(&_ERC1155NFTCustom.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _ERC1155NFTCustom.Contract.RoyaltyInfo(&_ERC1155NFTCustom.CallOpts, tokenId, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) SupportsInterface(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, interfaceId [4]byte) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "supportsInterface", interfaceId)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "supportsInterface")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155NFTCustom.Contract.SupportsInterface(&_ERC1155NFTCustom.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155NFTCustom.Contract.SupportsInterface(&_ERC1155NFTCustom.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "symbol")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) Symbol(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "symbol")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "symbol")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) Symbol() (string, error) {
	return _ERC1155NFTCustom.Contract.Symbol(&_ERC1155NFTCustom.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) Symbol() (string, error) {
	return _ERC1155NFTCustom.Contract.Symbol(&_ERC1155NFTCustom.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "tokenByIndex", index)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) TokenByIndex(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, index *big.Int) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "tokenByIndex", index)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "tokenByIndex")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC1155NFTCustom.Contract.TokenByIndex(&_ERC1155NFTCustom.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC1155NFTCustom.Contract.TokenByIndex(&_ERC1155NFTCustom.CallOpts, index)
}

// TokenCountOf is a free data retrieval call binding the contract method 0xb722938a.
//
// Solidity: function tokenCountOf(address owner) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) TokenCountOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "tokenCountOf", owner)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// TokenCountOf is a free data retrieval call binding the contract method 0xb722938a.
//
// Solidity: function tokenCountOf(address owner) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) TokenCountOf(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, owner common.Address) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "tokenCountOf", owner)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "tokenCountOf")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TokenCountOf is a free data retrieval call binding the contract method 0xb722938a.
//
// Solidity: function tokenCountOf(address owner) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) TokenCountOf(owner common.Address) (*big.Int, error) {
	return _ERC1155NFTCustom.Contract.TokenCountOf(&_ERC1155NFTCustom.CallOpts, owner)
}

// TokenCountOf is a free data retrieval call binding the contract method 0xb722938a.
//
// Solidity: function tokenCountOf(address owner) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) TokenCountOf(owner common.Address) (*big.Int, error) {
	return _ERC1155NFTCustom.Contract.TokenCountOf(&_ERC1155NFTCustom.CallOpts, owner)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) TokenOfOwnerByIndex(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, owner common.Address, index *big.Int) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "tokenOfOwnerByIndex", owner, index)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "tokenOfOwnerByIndex")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC1155NFTCustom.Contract.TokenOfOwnerByIndex(&_ERC1155NFTCustom.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC1155NFTCustom.Contract.TokenOfOwnerByIndex(&_ERC1155NFTCustom.CallOpts, owner, index)
}

// Tokens is a free data retrieval call binding the contract method 0x8b4864d6.
//
// Solidity: function tokens(uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) Tokens(opts *bind.CallOpts, offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "tokens", offset, limit)

	outstruct := new(struct {
		Total    *big.Int
		TokenIds []*big.Int
	})
	if __err != nil {
		return *outstruct, __err
	}

	outstruct.Total = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenIds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, __err

}

// Tokens is a free data retrieval call binding the contract method 0x8b4864d6.
//
// Solidity: function tokens(uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) Tokens(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, offset *big.Int, limit *big.Int) (*struct {
	Total    *big.Int
	TokenIds []*big.Int
}, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "tokens", offset, limit)

	outstruct := new(struct {
		Total    *big.Int
		TokenIds []*big.Int
	})

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "tokens")
		if err != nil {
			return err
		}

		outstruct.Total = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
		outstruct.TokenIds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return outstruct, __err

}

// Tokens is a free data retrieval call binding the contract method 0x8b4864d6.
//
// Solidity: function tokens(uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) Tokens(offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	return _ERC1155NFTCustom.Contract.Tokens(&_ERC1155NFTCustom.CallOpts, offset, limit)
}

// Tokens is a free data retrieval call binding the contract method 0x8b4864d6.
//
// Solidity: function tokens(uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) Tokens(offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	return _ERC1155NFTCustom.Contract.Tokens(&_ERC1155NFTCustom.CallOpts, offset, limit)
}

// TokensBurnable is a free data retrieval call binding the contract method 0xe3d52072.
//
// Solidity: function tokensBurnable() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) TokensBurnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "tokensBurnable")

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// TokensBurnable is a free data retrieval call binding the contract method 0xe3d52072.
//
// Solidity: function tokensBurnable() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) TokensBurnable(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "tokensBurnable")

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "tokensBurnable")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TokensBurnable is a free data retrieval call binding the contract method 0xe3d52072.
//
// Solidity: function tokensBurnable() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) TokensBurnable() (bool, error) {
	return _ERC1155NFTCustom.Contract.TokensBurnable(&_ERC1155NFTCustom.CallOpts)
}

// TokensBurnable is a free data retrieval call binding the contract method 0xe3d52072.
//
// Solidity: function tokensBurnable() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) TokensBurnable() (bool, error) {
	return _ERC1155NFTCustom.Contract.TokensBurnable(&_ERC1155NFTCustom.CallOpts)
}

// TokensOf is a free data retrieval call binding the contract method 0x23185dc9.
//
// Solidity: function tokensOf(address owner, uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) TokensOf(opts *bind.CallOpts, owner common.Address, offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "tokensOf", owner, offset, limit)

	outstruct := new(struct {
		Total    *big.Int
		TokenIds []*big.Int
	})
	if __err != nil {
		return *outstruct, __err
	}

	outstruct.Total = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenIds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, __err

}

// TokensOf is a free data retrieval call binding the contract method 0x23185dc9.
//
// Solidity: function tokensOf(address owner, uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) TokensOf(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, owner common.Address, offset *big.Int, limit *big.Int) (*struct {
	Total    *big.Int
	TokenIds []*big.Int
}, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "tokensOf", owner, offset, limit)

	outstruct := new(struct {
		Total    *big.Int
		TokenIds []*big.Int
	})

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "tokensOf")
		if err != nil {
			return err
		}

		outstruct.Total = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
		outstruct.TokenIds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return outstruct, __err

}

// TokensOf is a free data retrieval call binding the contract method 0x23185dc9.
//
// Solidity: function tokensOf(address owner, uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) TokensOf(owner common.Address, offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	return _ERC1155NFTCustom.Contract.TokensOf(&_ERC1155NFTCustom.CallOpts, owner, offset, limit)
}

// TokensOf is a free data retrieval call binding the contract method 0x23185dc9.
//
// Solidity: function tokensOf(address owner, uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) TokensOf(owner common.Address, offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	return _ERC1155NFTCustom.Contract.TokensOf(&_ERC1155NFTCustom.CallOpts, owner, offset, limit)
}

// TokensTransferableByAdmin is a free data retrieval call binding the contract method 0xd9fab275.
//
// Solidity: function tokensTransferableByAdmin() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) TokensTransferableByAdmin(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "tokensTransferableByAdmin")

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// TokensTransferableByAdmin is a free data retrieval call binding the contract method 0xd9fab275.
//
// Solidity: function tokensTransferableByAdmin() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) TokensTransferableByAdmin(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "tokensTransferableByAdmin")

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "tokensTransferableByAdmin")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TokensTransferableByAdmin is a free data retrieval call binding the contract method 0xd9fab275.
//
// Solidity: function tokensTransferableByAdmin() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) TokensTransferableByAdmin() (bool, error) {
	return _ERC1155NFTCustom.Contract.TokensTransferableByAdmin(&_ERC1155NFTCustom.CallOpts)
}

// TokensTransferableByAdmin is a free data retrieval call binding the contract method 0xd9fab275.
//
// Solidity: function tokensTransferableByAdmin() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) TokensTransferableByAdmin() (bool, error) {
	return _ERC1155NFTCustom.Contract.TokensTransferableByAdmin(&_ERC1155NFTCustom.CallOpts)
}

// TokensTransferableByUser is a free data retrieval call binding the contract method 0x7915c570.
//
// Solidity: function tokensTransferableByUser() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) TokensTransferableByUser(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "tokensTransferableByUser")

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// TokensTransferableByUser is a free data retrieval call binding the contract method 0x7915c570.
//
// Solidity: function tokensTransferableByUser() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) TokensTransferableByUser(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "tokensTransferableByUser")

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "tokensTransferableByUser")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TokensTransferableByUser is a free data retrieval call binding the contract method 0x7915c570.
//
// Solidity: function tokensTransferableByUser() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) TokensTransferableByUser() (bool, error) {
	return _ERC1155NFTCustom.Contract.TokensTransferableByUser(&_ERC1155NFTCustom.CallOpts)
}

// TokensTransferableByUser is a free data retrieval call binding the contract method 0x7915c570.
//
// Solidity: function tokensTransferableByUser() view returns(bool)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) TokensTransferableByUser() (bool, error) {
	return _ERC1155NFTCustom.Contract.TokensTransferableByUser(&_ERC1155NFTCustom.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "totalSupply")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) TotalSupply(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "totalSupply")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "totalSupply")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) TotalSupply() (*big.Int, error) {
	return _ERC1155NFTCustom.Contract.TotalSupply(&_ERC1155NFTCustom.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC1155NFTCustom.Contract.TotalSupply(&_ERC1155NFTCustom.CallOpts)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) TotalSupply0(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "totalSupply0", tokenId)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) TotalSupply0(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, tokenId *big.Int) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "totalSupply0", tokenId)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "totalSupply0")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) TotalSupply0(tokenId *big.Int) (*big.Int, error) {
	return _ERC1155NFTCustom.Contract.TotalSupply0(&_ERC1155NFTCustom.CallOpts, tokenId)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) TotalSupply0(tokenId *big.Int) (*big.Int, error) {
	return _ERC1155NFTCustom.Contract.TotalSupply0(&_ERC1155NFTCustom.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	__err := _ERC1155NFTCustom.contract.Call(opts, &out, "uri", tokenId)

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkCaller) Uri(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, tokenId *big.Int) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFTCustom.contract.GenRequest(opts, "uri", tokenId)

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFTCustom.contract.DecodeOutput(&out, rawOut, "uri")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) Uri(tokenId *big.Int) (string, error) {
	return _ERC1155NFTCustom.Contract.Uri(&_ERC1155NFTCustom.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_ERC1155NFTCustom *ERC1155NFTCustomCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _ERC1155NFTCustom.Contract.Uri(&_ERC1155NFTCustom.CallOpts, tokenId)
}

// AddSponsorPrivilege is a paid mutator transaction binding the contract method 0x938ead7c.
//
// Solidity: function addSponsorPrivilege(address[] whites) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) AddSponsorPrivilege(opts *bind.TransactOpts, whites []common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "addSponsorPrivilege", whites)
}

// AddSponsorPrivilege is a paid mutator transaction binding the contract method 0x938ead7c.
//
// Solidity: function addSponsorPrivilege(address[] whites) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) AddSponsorPrivilege(opts *bind.TransactOpts, whites []common.Address) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "addSponsorPrivilege", whites)
}

// AddSponsorPrivilege is a paid mutator transaction binding the contract method 0x938ead7c.
//
// Solidity: function addSponsorPrivilege(address[] whites) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) AddSponsorPrivilege(whites []common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.AddSponsorPrivilege(&_ERC1155NFTCustom.TransactOpts, whites)
}

// AddSponsorPrivilege is a paid mutator transaction binding the contract method 0x938ead7c.
//
// Solidity: function addSponsorPrivilege(address[] whites) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) AddSponsorPrivilege(whites []common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.AddSponsorPrivilege(&_ERC1155NFTCustom.TransactOpts, whites)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 id, uint256 value) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) Burn(opts *bind.TransactOpts, user common.Address, id *big.Int, value *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "burn", user, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 id, uint256 value) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) Burn(opts *bind.TransactOpts, user common.Address, id *big.Int, value *big.Int) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "burn", user, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 id, uint256 value) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) Burn(user common.Address, id *big.Int, value *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.Burn(&_ERC1155NFTCustom.TransactOpts, user, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 id, uint256 value) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) Burn(user common.Address, id *big.Int, value *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.Burn(&_ERC1155NFTCustom.TransactOpts, user, id, value)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address user, uint256[] ids, uint256[] values) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) BurnBatch(opts *bind.TransactOpts, user common.Address, ids []*big.Int, values []*big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "burnBatch", user, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address user, uint256[] ids, uint256[] values) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) BurnBatch(opts *bind.TransactOpts, user common.Address, ids []*big.Int, values []*big.Int) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "burnBatch", user, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address user, uint256[] ids, uint256[] values) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) BurnBatch(user common.Address, ids []*big.Int, values []*big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.BurnBatch(&_ERC1155NFTCustom.TransactOpts, user, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address user, uint256[] ids, uint256[] values) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) BurnBatch(user common.Address, ids []*big.Int, values []*big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.BurnBatch(&_ERC1155NFTCustom.TransactOpts, user, ids, values)
}

// FreezeGlobalMetadata is a paid mutator transaction binding the contract method 0x092e7106.
//
// Solidity: function freezeGlobalMetadata() returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) FreezeGlobalMetadata(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "freezeGlobalMetadata")
}

// FreezeGlobalMetadata is a paid mutator transaction binding the contract method 0x092e7106.
//
// Solidity: function freezeGlobalMetadata() returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) FreezeGlobalMetadata(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "freezeGlobalMetadata")
}

// FreezeGlobalMetadata is a paid mutator transaction binding the contract method 0x092e7106.
//
// Solidity: function freezeGlobalMetadata() returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) FreezeGlobalMetadata() (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.FreezeGlobalMetadata(&_ERC1155NFTCustom.TransactOpts)
}

// FreezeGlobalMetadata is a paid mutator transaction binding the contract method 0x092e7106.
//
// Solidity: function freezeGlobalMetadata() returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) FreezeGlobalMetadata() (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.FreezeGlobalMetadata(&_ERC1155NFTCustom.TransactOpts)
}

// FreezeTokenURI is a paid mutator transaction binding the contract method 0x385c0eb0.
//
// Solidity: function freezeTokenURI(uint256 tokenId) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) FreezeTokenURI(opts *bind.TransactOpts, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "freezeTokenURI", tokenId)
}

// FreezeTokenURI is a paid mutator transaction binding the contract method 0x385c0eb0.
//
// Solidity: function freezeTokenURI(uint256 tokenId) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) FreezeTokenURI(opts *bind.TransactOpts, tokenId *big.Int) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "freezeTokenURI", tokenId)
}

// FreezeTokenURI is a paid mutator transaction binding the contract method 0x385c0eb0.
//
// Solidity: function freezeTokenURI(uint256 tokenId) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) FreezeTokenURI(tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.FreezeTokenURI(&_ERC1155NFTCustom.TransactOpts, tokenId)
}

// FreezeTokenURI is a paid mutator transaction binding the contract method 0x385c0eb0.
//
// Solidity: function freezeTokenURI(uint256 tokenId) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) FreezeTokenURI(tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.FreezeTokenURI(&_ERC1155NFTCustom.TransactOpts, tokenId)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address user) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) GrantAdminRole(opts *bind.TransactOpts, user common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "grantAdminRole", user)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address user) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) GrantAdminRole(opts *bind.TransactOpts, user common.Address) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "grantAdminRole", user)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address user) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) GrantAdminRole(user common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.GrantAdminRole(&_ERC1155NFTCustom.TransactOpts, user)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address user) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) GrantAdminRole(user common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.GrantAdminRole(&_ERC1155NFTCustom.TransactOpts, user)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address user) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) GrantMintRole(opts *bind.TransactOpts, user common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "grantMintRole", user)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address user) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) GrantMintRole(opts *bind.TransactOpts, user common.Address) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "grantMintRole", user)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address user) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) GrantMintRole(user common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.GrantMintRole(&_ERC1155NFTCustom.TransactOpts, user)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address user) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) GrantMintRole(user common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.GrantMintRole(&_ERC1155NFTCustom.TransactOpts, user)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) GrantRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.GrantRole(&_ERC1155NFTCustom.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.GrantRole(&_ERC1155NFTCustom.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x7d251387.
//
// Solidity: function initialize(string _name, string _symbol, string baseURI, uint256 royaltiesBps, address royaltiesAddress, address[] owners, bool tokensTransferableByAdmin, bool tokensTransferableByUser, bool isSetSponsorWhitelistForAllUser) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, baseURI string, royaltiesBps *big.Int, royaltiesAddress common.Address, owners []common.Address, tokensTransferableByAdmin bool, tokensTransferableByUser bool, isSetSponsorWhitelistForAllUser bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "initialize", _name, _symbol, baseURI, royaltiesBps, royaltiesAddress, owners, tokensTransferableByAdmin, tokensTransferableByUser, isSetSponsorWhitelistForAllUser)
}

// Initialize is a paid mutator transaction binding the contract method 0x7d251387.
//
// Solidity: function initialize(string _name, string _symbol, string baseURI, uint256 royaltiesBps, address royaltiesAddress, address[] owners, bool tokensTransferableByAdmin, bool tokensTransferableByUser, bool isSetSponsorWhitelistForAllUser) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, baseURI string, royaltiesBps *big.Int, royaltiesAddress common.Address, owners []common.Address, tokensTransferableByAdmin bool, tokensTransferableByUser bool, isSetSponsorWhitelistForAllUser bool) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "initialize", _name, _symbol, baseURI, royaltiesBps, royaltiesAddress, owners, tokensTransferableByAdmin, tokensTransferableByUser, isSetSponsorWhitelistForAllUser)
}

// Initialize is a paid mutator transaction binding the contract method 0x7d251387.
//
// Solidity: function initialize(string _name, string _symbol, string baseURI, uint256 royaltiesBps, address royaltiesAddress, address[] owners, bool tokensTransferableByAdmin, bool tokensTransferableByUser, bool isSetSponsorWhitelistForAllUser) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) Initialize(_name string, _symbol string, baseURI string, royaltiesBps *big.Int, royaltiesAddress common.Address, owners []common.Address, tokensTransferableByAdmin bool, tokensTransferableByUser bool, isSetSponsorWhitelistForAllUser bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.Initialize(&_ERC1155NFTCustom.TransactOpts, _name, _symbol, baseURI, royaltiesBps, royaltiesAddress, owners, tokensTransferableByAdmin, tokensTransferableByUser, isSetSponsorWhitelistForAllUser)
}

// Initialize is a paid mutator transaction binding the contract method 0x7d251387.
//
// Solidity: function initialize(string _name, string _symbol, string baseURI, uint256 royaltiesBps, address royaltiesAddress, address[] owners, bool tokensTransferableByAdmin, bool tokensTransferableByUser, bool isSetSponsorWhitelistForAllUser) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) Initialize(_name string, _symbol string, baseURI string, royaltiesBps *big.Int, royaltiesAddress common.Address, owners []common.Address, tokensTransferableByAdmin bool, tokensTransferableByUser bool, isSetSponsorWhitelistForAllUser bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.Initialize(&_ERC1155NFTCustom.TransactOpts, _name, _symbol, baseURI, royaltiesBps, royaltiesAddress, owners, tokensTransferableByAdmin, tokensTransferableByUser, isSetSponsorWhitelistForAllUser)
}

// MintTo is a paid mutator transaction binding the contract method 0x3dbd5b25.
//
// Solidity: function mintTo(address to, uint256 id, uint256 amount, string tokenUri) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) MintTo(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int, tokenUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "mintTo", to, id, amount, tokenUri)
}

// MintTo is a paid mutator transaction binding the contract method 0x3dbd5b25.
//
// Solidity: function mintTo(address to, uint256 id, uint256 amount, string tokenUri) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) MintTo(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int, tokenUri string) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "mintTo", to, id, amount, tokenUri)
}

// MintTo is a paid mutator transaction binding the contract method 0x3dbd5b25.
//
// Solidity: function mintTo(address to, uint256 id, uint256 amount, string tokenUri) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) MintTo(to common.Address, id *big.Int, amount *big.Int, tokenUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.MintTo(&_ERC1155NFTCustom.TransactOpts, to, id, amount, tokenUri)
}

// MintTo is a paid mutator transaction binding the contract method 0x3dbd5b25.
//
// Solidity: function mintTo(address to, uint256 id, uint256 amount, string tokenUri) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) MintTo(to common.Address, id *big.Int, amount *big.Int, tokenUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.MintTo(&_ERC1155NFTCustom.TransactOpts, to, id, amount, tokenUri)
}

// MintTo0 is a paid mutator transaction binding the contract method 0x9f6ed25f.
//
// Solidity: function mintTo(address to, uint256 id, string tokenUri) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) MintTo0(opts *bind.TransactOpts, to common.Address, id *big.Int, tokenUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "mintTo0", to, id, tokenUri)
}

// MintTo0 is a paid mutator transaction binding the contract method 0x9f6ed25f.
//
// Solidity: function mintTo(address to, uint256 id, string tokenUri) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) MintTo0(opts *bind.TransactOpts, to common.Address, id *big.Int, tokenUri string) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "mintTo0", to, id, tokenUri)
}

// MintTo0 is a paid mutator transaction binding the contract method 0x9f6ed25f.
//
// Solidity: function mintTo(address to, uint256 id, string tokenUri) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) MintTo0(to common.Address, id *big.Int, tokenUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.MintTo0(&_ERC1155NFTCustom.TransactOpts, to, id, tokenUri)
}

// MintTo0 is a paid mutator transaction binding the contract method 0x9f6ed25f.
//
// Solidity: function mintTo(address to, uint256 id, string tokenUri) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) MintTo0(to common.Address, id *big.Int, tokenUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.MintTo0(&_ERC1155NFTCustom.TransactOpts, to, id, tokenUri)
}

// MintToBatch is a paid mutator transaction binding the contract method 0xe21e29c6.
//
// Solidity: function mintToBatch(address[] tos, uint256[] ids, uint256[] amounts, string[] uris) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) MintToBatch(opts *bind.TransactOpts, tos []common.Address, ids []*big.Int, amounts []*big.Int, uris []string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "mintToBatch", tos, ids, amounts, uris)
}

// MintToBatch is a paid mutator transaction binding the contract method 0xe21e29c6.
//
// Solidity: function mintToBatch(address[] tos, uint256[] ids, uint256[] amounts, string[] uris) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) MintToBatch(opts *bind.TransactOpts, tos []common.Address, ids []*big.Int, amounts []*big.Int, uris []string) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "mintToBatch", tos, ids, amounts, uris)
}

// MintToBatch is a paid mutator transaction binding the contract method 0xe21e29c6.
//
// Solidity: function mintToBatch(address[] tos, uint256[] ids, uint256[] amounts, string[] uris) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) MintToBatch(tos []common.Address, ids []*big.Int, amounts []*big.Int, uris []string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.MintToBatch(&_ERC1155NFTCustom.TransactOpts, tos, ids, amounts, uris)
}

// MintToBatch is a paid mutator transaction binding the contract method 0xe21e29c6.
//
// Solidity: function mintToBatch(address[] tos, uint256[] ids, uint256[] amounts, string[] uris) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) MintToBatch(tos []common.Address, ids []*big.Int, amounts []*big.Int, uris []string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.MintToBatch(&_ERC1155NFTCustom.TransactOpts, tos, ids, amounts, uris)
}

// RemoveSponsorPrivilege is a paid mutator transaction binding the contract method 0x76104e53.
//
// Solidity: function removeSponsorPrivilege(address[] whites) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) RemoveSponsorPrivilege(opts *bind.TransactOpts, whites []common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "removeSponsorPrivilege", whites)
}

// RemoveSponsorPrivilege is a paid mutator transaction binding the contract method 0x76104e53.
//
// Solidity: function removeSponsorPrivilege(address[] whites) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) RemoveSponsorPrivilege(opts *bind.TransactOpts, whites []common.Address) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "removeSponsorPrivilege", whites)
}

// RemoveSponsorPrivilege is a paid mutator transaction binding the contract method 0x76104e53.
//
// Solidity: function removeSponsorPrivilege(address[] whites) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) RemoveSponsorPrivilege(whites []common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.RemoveSponsorPrivilege(&_ERC1155NFTCustom.TransactOpts, whites)
}

// RemoveSponsorPrivilege is a paid mutator transaction binding the contract method 0x76104e53.
//
// Solidity: function removeSponsorPrivilege(address[] whites) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) RemoveSponsorPrivilege(whites []common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.RemoveSponsorPrivilege(&_ERC1155NFTCustom.TransactOpts, whites)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) RenounceRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.RenounceRole(&_ERC1155NFTCustom.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.RenounceRole(&_ERC1155NFTCustom.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) RevokeRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.RevokeRole(&_ERC1155NFTCustom.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.RevokeRole(&_ERC1155NFTCustom.TransactOpts, role, account)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.SafeBatchTransferFrom(&_ERC1155NFTCustom.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.SafeBatchTransferFrom(&_ERC1155NFTCustom.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.SafeTransferFrom(&_ERC1155NFTCustom.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.SafeTransferFrom(&_ERC1155NFTCustom.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) SetApprovalForAll(operator common.Address, approved bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.SetApprovalForAll(&_ERC1155NFTCustom.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.SetApprovalForAll(&_ERC1155NFTCustom.TransactOpts, operator, approved)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x5de42985.
//
// Solidity: function setRoyalties(uint256 _royaltiesBps, address _royaltiesAddress) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) SetRoyalties(opts *bind.TransactOpts, _royaltiesBps *big.Int, _royaltiesAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "setRoyalties", _royaltiesBps, _royaltiesAddress)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x5de42985.
//
// Solidity: function setRoyalties(uint256 _royaltiesBps, address _royaltiesAddress) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) SetRoyalties(opts *bind.TransactOpts, _royaltiesBps *big.Int, _royaltiesAddress common.Address) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "setRoyalties", _royaltiesBps, _royaltiesAddress)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x5de42985.
//
// Solidity: function setRoyalties(uint256 _royaltiesBps, address _royaltiesAddress) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) SetRoyalties(_royaltiesBps *big.Int, _royaltiesAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.SetRoyalties(&_ERC1155NFTCustom.TransactOpts, _royaltiesBps, _royaltiesAddress)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x5de42985.
//
// Solidity: function setRoyalties(uint256 _royaltiesBps, address _royaltiesAddress) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) SetRoyalties(_royaltiesBps *big.Int, _royaltiesAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.SetRoyalties(&_ERC1155NFTCustom.TransactOpts, _royaltiesBps, _royaltiesAddress)
}

// SetTokensTransferable is a paid mutator transaction binding the contract method 0xd4e0456b.
//
// Solidity: function setTokensTransferable(bool transferableByAdmin, bool transferableByUser) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) SetTokensTransferable(opts *bind.TransactOpts, transferableByAdmin bool, transferableByUser bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "setTokensTransferable", transferableByAdmin, transferableByUser)
}

// SetTokensTransferable is a paid mutator transaction binding the contract method 0xd4e0456b.
//
// Solidity: function setTokensTransferable(bool transferableByAdmin, bool transferableByUser) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) SetTokensTransferable(opts *bind.TransactOpts, transferableByAdmin bool, transferableByUser bool) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "setTokensTransferable", transferableByAdmin, transferableByUser)
}

// SetTokensTransferable is a paid mutator transaction binding the contract method 0xd4e0456b.
//
// Solidity: function setTokensTransferable(bool transferableByAdmin, bool transferableByUser) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) SetTokensTransferable(transferableByAdmin bool, transferableByUser bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.SetTokensTransferable(&_ERC1155NFTCustom.TransactOpts, transferableByAdmin, transferableByUser)
}

// SetTokensTransferable is a paid mutator transaction binding the contract method 0xd4e0456b.
//
// Solidity: function setTokensTransferable(bool transferableByAdmin, bool transferableByUser) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) SetTokensTransferable(transferableByAdmin bool, transferableByUser bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.SetTokensTransferable(&_ERC1155NFTCustom.TransactOpts, transferableByAdmin, transferableByUser)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newURI) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) SetURI(opts *bind.TransactOpts, newURI string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "setURI", newURI)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newURI) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) SetURI(opts *bind.TransactOpts, newURI string) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "setURI", newURI)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newURI) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) SetURI(newURI string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.SetURI(&_ERC1155NFTCustom.TransactOpts, newURI)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newURI) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) SetURI(newURI string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.SetURI(&_ERC1155NFTCustom.TransactOpts, newURI)
}

// TransferBatchByAdmin is a paid mutator transaction binding the contract method 0x28b76be3.
//
// Solidity: function transferBatchByAdmin(address[] users, address[] to, uint256[] ids, uint256[] amounts) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) TransferBatchByAdmin(opts *bind.TransactOpts, users []common.Address, to []common.Address, ids []*big.Int, amounts []*big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "transferBatchByAdmin", users, to, ids, amounts)
}

// TransferBatchByAdmin is a paid mutator transaction binding the contract method 0x28b76be3.
//
// Solidity: function transferBatchByAdmin(address[] users, address[] to, uint256[] ids, uint256[] amounts) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) TransferBatchByAdmin(opts *bind.TransactOpts, users []common.Address, to []common.Address, ids []*big.Int, amounts []*big.Int) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "transferBatchByAdmin", users, to, ids, amounts)
}

// TransferBatchByAdmin is a paid mutator transaction binding the contract method 0x28b76be3.
//
// Solidity: function transferBatchByAdmin(address[] users, address[] to, uint256[] ids, uint256[] amounts) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) TransferBatchByAdmin(users []common.Address, to []common.Address, ids []*big.Int, amounts []*big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.TransferBatchByAdmin(&_ERC1155NFTCustom.TransactOpts, users, to, ids, amounts)
}

// TransferBatchByAdmin is a paid mutator transaction binding the contract method 0x28b76be3.
//
// Solidity: function transferBatchByAdmin(address[] users, address[] to, uint256[] ids, uint256[] amounts) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) TransferBatchByAdmin(users []common.Address, to []common.Address, ids []*big.Int, amounts []*big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.TransferBatchByAdmin(&_ERC1155NFTCustom.TransactOpts, users, to, ids, amounts)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0x41daa6bc.
//
// Solidity: function transferByAdmin(address user, address to, uint256 id, uint256 amount) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) TransferByAdmin(opts *bind.TransactOpts, user common.Address, to common.Address, id *big.Int, amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "transferByAdmin", user, to, id, amount)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0x41daa6bc.
//
// Solidity: function transferByAdmin(address user, address to, uint256 id, uint256 amount) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) TransferByAdmin(opts *bind.TransactOpts, user common.Address, to common.Address, id *big.Int, amount *big.Int) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "transferByAdmin", user, to, id, amount)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0x41daa6bc.
//
// Solidity: function transferByAdmin(address user, address to, uint256 id, uint256 amount) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) TransferByAdmin(user common.Address, to common.Address, id *big.Int, amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.TransferByAdmin(&_ERC1155NFTCustom.TransactOpts, user, to, id, amount)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0x41daa6bc.
//
// Solidity: function transferByAdmin(address user, address to, uint256 id, uint256 amount) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) TransferByAdmin(user common.Address, to common.Address, id *big.Int, amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.TransferByAdmin(&_ERC1155NFTCustom.TransactOpts, user, to, id, amount)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string newUri) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactor) UpdateTokenURI(opts *bind.TransactOpts, tokenId *big.Int, newUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.contract.Transact(opts, "updateTokenURI", tokenId, newUri)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string newUri) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomBulkTransactor) UpdateTokenURI(opts *bind.TransactOpts, tokenId *big.Int, newUri string) types.UnsignedTransaction {
	return _ERC1155NFTCustom.contract.GenUnsignedTransaction(opts, "updateTokenURI", tokenId, newUri)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string newUri) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomSession) UpdateTokenURI(tokenId *big.Int, newUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.UpdateTokenURI(&_ERC1155NFTCustom.TransactOpts, tokenId, newUri)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string newUri) returns()
func (_ERC1155NFTCustom *ERC1155NFTCustomTransactorSession) UpdateTokenURI(tokenId *big.Int, newUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFTCustom.Contract.UpdateTokenURI(&_ERC1155NFTCustom.TransactOpts, tokenId, newUri)
}

// ERC1155NFTCustomApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomApprovalForAllIterator struct {
	Event *ERC1155NFTCustomApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155NFTCustomApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTCustomApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155NFTCustomApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155NFTCustomApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTCustomApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTCustomApprovalForAll represents a ApprovalForAll event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// ERC1155NFTCustomApprovalForAllOrChainReorg represents a ApprovalForAll subscription event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomApprovalForAllOrChainReorg struct {
	Event      *ERC1155NFTCustomApprovalForAll
	ChainReorg *types.ChainReorg
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ERC1155NFTCustomApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, err := _ERC1155NFTCustom.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomApprovalForAllIterator{contract: _ERC1155NFTCustom.contract, event: "ApprovalForAll", logs: logs}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC1155NFTCustomApprovalForAllOrChainReorg, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC1155NFTCustom.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTCustomApprovalForAllOrChainReorg)
				event.Event = new(ERC1155NFTCustomApprovalForAll)

				if log.ChainReorg == nil {
					if err := _ERC1155NFTCustom.contract.UnpackLog(event.Event, "ApprovalForAll", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) ParseApprovalForAll(log types.Log) (*ERC1155NFTCustomApprovalForAll, error) {
	event := new(ERC1155NFTCustomApprovalForAll)
	if err := _ERC1155NFTCustom.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTCustomInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomInitializedIterator struct {
	Event *ERC1155NFTCustomInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155NFTCustomInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTCustomInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155NFTCustomInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155NFTCustomInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTCustomInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTCustomInitialized represents a Initialized event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// ERC1155NFTCustomInitializedOrChainReorg represents a Initialized subscription event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomInitializedOrChainReorg struct {
	Event      *ERC1155NFTCustomInitialized
	ChainReorg *types.ChainReorg
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC1155NFTCustomInitializedIterator, error) {

	logs, err := _ERC1155NFTCustom.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomInitializedIterator{contract: _ERC1155NFTCustom.contract, event: "Initialized", logs: logs}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC1155NFTCustomInitializedOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _ERC1155NFTCustom.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTCustomInitializedOrChainReorg)
				event.Event = new(ERC1155NFTCustomInitialized)

				if log.ChainReorg == nil {
					if err := _ERC1155NFTCustom.contract.UnpackLog(event.Event, "Initialized", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) ParseInitialized(log types.Log) (*ERC1155NFTCustomInitialized, error) {
	event := new(ERC1155NFTCustomInitialized)
	if err := _ERC1155NFTCustom.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTCustomPermanentURIIterator is returned from FilterPermanentURI and is used to iterate over the raw logs and unpacked data for PermanentURI events raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomPermanentURIIterator struct {
	Event *ERC1155NFTCustomPermanentURI // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155NFTCustomPermanentURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTCustomPermanentURI)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155NFTCustomPermanentURI)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155NFTCustomPermanentURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTCustomPermanentURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTCustomPermanentURI represents a PermanentURI event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomPermanentURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// ERC1155NFTCustomPermanentURIOrChainReorg represents a PermanentURI subscription event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomPermanentURIOrChainReorg struct {
	Event      *ERC1155NFTCustomPermanentURI
	ChainReorg *types.ChainReorg
}

// FilterPermanentURI is a free log retrieval operation binding the contract event 0xa109ba539900bf1b633f956d63c96fc89b814c7287f7aa50a9216d0b55657207.
//
// Solidity: event PermanentURI(string _value, uint256 indexed _id)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) FilterPermanentURI(opts *bind.FilterOpts, _id []*big.Int) (*ERC1155NFTCustomPermanentURIIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, err := _ERC1155NFTCustom.contract.FilterLogs(opts, "PermanentURI", _idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomPermanentURIIterator{contract: _ERC1155NFTCustom.contract, event: "PermanentURI", logs: logs}, nil
}

// WatchPermanentURI is a free log subscription operation binding the contract event 0xa109ba539900bf1b633f956d63c96fc89b814c7287f7aa50a9216d0b55657207.
//
// Solidity: event PermanentURI(string _value, uint256 indexed _id)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) WatchPermanentURI(opts *bind.WatchOpts, sink chan<- *ERC1155NFTCustomPermanentURIOrChainReorg, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _ERC1155NFTCustom.contract.WatchLogs(opts, "PermanentURI", _idRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTCustomPermanentURIOrChainReorg)
				event.Event = new(ERC1155NFTCustomPermanentURI)

				if log.ChainReorg == nil {
					if err := _ERC1155NFTCustom.contract.UnpackLog(event.Event, "PermanentURI", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePermanentURI is a log parse operation binding the contract event 0xa109ba539900bf1b633f956d63c96fc89b814c7287f7aa50a9216d0b55657207.
//
// Solidity: event PermanentURI(string _value, uint256 indexed _id)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) ParsePermanentURI(log types.Log) (*ERC1155NFTCustomPermanentURI, error) {
	event := new(ERC1155NFTCustomPermanentURI)
	if err := _ERC1155NFTCustom.contract.UnpackLog(event, "PermanentURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTCustomPermanentURIGlobalIterator is returned from FilterPermanentURIGlobal and is used to iterate over the raw logs and unpacked data for PermanentURIGlobal events raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomPermanentURIGlobalIterator struct {
	Event *ERC1155NFTCustomPermanentURIGlobal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155NFTCustomPermanentURIGlobalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTCustomPermanentURIGlobal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155NFTCustomPermanentURIGlobal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155NFTCustomPermanentURIGlobalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTCustomPermanentURIGlobalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTCustomPermanentURIGlobal represents a PermanentURIGlobal event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomPermanentURIGlobal struct {
	Raw types.Log // Blockchain specific contextual infos
}

// ERC1155NFTCustomPermanentURIGlobalOrChainReorg represents a PermanentURIGlobal subscription event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomPermanentURIGlobalOrChainReorg struct {
	Event      *ERC1155NFTCustomPermanentURIGlobal
	ChainReorg *types.ChainReorg
}

// FilterPermanentURIGlobal is a free log retrieval operation binding the contract event 0xb59f45df38ec0d34114b1248c38a29cdbccbf3e745ae3ef310ac66199a4ceccf.
//
// Solidity: event PermanentURIGlobal()
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) FilterPermanentURIGlobal(opts *bind.FilterOpts) (*ERC1155NFTCustomPermanentURIGlobalIterator, error) {

	logs, err := _ERC1155NFTCustom.contract.FilterLogs(opts, "PermanentURIGlobal")
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomPermanentURIGlobalIterator{contract: _ERC1155NFTCustom.contract, event: "PermanentURIGlobal", logs: logs}, nil
}

// WatchPermanentURIGlobal is a free log subscription operation binding the contract event 0xb59f45df38ec0d34114b1248c38a29cdbccbf3e745ae3ef310ac66199a4ceccf.
//
// Solidity: event PermanentURIGlobal()
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) WatchPermanentURIGlobal(opts *bind.WatchOpts, sink chan<- *ERC1155NFTCustomPermanentURIGlobalOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _ERC1155NFTCustom.contract.WatchLogs(opts, "PermanentURIGlobal")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTCustomPermanentURIGlobalOrChainReorg)
				event.Event = new(ERC1155NFTCustomPermanentURIGlobal)

				if log.ChainReorg == nil {
					if err := _ERC1155NFTCustom.contract.UnpackLog(event.Event, "PermanentURIGlobal", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePermanentURIGlobal is a log parse operation binding the contract event 0xb59f45df38ec0d34114b1248c38a29cdbccbf3e745ae3ef310ac66199a4ceccf.
//
// Solidity: event PermanentURIGlobal()
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) ParsePermanentURIGlobal(log types.Log) (*ERC1155NFTCustomPermanentURIGlobal, error) {
	event := new(ERC1155NFTCustomPermanentURIGlobal)
	if err := _ERC1155NFTCustom.contract.UnpackLog(event, "PermanentURIGlobal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTCustomRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomRoleAdminChangedIterator struct {
	Event *ERC1155NFTCustomRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155NFTCustomRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTCustomRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155NFTCustomRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155NFTCustomRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTCustomRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTCustomRoleAdminChanged represents a RoleAdminChanged event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// ERC1155NFTCustomRoleAdminChangedOrChainReorg represents a RoleAdminChanged subscription event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomRoleAdminChangedOrChainReorg struct {
	Event      *ERC1155NFTCustomRoleAdminChanged
	ChainReorg *types.ChainReorg
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ERC1155NFTCustomRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, err := _ERC1155NFTCustom.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomRoleAdminChangedIterator{contract: _ERC1155NFTCustom.contract, event: "RoleAdminChanged", logs: logs}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC1155NFTCustomRoleAdminChangedOrChainReorg, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ERC1155NFTCustom.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTCustomRoleAdminChangedOrChainReorg)
				event.Event = new(ERC1155NFTCustomRoleAdminChanged)

				if log.ChainReorg == nil {
					if err := _ERC1155NFTCustom.contract.UnpackLog(event.Event, "RoleAdminChanged", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) ParseRoleAdminChanged(log types.Log) (*ERC1155NFTCustomRoleAdminChanged, error) {
	event := new(ERC1155NFTCustomRoleAdminChanged)
	if err := _ERC1155NFTCustom.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTCustomRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomRoleGrantedIterator struct {
	Event *ERC1155NFTCustomRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155NFTCustomRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTCustomRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155NFTCustomRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155NFTCustomRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTCustomRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTCustomRoleGranted represents a RoleGranted event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// ERC1155NFTCustomRoleGrantedOrChainReorg represents a RoleGranted subscription event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomRoleGrantedOrChainReorg struct {
	Event      *ERC1155NFTCustomRoleGranted
	ChainReorg *types.ChainReorg
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC1155NFTCustomRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, err := _ERC1155NFTCustom.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomRoleGrantedIterator{contract: _ERC1155NFTCustom.contract, event: "RoleGranted", logs: logs}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ERC1155NFTCustomRoleGrantedOrChainReorg, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC1155NFTCustom.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTCustomRoleGrantedOrChainReorg)
				event.Event = new(ERC1155NFTCustomRoleGranted)

				if log.ChainReorg == nil {
					if err := _ERC1155NFTCustom.contract.UnpackLog(event.Event, "RoleGranted", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) ParseRoleGranted(log types.Log) (*ERC1155NFTCustomRoleGranted, error) {
	event := new(ERC1155NFTCustomRoleGranted)
	if err := _ERC1155NFTCustom.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTCustomRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomRoleRevokedIterator struct {
	Event *ERC1155NFTCustomRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155NFTCustomRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTCustomRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155NFTCustomRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155NFTCustomRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTCustomRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTCustomRoleRevoked represents a RoleRevoked event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// ERC1155NFTCustomRoleRevokedOrChainReorg represents a RoleRevoked subscription event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomRoleRevokedOrChainReorg struct {
	Event      *ERC1155NFTCustomRoleRevoked
	ChainReorg *types.ChainReorg
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC1155NFTCustomRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, err := _ERC1155NFTCustom.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomRoleRevokedIterator{contract: _ERC1155NFTCustom.contract, event: "RoleRevoked", logs: logs}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ERC1155NFTCustomRoleRevokedOrChainReorg, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC1155NFTCustom.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTCustomRoleRevokedOrChainReorg)
				event.Event = new(ERC1155NFTCustomRoleRevoked)

				if log.ChainReorg == nil {
					if err := _ERC1155NFTCustom.contract.UnpackLog(event.Event, "RoleRevoked", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) ParseRoleRevoked(log types.Log) (*ERC1155NFTCustomRoleRevoked, error) {
	event := new(ERC1155NFTCustomRoleRevoked)
	if err := _ERC1155NFTCustom.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTCustomRoyaltyUpdatedIterator is returned from FilterRoyaltyUpdated and is used to iterate over the raw logs and unpacked data for RoyaltyUpdated events raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomRoyaltyUpdatedIterator struct {
	Event *ERC1155NFTCustomRoyaltyUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155NFTCustomRoyaltyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTCustomRoyaltyUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155NFTCustomRoyaltyUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155NFTCustomRoyaltyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTCustomRoyaltyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTCustomRoyaltyUpdated represents a RoyaltyUpdated event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomRoyaltyUpdated struct {
	RoyaltiesBps     *big.Int
	RoyaltiesAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// ERC1155NFTCustomRoyaltyUpdatedOrChainReorg represents a RoyaltyUpdated subscription event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomRoyaltyUpdatedOrChainReorg struct {
	Event      *ERC1155NFTCustomRoyaltyUpdated
	ChainReorg *types.ChainReorg
}

// FilterRoyaltyUpdated is a free log retrieval operation binding the contract event 0x13ceafb5e8ec39102bd452913a5c00a05f3b060d636d3a567c1c80e8b4321fe7.
//
// Solidity: event RoyaltyUpdated(uint256 royaltiesBps, address royaltiesAddress)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) FilterRoyaltyUpdated(opts *bind.FilterOpts) (*ERC1155NFTCustomRoyaltyUpdatedIterator, error) {

	logs, err := _ERC1155NFTCustom.contract.FilterLogs(opts, "RoyaltyUpdated")
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomRoyaltyUpdatedIterator{contract: _ERC1155NFTCustom.contract, event: "RoyaltyUpdated", logs: logs}, nil
}

// WatchRoyaltyUpdated is a free log subscription operation binding the contract event 0x13ceafb5e8ec39102bd452913a5c00a05f3b060d636d3a567c1c80e8b4321fe7.
//
// Solidity: event RoyaltyUpdated(uint256 royaltiesBps, address royaltiesAddress)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) WatchRoyaltyUpdated(opts *bind.WatchOpts, sink chan<- *ERC1155NFTCustomRoyaltyUpdatedOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _ERC1155NFTCustom.contract.WatchLogs(opts, "RoyaltyUpdated")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTCustomRoyaltyUpdatedOrChainReorg)
				event.Event = new(ERC1155NFTCustomRoyaltyUpdated)

				if log.ChainReorg == nil {
					if err := _ERC1155NFTCustom.contract.UnpackLog(event.Event, "RoyaltyUpdated", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoyaltyUpdated is a log parse operation binding the contract event 0x13ceafb5e8ec39102bd452913a5c00a05f3b060d636d3a567c1c80e8b4321fe7.
//
// Solidity: event RoyaltyUpdated(uint256 royaltiesBps, address royaltiesAddress)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) ParseRoyaltyUpdated(log types.Log) (*ERC1155NFTCustomRoyaltyUpdated, error) {
	event := new(ERC1155NFTCustomRoyaltyUpdated)
	if err := _ERC1155NFTCustom.contract.UnpackLog(event, "RoyaltyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTCustomTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomTransferBatchIterator struct {
	Event *ERC1155NFTCustomTransferBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155NFTCustomTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTCustomTransferBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155NFTCustomTransferBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155NFTCustomTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTCustomTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTCustomTransferBatch represents a TransferBatch event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// ERC1155NFTCustomTransferBatchOrChainReorg represents a TransferBatch subscription event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomTransferBatchOrChainReorg struct {
	Event      *ERC1155NFTCustomTransferBatch
	ChainReorg *types.ChainReorg
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ERC1155NFTCustomTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, err := _ERC1155NFTCustom.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomTransferBatchIterator{contract: _ERC1155NFTCustom.contract, event: "TransferBatch", logs: logs}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ERC1155NFTCustomTransferBatchOrChainReorg, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155NFTCustom.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTCustomTransferBatchOrChainReorg)
				event.Event = new(ERC1155NFTCustomTransferBatch)

				if log.ChainReorg == nil {
					if err := _ERC1155NFTCustom.contract.UnpackLog(event.Event, "TransferBatch", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) ParseTransferBatch(log types.Log) (*ERC1155NFTCustomTransferBatch, error) {
	event := new(ERC1155NFTCustomTransferBatch)
	if err := _ERC1155NFTCustom.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTCustomTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomTransferSingleIterator struct {
	Event *ERC1155NFTCustomTransferSingle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155NFTCustomTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTCustomTransferSingle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155NFTCustomTransferSingle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155NFTCustomTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTCustomTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTCustomTransferSingle represents a TransferSingle event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// ERC1155NFTCustomTransferSingleOrChainReorg represents a TransferSingle subscription event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomTransferSingleOrChainReorg struct {
	Event      *ERC1155NFTCustomTransferSingle
	ChainReorg *types.ChainReorg
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ERC1155NFTCustomTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, err := _ERC1155NFTCustom.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomTransferSingleIterator{contract: _ERC1155NFTCustom.contract, event: "TransferSingle", logs: logs}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ERC1155NFTCustomTransferSingleOrChainReorg, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155NFTCustom.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTCustomTransferSingleOrChainReorg)
				event.Event = new(ERC1155NFTCustomTransferSingle)

				if log.ChainReorg == nil {
					if err := _ERC1155NFTCustom.contract.UnpackLog(event.Event, "TransferSingle", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) ParseTransferSingle(log types.Log) (*ERC1155NFTCustomTransferSingle, error) {
	event := new(ERC1155NFTCustomTransferSingle)
	if err := _ERC1155NFTCustom.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTCustomTransferableChangedIterator is returned from FilterTransferableChanged and is used to iterate over the raw logs and unpacked data for TransferableChanged events raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomTransferableChangedIterator struct {
	Event *ERC1155NFTCustomTransferableChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155NFTCustomTransferableChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTCustomTransferableChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155NFTCustomTransferableChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155NFTCustomTransferableChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTCustomTransferableChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTCustomTransferableChanged represents a TransferableChanged event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomTransferableChanged struct {
	TransferableByAdmin bool
	TransferableByUser  bool
	Raw                 types.Log // Blockchain specific contextual infos
}

// ERC1155NFTCustomTransferableChangedOrChainReorg represents a TransferableChanged subscription event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomTransferableChangedOrChainReorg struct {
	Event      *ERC1155NFTCustomTransferableChanged
	ChainReorg *types.ChainReorg
}

// FilterTransferableChanged is a free log retrieval operation binding the contract event 0xab67e71edde643937fe4eedd295209d1ea4a844ff0f9203a60e651dd3e46687f.
//
// Solidity: event TransferableChanged(bool transferableByAdmin, bool transferableByUser)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) FilterTransferableChanged(opts *bind.FilterOpts) (*ERC1155NFTCustomTransferableChangedIterator, error) {

	logs, err := _ERC1155NFTCustom.contract.FilterLogs(opts, "TransferableChanged")
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomTransferableChangedIterator{contract: _ERC1155NFTCustom.contract, event: "TransferableChanged", logs: logs}, nil
}

// WatchTransferableChanged is a free log subscription operation binding the contract event 0xab67e71edde643937fe4eedd295209d1ea4a844ff0f9203a60e651dd3e46687f.
//
// Solidity: event TransferableChanged(bool transferableByAdmin, bool transferableByUser)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) WatchTransferableChanged(opts *bind.WatchOpts, sink chan<- *ERC1155NFTCustomTransferableChangedOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _ERC1155NFTCustom.contract.WatchLogs(opts, "TransferableChanged")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTCustomTransferableChangedOrChainReorg)
				event.Event = new(ERC1155NFTCustomTransferableChanged)

				if log.ChainReorg == nil {
					if err := _ERC1155NFTCustom.contract.UnpackLog(event.Event, "TransferableChanged", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferableChanged is a log parse operation binding the contract event 0xab67e71edde643937fe4eedd295209d1ea4a844ff0f9203a60e651dd3e46687f.
//
// Solidity: event TransferableChanged(bool transferableByAdmin, bool transferableByUser)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) ParseTransferableChanged(log types.Log) (*ERC1155NFTCustomTransferableChanged, error) {
	event := new(ERC1155NFTCustomTransferableChanged)
	if err := _ERC1155NFTCustom.contract.UnpackLog(event, "TransferableChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTCustomURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomURIIterator struct {
	Event *ERC1155NFTCustomURI // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155NFTCustomURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTCustomURI)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155NFTCustomURI)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155NFTCustomURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTCustomURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTCustomURI represents a URI event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// ERC1155NFTCustomURIOrChainReorg represents a URI subscription event raised by the ERC1155NFTCustom contract.
type ERC1155NFTCustomURIOrChainReorg struct {
	Event      *ERC1155NFTCustomURI
	ChainReorg *types.ChainReorg
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ERC1155NFTCustomURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, err := _ERC1155NFTCustom.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCustomURIIterator{contract: _ERC1155NFTCustom.contract, event: "URI", logs: logs}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ERC1155NFTCustomURIOrChainReorg, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC1155NFTCustom.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTCustomURIOrChainReorg)
				event.Event = new(ERC1155NFTCustomURI)

				if log.ChainReorg == nil {
					if err := _ERC1155NFTCustom.contract.UnpackLog(event.Event, "URI", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155NFTCustom *ERC1155NFTCustomFilterer) ParseURI(log types.Log) (*ERC1155NFTCustomURI, error) {
	event := new(ERC1155NFTCustomURI)
	if err := _ERC1155NFTCustom.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
