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

// ERC721NFTCustomABI is the input ABI used to generate the binding from.
const ERC721NFTCustomABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"PermanentURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PermanentURIGlobal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltiesBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"royaltiesAddress\",\"type\":\"address\"}],\"name\":\"RoyaltyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferCooldownTimeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"transferableByAdmin\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"transferableByUser\",\"type\":\"bool\"}],\"name\":\"TransferableChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"whites\",\"type\":\"address[]\"}],\"name\":\"addSponsorPrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freezeGlobalMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"freezeTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"freezeTokenUris\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"grantAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"grantMintRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"royaltiesBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"royaltiesAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"tokensTransferableByAdmin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"tokensTransferableByUser\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"transferCooldownTime_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSetSponsorWhitelistForAllUser\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastTransferTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listSponsorPrivilege\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataUpdatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tos\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"}],\"name\":\"mintToBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"whites\",\"type\":\"address[]\"}],\"name\":\"removeSponsorPrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltiesAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltiesBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_royaltiesBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_royaltiesAddress\",\"type\":\"address\"}],\"name\":\"setRoyalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"transferableByAdmin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"transferableByUser\",\"type\":\"bool\"}],\"name\":\"setTokensTransferable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setTransferCooldownTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensBurnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"tokensOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensTransferableByAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensTransferableByUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"transferBatchByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"transferByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferCooldownTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"updateTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721NFTCustom is an auto generated Go binding around an Conflux contract.
type ERC721NFTCustom struct {
	ERC721NFTCustomCaller     // Read-only binding to the contract
	ERC721NFTCustomTransactor // Write-only binding to the contract
	ERC721NFTCustomFilterer   // Log filterer for contract events
}

// ERC721NFTCustomCaller is an auto generated read-only Go binding around an Conflux contract.
type ERC721NFTCustomCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721NFTCustomBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type ERC721NFTCustomBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721NFTCustomTransactor is an auto generated write-only Go binding around an Conflux contract.
type ERC721NFTCustomTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721NFTCustomBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type ERC721NFTCustomBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721NFTCustomFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type ERC721NFTCustomFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721NFTCustomSession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type ERC721NFTCustomSession struct {
	Contract     *ERC721NFTCustom  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721NFTCustomCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type ERC721NFTCustomCallerSession struct {
	Contract *ERC721NFTCustomCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ERC721NFTCustomTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type ERC721NFTCustomTransactorSession struct {
	Contract     *ERC721NFTCustomTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC721NFTCustomRaw is an auto generated low-level Go binding around an Conflux contract.
type ERC721NFTCustomRaw struct {
	Contract *ERC721NFTCustom // Generic contract binding to access the raw methods on
}

// ERC721NFTCustomCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type ERC721NFTCustomCallerRaw struct {
	Contract *ERC721NFTCustomCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721NFTCustomTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type ERC721NFTCustomTransactorRaw struct {
	Contract *ERC721NFTCustomTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721NFTCustom creates a new instance of ERC721NFTCustom, bound to a specific deployed contract.
func NewERC721NFTCustom(address types.Address, backend bind.ContractBackend) (*ERC721NFTCustom, error) {
	contract, err := bindERC721NFTCustom(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustom{ERC721NFTCustomCaller: ERC721NFTCustomCaller{contract: contract}, ERC721NFTCustomTransactor: ERC721NFTCustomTransactor{contract: contract}, ERC721NFTCustomFilterer: ERC721NFTCustomFilterer{contract: contract}}, nil
}

// NewERC721NFTCustomCaller creates a new read-only instance of ERC721NFTCustom, bound to a specific deployed contract.
func NewERC721NFTCustomCaller(address types.Address, caller bind.ContractCaller) (*ERC721NFTCustomCaller, error) {
	contract, err := bindERC721NFTCustom(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomCaller{contract: contract}, nil
}

// NewERC721NFTCustomTransactor creates a new write-only instance of ERC721NFTCustom, bound to a specific deployed contract.
func NewERC721NFTCustomTransactor(address types.Address, transactor bind.ContractTransactor) (*ERC721NFTCustomTransactor, error) {
	contract, err := bindERC721NFTCustom(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomTransactor{contract: contract}, nil
}

// NewERC721NFTCustomFilterer creates a new log filterer instance of ERC721NFTCustom, bound to a specific deployed contract.
func NewERC721NFTCustomFilterer(address types.Address, filterer bind.ContractFilterer) (*ERC721NFTCustomFilterer, error) {
	contract, err := bindERC721NFTCustom(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomFilterer{contract: contract}, nil
}

// NewERC721NFTCustomCaller creates a new read-only instance of ERC721NFTCustom, bound to a specific deployed contract.
func NewERC721NFTCustomBulkCaller(address types.Address, caller bind.ContractCaller) (*ERC721NFTCustomBulkCaller, error) {
	contract, err := bindERC721NFTCustom(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomBulkCaller{contract: contract}, nil
}

// NewERC721NFTCustomBulkTransactor creates a new write-only instance of ERC721NFTCustom, bound to a specific deployed contract.
func NewERC721NFTCustomBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*ERC721NFTCustomBulkTransactor, error) {
	contract, err := bindERC721NFTCustom(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomBulkTransactor{contract: contract}, nil
}

// bindERC721NFTCustom binds a generic wrapper to an already deployed contract.
func bindERC721NFTCustom(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721NFTCustomABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721NFTCustom *ERC721NFTCustomRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721NFTCustom.Contract.ERC721NFTCustomCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721NFTCustom *ERC721NFTCustomRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.ERC721NFTCustomTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721NFTCustom *ERC721NFTCustomRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.ERC721NFTCustomTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721NFTCustom *ERC721NFTCustomCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721NFTCustom.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721NFTCustom *ERC721NFTCustomTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721NFTCustom *ERC721NFTCustomTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) DEFAULTADMINROLE(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "DEFAULT_ADMIN_ROLE")

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "DEFAULT_ADMIN_ROLE")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC721NFTCustom.Contract.DEFAULTADMINROLE(&_ERC721NFTCustom.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC721NFTCustom.Contract.DEFAULTADMINROLE(&_ERC721NFTCustom.CallOpts)
}

// ZERO is a free data retrieval call binding the contract method 0x58fa63ca.
//
// Solidity: function ZERO() view returns(address)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) ZERO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "ZERO")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// ZERO is a free data retrieval call binding the contract method 0x58fa63ca.
//
// Solidity: function ZERO() view returns(address)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) ZERO(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "ZERO")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "ZERO")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) ZERO() (common.Address, error) {
	return _ERC721NFTCustom.Contract.ZERO(&_ERC721NFTCustom.CallOpts)
}

// ZERO is a free data retrieval call binding the contract method 0x58fa63ca.
//
// Solidity: function ZERO() view returns(address)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) ZERO() (common.Address, error) {
	return _ERC721NFTCustom.Contract.ZERO(&_ERC721NFTCustom.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "balanceOf", owner)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) BalanceOf(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, owner common.Address) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "balanceOf", owner)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "balanceOf")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721NFTCustom.Contract.BalanceOf(&_ERC721NFTCustom.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721NFTCustom.Contract.BalanceOf(&_ERC721NFTCustom.CallOpts, owner)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "contractURI")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) ContractURI(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "contractURI")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "contractURI")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) ContractURI() (string, error) {
	return _ERC721NFTCustom.Contract.ContractURI(&_ERC721NFTCustom.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) ContractURI() (string, error) {
	return _ERC721NFTCustom.Contract.ContractURI(&_ERC721NFTCustom.CallOpts)
}

// FreezeTokenUris is a free data retrieval call binding the contract method 0x8d010db3.
//
// Solidity: function freezeTokenUris(uint256 ) view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) FreezeTokenUris(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "freezeTokenUris", arg0)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// FreezeTokenUris is a free data retrieval call binding the contract method 0x8d010db3.
//
// Solidity: function freezeTokenUris(uint256 ) view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) FreezeTokenUris(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, arg0 *big.Int) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "freezeTokenUris", arg0)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "freezeTokenUris")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) FreezeTokenUris(arg0 *big.Int) (bool, error) {
	return _ERC721NFTCustom.Contract.FreezeTokenUris(&_ERC721NFTCustom.CallOpts, arg0)
}

// FreezeTokenUris is a free data retrieval call binding the contract method 0x8d010db3.
//
// Solidity: function freezeTokenUris(uint256 ) view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) FreezeTokenUris(arg0 *big.Int) (bool, error) {
	return _ERC721NFTCustom.Contract.FreezeTokenUris(&_ERC721NFTCustom.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "getApproved", tokenId)

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) GetApproved(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, tokenId *big.Int) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "getApproved", tokenId)

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "getApproved")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721NFTCustom *ERC721NFTCustomSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721NFTCustom.Contract.GetApproved(&_ERC721NFTCustom.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721NFTCustom.Contract.GetApproved(&_ERC721NFTCustom.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "getRoleAdmin", role)

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) GetRoleAdmin(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, role [32]byte) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "getRoleAdmin", role)

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "getRoleAdmin")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC721NFTCustom.Contract.GetRoleAdmin(&_ERC721NFTCustom.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC721NFTCustom.Contract.GetRoleAdmin(&_ERC721NFTCustom.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "hasRole", role, account)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) HasRole(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, role [32]byte, account common.Address) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "hasRole", role, account)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "hasRole")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC721NFTCustom.Contract.HasRole(&_ERC721NFTCustom.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC721NFTCustom.Contract.HasRole(&_ERC721NFTCustom.CallOpts, role, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0xb6db75a0.
//
// Solidity: function isAdmin() view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) IsAdmin(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "isAdmin")

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// IsAdmin is a free data retrieval call binding the contract method 0xb6db75a0.
//
// Solidity: function isAdmin() view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) IsAdmin(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "isAdmin")

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "isAdmin")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) IsAdmin() (bool, error) {
	return _ERC721NFTCustom.Contract.IsAdmin(&_ERC721NFTCustom.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0xb6db75a0.
//
// Solidity: function isAdmin() view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) IsAdmin() (bool, error) {
	return _ERC721NFTCustom.Contract.IsAdmin(&_ERC721NFTCustom.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) IsApprovedForAll(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, owner common.Address, operator common.Address) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "isApprovedForAll", owner, operator)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "isApprovedForAll")
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
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721NFTCustom.Contract.IsApprovedForAll(&_ERC721NFTCustom.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721NFTCustom.Contract.IsApprovedForAll(&_ERC721NFTCustom.CallOpts, owner, operator)
}

// LastTransferTimes is a free data retrieval call binding the contract method 0x0e66eac5.
//
// Solidity: function lastTransferTimes(uint256 ) view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) LastTransferTimes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "lastTransferTimes", arg0)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// LastTransferTimes is a free data retrieval call binding the contract method 0x0e66eac5.
//
// Solidity: function lastTransferTimes(uint256 ) view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) LastTransferTimes(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, arg0 *big.Int) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "lastTransferTimes", arg0)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "lastTransferTimes")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// LastTransferTimes is a free data retrieval call binding the contract method 0x0e66eac5.
//
// Solidity: function lastTransferTimes(uint256 ) view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomSession) LastTransferTimes(arg0 *big.Int) (*big.Int, error) {
	return _ERC721NFTCustom.Contract.LastTransferTimes(&_ERC721NFTCustom.CallOpts, arg0)
}

// LastTransferTimes is a free data retrieval call binding the contract method 0x0e66eac5.
//
// Solidity: function lastTransferTimes(uint256 ) view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) LastTransferTimes(arg0 *big.Int) (*big.Int, error) {
	return _ERC721NFTCustom.Contract.LastTransferTimes(&_ERC721NFTCustom.CallOpts, arg0)
}

// ListSponsorPrivilege is a free data retrieval call binding the contract method 0x28630e1d.
//
// Solidity: function listSponsorPrivilege() view returns(address[])
func (_ERC721NFTCustom *ERC721NFTCustomCaller) ListSponsorPrivilege(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "listSponsorPrivilege")

	if __err != nil {
		return *new([]common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, __err

}

// ListSponsorPrivilege is a free data retrieval call binding the contract method 0x28630e1d.
//
// Solidity: function listSponsorPrivilege() view returns(address[])
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) ListSponsorPrivilege(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*[]common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "listSponsorPrivilege")

	out0 := new([]common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "listSponsorPrivilege")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) ListSponsorPrivilege() ([]common.Address, error) {
	return _ERC721NFTCustom.Contract.ListSponsorPrivilege(&_ERC721NFTCustom.CallOpts)
}

// ListSponsorPrivilege is a free data retrieval call binding the contract method 0x28630e1d.
//
// Solidity: function listSponsorPrivilege() view returns(address[])
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) ListSponsorPrivilege() ([]common.Address, error) {
	return _ERC721NFTCustom.Contract.ListSponsorPrivilege(&_ERC721NFTCustom.CallOpts)
}

// MetadataUpdatable is a free data retrieval call binding the contract method 0x4e6f9dd6.
//
// Solidity: function metadataUpdatable() view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) MetadataUpdatable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "metadataUpdatable")

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// MetadataUpdatable is a free data retrieval call binding the contract method 0x4e6f9dd6.
//
// Solidity: function metadataUpdatable() view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) MetadataUpdatable(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "metadataUpdatable")

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "metadataUpdatable")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) MetadataUpdatable() (bool, error) {
	return _ERC721NFTCustom.Contract.MetadataUpdatable(&_ERC721NFTCustom.CallOpts)
}

// MetadataUpdatable is a free data retrieval call binding the contract method 0x4e6f9dd6.
//
// Solidity: function metadataUpdatable() view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) MetadataUpdatable() (bool, error) {
	return _ERC721NFTCustom.Contract.MetadataUpdatable(&_ERC721NFTCustom.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "name")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) Name(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "name")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "name")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) Name() (string, error) {
	return _ERC721NFTCustom.Contract.Name(&_ERC721NFTCustom.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) Name() (string, error) {
	return _ERC721NFTCustom.Contract.Name(&_ERC721NFTCustom.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "ownerOf", tokenId)

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) OwnerOf(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, tokenId *big.Int) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "ownerOf", tokenId)

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "ownerOf")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721NFTCustom *ERC721NFTCustomSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721NFTCustom.Contract.OwnerOf(&_ERC721NFTCustom.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721NFTCustom.Contract.OwnerOf(&_ERC721NFTCustom.CallOpts, tokenId)
}

// RoyaltiesAddress is a free data retrieval call binding the contract method 0x32882535.
//
// Solidity: function royaltiesAddress() view returns(address)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) RoyaltiesAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "royaltiesAddress")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// RoyaltiesAddress is a free data retrieval call binding the contract method 0x32882535.
//
// Solidity: function royaltiesAddress() view returns(address)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) RoyaltiesAddress(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "royaltiesAddress")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "royaltiesAddress")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) RoyaltiesAddress() (common.Address, error) {
	return _ERC721NFTCustom.Contract.RoyaltiesAddress(&_ERC721NFTCustom.CallOpts)
}

// RoyaltiesAddress is a free data retrieval call binding the contract method 0x32882535.
//
// Solidity: function royaltiesAddress() view returns(address)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) RoyaltiesAddress() (common.Address, error) {
	return _ERC721NFTCustom.Contract.RoyaltiesAddress(&_ERC721NFTCustom.CallOpts)
}

// RoyaltiesBps is a free data retrieval call binding the contract method 0x99d89f9d.
//
// Solidity: function royaltiesBps() view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) RoyaltiesBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "royaltiesBps")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// RoyaltiesBps is a free data retrieval call binding the contract method 0x99d89f9d.
//
// Solidity: function royaltiesBps() view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) RoyaltiesBps(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "royaltiesBps")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "royaltiesBps")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) RoyaltiesBps() (*big.Int, error) {
	return _ERC721NFTCustom.Contract.RoyaltiesBps(&_ERC721NFTCustom.CallOpts)
}

// RoyaltiesBps is a free data retrieval call binding the contract method 0x99d89f9d.
//
// Solidity: function royaltiesBps() view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) RoyaltiesBps() (*big.Int, error) {
	return _ERC721NFTCustom.Contract.RoyaltiesBps(&_ERC721NFTCustom.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

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
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) RoyaltyInfo(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (*common.Address, **big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "royaltyInfo", tokenId, salePrice)

	out0 := new(common.Address)
	out1 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "royaltyInfo")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _ERC721NFTCustom.Contract.RoyaltyInfo(&_ERC721NFTCustom.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _ERC721NFTCustom.Contract.RoyaltyInfo(&_ERC721NFTCustom.CallOpts, tokenId, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) SupportsInterface(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, interfaceId [4]byte) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "supportsInterface", interfaceId)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "supportsInterface")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721NFTCustom.Contract.SupportsInterface(&_ERC721NFTCustom.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721NFTCustom.Contract.SupportsInterface(&_ERC721NFTCustom.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "symbol")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) Symbol(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "symbol")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "symbol")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) Symbol() (string, error) {
	return _ERC721NFTCustom.Contract.Symbol(&_ERC721NFTCustom.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) Symbol() (string, error) {
	return _ERC721NFTCustom.Contract.Symbol(&_ERC721NFTCustom.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "tokenByIndex", index)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) TokenByIndex(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, index *big.Int) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "tokenByIndex", index)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "tokenByIndex")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721NFTCustom.Contract.TokenByIndex(&_ERC721NFTCustom.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721NFTCustom.Contract.TokenByIndex(&_ERC721NFTCustom.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) TokenOfOwnerByIndex(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, owner common.Address, index *big.Int) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "tokenOfOwnerByIndex", owner, index)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "tokenOfOwnerByIndex")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721NFTCustom.Contract.TokenOfOwnerByIndex(&_ERC721NFTCustom.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721NFTCustom.Contract.TokenOfOwnerByIndex(&_ERC721NFTCustom.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "tokenURI", tokenId)

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) TokenURI(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, tokenId *big.Int) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "tokenURI", tokenId)

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "tokenURI")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721NFTCustom *ERC721NFTCustomSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721NFTCustom.Contract.TokenURI(&_ERC721NFTCustom.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721NFTCustom.Contract.TokenURI(&_ERC721NFTCustom.CallOpts, tokenId)
}

// Tokens is a free data retrieval call binding the contract method 0x8b4864d6.
//
// Solidity: function tokens(uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) Tokens(opts *bind.CallOpts, offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "tokens", offset, limit)

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
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) Tokens(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, offset *big.Int, limit *big.Int) (*struct {
	Total    *big.Int
	TokenIds []*big.Int
}, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "tokens", offset, limit)

	outstruct := new(struct {
		Total    *big.Int
		TokenIds []*big.Int
	})

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "tokens")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) Tokens(offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	return _ERC721NFTCustom.Contract.Tokens(&_ERC721NFTCustom.CallOpts, offset, limit)
}

// Tokens is a free data retrieval call binding the contract method 0x8b4864d6.
//
// Solidity: function tokens(uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) Tokens(offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	return _ERC721NFTCustom.Contract.Tokens(&_ERC721NFTCustom.CallOpts, offset, limit)
}

// TokensBurnable is a free data retrieval call binding the contract method 0xe3d52072.
//
// Solidity: function tokensBurnable() view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) TokensBurnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "tokensBurnable")

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// TokensBurnable is a free data retrieval call binding the contract method 0xe3d52072.
//
// Solidity: function tokensBurnable() view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) TokensBurnable(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "tokensBurnable")

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "tokensBurnable")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) TokensBurnable() (bool, error) {
	return _ERC721NFTCustom.Contract.TokensBurnable(&_ERC721NFTCustom.CallOpts)
}

// TokensBurnable is a free data retrieval call binding the contract method 0xe3d52072.
//
// Solidity: function tokensBurnable() view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) TokensBurnable() (bool, error) {
	return _ERC721NFTCustom.Contract.TokensBurnable(&_ERC721NFTCustom.CallOpts)
}

// TokensOf is a free data retrieval call binding the contract method 0x23185dc9.
//
// Solidity: function tokensOf(address owner, uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) TokensOf(opts *bind.CallOpts, owner common.Address, offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "tokensOf", owner, offset, limit)

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
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) TokensOf(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, owner common.Address, offset *big.Int, limit *big.Int) (*struct {
	Total    *big.Int
	TokenIds []*big.Int
}, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "tokensOf", owner, offset, limit)

	outstruct := new(struct {
		Total    *big.Int
		TokenIds []*big.Int
	})

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "tokensOf")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) TokensOf(owner common.Address, offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	return _ERC721NFTCustom.Contract.TokensOf(&_ERC721NFTCustom.CallOpts, owner, offset, limit)
}

// TokensOf is a free data retrieval call binding the contract method 0x23185dc9.
//
// Solidity: function tokensOf(address owner, uint256 offset, uint256 limit) view returns(uint256 total, uint256[] tokenIds)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) TokensOf(owner common.Address, offset *big.Int, limit *big.Int) (struct {
	Total    *big.Int
	TokenIds []*big.Int
}, error) {
	return _ERC721NFTCustom.Contract.TokensOf(&_ERC721NFTCustom.CallOpts, owner, offset, limit)
}

// TokensTransferableByAdmin is a free data retrieval call binding the contract method 0xd9fab275.
//
// Solidity: function tokensTransferableByAdmin() view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) TokensTransferableByAdmin(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "tokensTransferableByAdmin")

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// TokensTransferableByAdmin is a free data retrieval call binding the contract method 0xd9fab275.
//
// Solidity: function tokensTransferableByAdmin() view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) TokensTransferableByAdmin(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "tokensTransferableByAdmin")

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "tokensTransferableByAdmin")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) TokensTransferableByAdmin() (bool, error) {
	return _ERC721NFTCustom.Contract.TokensTransferableByAdmin(&_ERC721NFTCustom.CallOpts)
}

// TokensTransferableByAdmin is a free data retrieval call binding the contract method 0xd9fab275.
//
// Solidity: function tokensTransferableByAdmin() view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) TokensTransferableByAdmin() (bool, error) {
	return _ERC721NFTCustom.Contract.TokensTransferableByAdmin(&_ERC721NFTCustom.CallOpts)
}

// TokensTransferableByUser is a free data retrieval call binding the contract method 0x7915c570.
//
// Solidity: function tokensTransferableByUser() view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) TokensTransferableByUser(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "tokensTransferableByUser")

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// TokensTransferableByUser is a free data retrieval call binding the contract method 0x7915c570.
//
// Solidity: function tokensTransferableByUser() view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) TokensTransferableByUser(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "tokensTransferableByUser")

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "tokensTransferableByUser")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) TokensTransferableByUser() (bool, error) {
	return _ERC721NFTCustom.Contract.TokensTransferableByUser(&_ERC721NFTCustom.CallOpts)
}

// TokensTransferableByUser is a free data retrieval call binding the contract method 0x7915c570.
//
// Solidity: function tokensTransferableByUser() view returns(bool)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) TokensTransferableByUser() (bool, error) {
	return _ERC721NFTCustom.Contract.TokensTransferableByUser(&_ERC721NFTCustom.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "totalSupply")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) TotalSupply(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "totalSupply")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "totalSupply")
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
func (_ERC721NFTCustom *ERC721NFTCustomSession) TotalSupply() (*big.Int, error) {
	return _ERC721NFTCustom.Contract.TotalSupply(&_ERC721NFTCustom.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721NFTCustom.Contract.TotalSupply(&_ERC721NFTCustom.CallOpts)
}

// TransferCooldownTime is a free data retrieval call binding the contract method 0xd32a81ab.
//
// Solidity: function transferCooldownTime() view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCaller) TransferCooldownTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _ERC721NFTCustom.contract.Call(opts, &out, "transferCooldownTime")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// TransferCooldownTime is a free data retrieval call binding the contract method 0xd32a81ab.
//
// Solidity: function transferCooldownTime() view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomBulkCaller) TransferCooldownTime(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC721NFTCustom.contract.GenRequest(opts, "transferCooldownTime")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC721NFTCustom.contract.DecodeOutput(&out, rawOut, "transferCooldownTime")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TransferCooldownTime is a free data retrieval call binding the contract method 0xd32a81ab.
//
// Solidity: function transferCooldownTime() view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomSession) TransferCooldownTime() (*big.Int, error) {
	return _ERC721NFTCustom.Contract.TransferCooldownTime(&_ERC721NFTCustom.CallOpts)
}

// TransferCooldownTime is a free data retrieval call binding the contract method 0xd32a81ab.
//
// Solidity: function transferCooldownTime() view returns(uint256)
func (_ERC721NFTCustom *ERC721NFTCustomCallerSession) TransferCooldownTime() (*big.Int, error) {
	return _ERC721NFTCustom.Contract.TransferCooldownTime(&_ERC721NFTCustom.CallOpts)
}

// AddSponsorPrivilege is a paid mutator transaction binding the contract method 0x938ead7c.
//
// Solidity: function addSponsorPrivilege(address[] whites) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) AddSponsorPrivilege(opts *bind.TransactOpts, whites []common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "addSponsorPrivilege", whites)
}

// AddSponsorPrivilege is a paid mutator transaction binding the contract method 0x938ead7c.
//
// Solidity: function addSponsorPrivilege(address[] whites) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) AddSponsorPrivilege(opts *bind.TransactOpts, whites []common.Address) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "addSponsorPrivilege", whites)
}

// AddSponsorPrivilege is a paid mutator transaction binding the contract method 0x938ead7c.
//
// Solidity: function addSponsorPrivilege(address[] whites) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) AddSponsorPrivilege(whites []common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.AddSponsorPrivilege(&_ERC721NFTCustom.TransactOpts, whites)
}

// AddSponsorPrivilege is a paid mutator transaction binding the contract method 0x938ead7c.
//
// Solidity: function addSponsorPrivilege(address[] whites) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) AddSponsorPrivilege(whites []common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.AddSponsorPrivilege(&_ERC721NFTCustom.TransactOpts, whites)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) Approve(to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.Approve(&_ERC721NFTCustom.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.Approve(&_ERC721NFTCustom.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 id) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) Burn(opts *bind.TransactOpts, id *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "burn", id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 id) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) Burn(opts *bind.TransactOpts, id *big.Int) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "burn", id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 id) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) Burn(id *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.Burn(&_ERC721NFTCustom.TransactOpts, id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 id) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) Burn(id *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.Burn(&_ERC721NFTCustom.TransactOpts, id)
}

// BurnBatch is a paid mutator transaction binding the contract method 0xe4623c1b.
//
// Solidity: function burnBatch(uint256[] ids) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) BurnBatch(opts *bind.TransactOpts, ids []*big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "burnBatch", ids)
}

// BurnBatch is a paid mutator transaction binding the contract method 0xe4623c1b.
//
// Solidity: function burnBatch(uint256[] ids) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) BurnBatch(opts *bind.TransactOpts, ids []*big.Int) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "burnBatch", ids)
}

// BurnBatch is a paid mutator transaction binding the contract method 0xe4623c1b.
//
// Solidity: function burnBatch(uint256[] ids) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) BurnBatch(ids []*big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.BurnBatch(&_ERC721NFTCustom.TransactOpts, ids)
}

// BurnBatch is a paid mutator transaction binding the contract method 0xe4623c1b.
//
// Solidity: function burnBatch(uint256[] ids) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) BurnBatch(ids []*big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.BurnBatch(&_ERC721NFTCustom.TransactOpts, ids)
}

// FreezeGlobalMetadata is a paid mutator transaction binding the contract method 0x092e7106.
//
// Solidity: function freezeGlobalMetadata() returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) FreezeGlobalMetadata(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "freezeGlobalMetadata")
}

// FreezeGlobalMetadata is a paid mutator transaction binding the contract method 0x092e7106.
//
// Solidity: function freezeGlobalMetadata() returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) FreezeGlobalMetadata(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "freezeGlobalMetadata")
}

// FreezeGlobalMetadata is a paid mutator transaction binding the contract method 0x092e7106.
//
// Solidity: function freezeGlobalMetadata() returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) FreezeGlobalMetadata() (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.FreezeGlobalMetadata(&_ERC721NFTCustom.TransactOpts)
}

// FreezeGlobalMetadata is a paid mutator transaction binding the contract method 0x092e7106.
//
// Solidity: function freezeGlobalMetadata() returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) FreezeGlobalMetadata() (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.FreezeGlobalMetadata(&_ERC721NFTCustom.TransactOpts)
}

// FreezeTokenURI is a paid mutator transaction binding the contract method 0x385c0eb0.
//
// Solidity: function freezeTokenURI(uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) FreezeTokenURI(opts *bind.TransactOpts, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "freezeTokenURI", tokenId)
}

// FreezeTokenURI is a paid mutator transaction binding the contract method 0x385c0eb0.
//
// Solidity: function freezeTokenURI(uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) FreezeTokenURI(opts *bind.TransactOpts, tokenId *big.Int) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "freezeTokenURI", tokenId)
}

// FreezeTokenURI is a paid mutator transaction binding the contract method 0x385c0eb0.
//
// Solidity: function freezeTokenURI(uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) FreezeTokenURI(tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.FreezeTokenURI(&_ERC721NFTCustom.TransactOpts, tokenId)
}

// FreezeTokenURI is a paid mutator transaction binding the contract method 0x385c0eb0.
//
// Solidity: function freezeTokenURI(uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) FreezeTokenURI(tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.FreezeTokenURI(&_ERC721NFTCustom.TransactOpts, tokenId)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address user) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) GrantAdminRole(opts *bind.TransactOpts, user common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "grantAdminRole", user)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address user) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) GrantAdminRole(opts *bind.TransactOpts, user common.Address) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "grantAdminRole", user)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address user) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) GrantAdminRole(user common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.GrantAdminRole(&_ERC721NFTCustom.TransactOpts, user)
}

// GrantAdminRole is a paid mutator transaction binding the contract method 0xc634b78e.
//
// Solidity: function grantAdminRole(address user) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) GrantAdminRole(user common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.GrantAdminRole(&_ERC721NFTCustom.TransactOpts, user)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address user) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) GrantMintRole(opts *bind.TransactOpts, user common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "grantMintRole", user)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address user) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) GrantMintRole(opts *bind.TransactOpts, user common.Address) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "grantMintRole", user)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address user) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) GrantMintRole(user common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.GrantMintRole(&_ERC721NFTCustom.TransactOpts, user)
}

// GrantMintRole is a paid mutator transaction binding the contract method 0xc2e3273d.
//
// Solidity: function grantMintRole(address user) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) GrantMintRole(user common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.GrantMintRole(&_ERC721NFTCustom.TransactOpts, user)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) GrantRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.GrantRole(&_ERC721NFTCustom.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.GrantRole(&_ERC721NFTCustom.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xcb257977.
//
// Solidity: function initialize(string name_, string symbol_, string uri, uint256 royaltiesBps, address royaltiesAddress, address[] owners, bool tokensTransferableByAdmin, bool tokensTransferableByUser, uint256 transferCooldownTime_, bool isSetSponsorWhitelistForAllUser) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, uri string, royaltiesBps *big.Int, royaltiesAddress common.Address, owners []common.Address, tokensTransferableByAdmin bool, tokensTransferableByUser bool, transferCooldownTime_ *big.Int, isSetSponsorWhitelistForAllUser bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "initialize", name_, symbol_, uri, royaltiesBps, royaltiesAddress, owners, tokensTransferableByAdmin, tokensTransferableByUser, transferCooldownTime_, isSetSponsorWhitelistForAllUser)
}

// Initialize is a paid mutator transaction binding the contract method 0xcb257977.
//
// Solidity: function initialize(string name_, string symbol_, string uri, uint256 royaltiesBps, address royaltiesAddress, address[] owners, bool tokensTransferableByAdmin, bool tokensTransferableByUser, uint256 transferCooldownTime_, bool isSetSponsorWhitelistForAllUser) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, uri string, royaltiesBps *big.Int, royaltiesAddress common.Address, owners []common.Address, tokensTransferableByAdmin bool, tokensTransferableByUser bool, transferCooldownTime_ *big.Int, isSetSponsorWhitelistForAllUser bool) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "initialize", name_, symbol_, uri, royaltiesBps, royaltiesAddress, owners, tokensTransferableByAdmin, tokensTransferableByUser, transferCooldownTime_, isSetSponsorWhitelistForAllUser)
}

// Initialize is a paid mutator transaction binding the contract method 0xcb257977.
//
// Solidity: function initialize(string name_, string symbol_, string uri, uint256 royaltiesBps, address royaltiesAddress, address[] owners, bool tokensTransferableByAdmin, bool tokensTransferableByUser, uint256 transferCooldownTime_, bool isSetSponsorWhitelistForAllUser) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) Initialize(name_ string, symbol_ string, uri string, royaltiesBps *big.Int, royaltiesAddress common.Address, owners []common.Address, tokensTransferableByAdmin bool, tokensTransferableByUser bool, transferCooldownTime_ *big.Int, isSetSponsorWhitelistForAllUser bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.Initialize(&_ERC721NFTCustom.TransactOpts, name_, symbol_, uri, royaltiesBps, royaltiesAddress, owners, tokensTransferableByAdmin, tokensTransferableByUser, transferCooldownTime_, isSetSponsorWhitelistForAllUser)
}

// Initialize is a paid mutator transaction binding the contract method 0xcb257977.
//
// Solidity: function initialize(string name_, string symbol_, string uri, uint256 royaltiesBps, address royaltiesAddress, address[] owners, bool tokensTransferableByAdmin, bool tokensTransferableByUser, uint256 transferCooldownTime_, bool isSetSponsorWhitelistForAllUser) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) Initialize(name_ string, symbol_ string, uri string, royaltiesBps *big.Int, royaltiesAddress common.Address, owners []common.Address, tokensTransferableByAdmin bool, tokensTransferableByUser bool, transferCooldownTime_ *big.Int, isSetSponsorWhitelistForAllUser bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.Initialize(&_ERC721NFTCustom.TransactOpts, name_, symbol_, uri, royaltiesBps, royaltiesAddress, owners, tokensTransferableByAdmin, tokensTransferableByUser, transferCooldownTime_, isSetSponsorWhitelistForAllUser)
}

// MintTo is a paid mutator transaction binding the contract method 0x9f6ed25f.
//
// Solidity: function mintTo(address to, uint256 id, string tokenUri) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) MintTo(opts *bind.TransactOpts, to common.Address, id *big.Int, tokenUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "mintTo", to, id, tokenUri)
}

// MintTo is a paid mutator transaction binding the contract method 0x9f6ed25f.
//
// Solidity: function mintTo(address to, uint256 id, string tokenUri) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) MintTo(opts *bind.TransactOpts, to common.Address, id *big.Int, tokenUri string) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "mintTo", to, id, tokenUri)
}

// MintTo is a paid mutator transaction binding the contract method 0x9f6ed25f.
//
// Solidity: function mintTo(address to, uint256 id, string tokenUri) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) MintTo(to common.Address, id *big.Int, tokenUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.MintTo(&_ERC721NFTCustom.TransactOpts, to, id, tokenUri)
}

// MintTo is a paid mutator transaction binding the contract method 0x9f6ed25f.
//
// Solidity: function mintTo(address to, uint256 id, string tokenUri) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) MintTo(to common.Address, id *big.Int, tokenUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.MintTo(&_ERC721NFTCustom.TransactOpts, to, id, tokenUri)
}

// MintToBatch is a paid mutator transaction binding the contract method 0x7398551b.
//
// Solidity: function mintToBatch(address[] tos, uint256[] ids, string[] uris) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) MintToBatch(opts *bind.TransactOpts, tos []common.Address, ids []*big.Int, uris []string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "mintToBatch", tos, ids, uris)
}

// MintToBatch is a paid mutator transaction binding the contract method 0x7398551b.
//
// Solidity: function mintToBatch(address[] tos, uint256[] ids, string[] uris) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) MintToBatch(opts *bind.TransactOpts, tos []common.Address, ids []*big.Int, uris []string) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "mintToBatch", tos, ids, uris)
}

// MintToBatch is a paid mutator transaction binding the contract method 0x7398551b.
//
// Solidity: function mintToBatch(address[] tos, uint256[] ids, string[] uris) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) MintToBatch(tos []common.Address, ids []*big.Int, uris []string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.MintToBatch(&_ERC721NFTCustom.TransactOpts, tos, ids, uris)
}

// MintToBatch is a paid mutator transaction binding the contract method 0x7398551b.
//
// Solidity: function mintToBatch(address[] tos, uint256[] ids, string[] uris) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) MintToBatch(tos []common.Address, ids []*big.Int, uris []string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.MintToBatch(&_ERC721NFTCustom.TransactOpts, tos, ids, uris)
}

// RemoveSponsorPrivilege is a paid mutator transaction binding the contract method 0x76104e53.
//
// Solidity: function removeSponsorPrivilege(address[] whites) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) RemoveSponsorPrivilege(opts *bind.TransactOpts, whites []common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "removeSponsorPrivilege", whites)
}

// RemoveSponsorPrivilege is a paid mutator transaction binding the contract method 0x76104e53.
//
// Solidity: function removeSponsorPrivilege(address[] whites) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) RemoveSponsorPrivilege(opts *bind.TransactOpts, whites []common.Address) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "removeSponsorPrivilege", whites)
}

// RemoveSponsorPrivilege is a paid mutator transaction binding the contract method 0x76104e53.
//
// Solidity: function removeSponsorPrivilege(address[] whites) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) RemoveSponsorPrivilege(whites []common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.RemoveSponsorPrivilege(&_ERC721NFTCustom.TransactOpts, whites)
}

// RemoveSponsorPrivilege is a paid mutator transaction binding the contract method 0x76104e53.
//
// Solidity: function removeSponsorPrivilege(address[] whites) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) RemoveSponsorPrivilege(whites []common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.RemoveSponsorPrivilege(&_ERC721NFTCustom.TransactOpts, whites)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) RenounceRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.RenounceRole(&_ERC721NFTCustom.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.RenounceRole(&_ERC721NFTCustom.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) RevokeRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.RevokeRole(&_ERC721NFTCustom.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.RevokeRole(&_ERC721NFTCustom.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.SafeTransferFrom(&_ERC721NFTCustom.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.SafeTransferFrom(&_ERC721NFTCustom.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.SafeTransferFrom0(&_ERC721NFTCustom.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.SafeTransferFrom0(&_ERC721NFTCustom.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) SetApprovalForAll(operator common.Address, approved bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.SetApprovalForAll(&_ERC721NFTCustom.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.SetApprovalForAll(&_ERC721NFTCustom.TransactOpts, operator, approved)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x5de42985.
//
// Solidity: function setRoyalties(uint256 _royaltiesBps, address _royaltiesAddress) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) SetRoyalties(opts *bind.TransactOpts, _royaltiesBps *big.Int, _royaltiesAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "setRoyalties", _royaltiesBps, _royaltiesAddress)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x5de42985.
//
// Solidity: function setRoyalties(uint256 _royaltiesBps, address _royaltiesAddress) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) SetRoyalties(opts *bind.TransactOpts, _royaltiesBps *big.Int, _royaltiesAddress common.Address) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "setRoyalties", _royaltiesBps, _royaltiesAddress)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x5de42985.
//
// Solidity: function setRoyalties(uint256 _royaltiesBps, address _royaltiesAddress) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) SetRoyalties(_royaltiesBps *big.Int, _royaltiesAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.SetRoyalties(&_ERC721NFTCustom.TransactOpts, _royaltiesBps, _royaltiesAddress)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0x5de42985.
//
// Solidity: function setRoyalties(uint256 _royaltiesBps, address _royaltiesAddress) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) SetRoyalties(_royaltiesBps *big.Int, _royaltiesAddress common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.SetRoyalties(&_ERC721NFTCustom.TransactOpts, _royaltiesBps, _royaltiesAddress)
}

// SetTokensTransferable is a paid mutator transaction binding the contract method 0xd4e0456b.
//
// Solidity: function setTokensTransferable(bool transferableByAdmin, bool transferableByUser) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) SetTokensTransferable(opts *bind.TransactOpts, transferableByAdmin bool, transferableByUser bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "setTokensTransferable", transferableByAdmin, transferableByUser)
}

// SetTokensTransferable is a paid mutator transaction binding the contract method 0xd4e0456b.
//
// Solidity: function setTokensTransferable(bool transferableByAdmin, bool transferableByUser) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) SetTokensTransferable(opts *bind.TransactOpts, transferableByAdmin bool, transferableByUser bool) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "setTokensTransferable", transferableByAdmin, transferableByUser)
}

// SetTokensTransferable is a paid mutator transaction binding the contract method 0xd4e0456b.
//
// Solidity: function setTokensTransferable(bool transferableByAdmin, bool transferableByUser) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) SetTokensTransferable(transferableByAdmin bool, transferableByUser bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.SetTokensTransferable(&_ERC721NFTCustom.TransactOpts, transferableByAdmin, transferableByUser)
}

// SetTokensTransferable is a paid mutator transaction binding the contract method 0xd4e0456b.
//
// Solidity: function setTokensTransferable(bool transferableByAdmin, bool transferableByUser) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) SetTokensTransferable(transferableByAdmin bool, transferableByUser bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.SetTokensTransferable(&_ERC721NFTCustom.TransactOpts, transferableByAdmin, transferableByUser)
}

// SetTransferCooldownTime is a paid mutator transaction binding the contract method 0x41a9f4f4.
//
// Solidity: function setTransferCooldownTime(uint256 val) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) SetTransferCooldownTime(opts *bind.TransactOpts, val *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "setTransferCooldownTime", val)
}

// SetTransferCooldownTime is a paid mutator transaction binding the contract method 0x41a9f4f4.
//
// Solidity: function setTransferCooldownTime(uint256 val) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) SetTransferCooldownTime(opts *bind.TransactOpts, val *big.Int) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "setTransferCooldownTime", val)
}

// SetTransferCooldownTime is a paid mutator transaction binding the contract method 0x41a9f4f4.
//
// Solidity: function setTransferCooldownTime(uint256 val) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) SetTransferCooldownTime(val *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.SetTransferCooldownTime(&_ERC721NFTCustom.TransactOpts, val)
}

// SetTransferCooldownTime is a paid mutator transaction binding the contract method 0x41a9f4f4.
//
// Solidity: function setTransferCooldownTime(uint256 val) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) SetTransferCooldownTime(val *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.SetTransferCooldownTime(&_ERC721NFTCustom.TransactOpts, val)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newURI) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) SetURI(opts *bind.TransactOpts, newURI string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "setURI", newURI)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newURI) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) SetURI(opts *bind.TransactOpts, newURI string) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "setURI", newURI)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newURI) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) SetURI(newURI string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.SetURI(&_ERC721NFTCustom.TransactOpts, newURI)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newURI) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) SetURI(newURI string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.SetURI(&_ERC721NFTCustom.TransactOpts, newURI)
}

// TransferBatchByAdmin is a paid mutator transaction binding the contract method 0x159ccf61.
//
// Solidity: function transferBatchByAdmin(address[] users, address[] to, uint256[] ids) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) TransferBatchByAdmin(opts *bind.TransactOpts, users []common.Address, to []common.Address, ids []*big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "transferBatchByAdmin", users, to, ids)
}

// TransferBatchByAdmin is a paid mutator transaction binding the contract method 0x159ccf61.
//
// Solidity: function transferBatchByAdmin(address[] users, address[] to, uint256[] ids) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) TransferBatchByAdmin(opts *bind.TransactOpts, users []common.Address, to []common.Address, ids []*big.Int) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "transferBatchByAdmin", users, to, ids)
}

// TransferBatchByAdmin is a paid mutator transaction binding the contract method 0x159ccf61.
//
// Solidity: function transferBatchByAdmin(address[] users, address[] to, uint256[] ids) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) TransferBatchByAdmin(users []common.Address, to []common.Address, ids []*big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.TransferBatchByAdmin(&_ERC721NFTCustom.TransactOpts, users, to, ids)
}

// TransferBatchByAdmin is a paid mutator transaction binding the contract method 0x159ccf61.
//
// Solidity: function transferBatchByAdmin(address[] users, address[] to, uint256[] ids) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) TransferBatchByAdmin(users []common.Address, to []common.Address, ids []*big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.TransferBatchByAdmin(&_ERC721NFTCustom.TransactOpts, users, to, ids)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0x8eb17dfe.
//
// Solidity: function transferByAdmin(address user, address to, uint256 id) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) TransferByAdmin(opts *bind.TransactOpts, user common.Address, to common.Address, id *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "transferByAdmin", user, to, id)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0x8eb17dfe.
//
// Solidity: function transferByAdmin(address user, address to, uint256 id) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) TransferByAdmin(opts *bind.TransactOpts, user common.Address, to common.Address, id *big.Int) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "transferByAdmin", user, to, id)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0x8eb17dfe.
//
// Solidity: function transferByAdmin(address user, address to, uint256 id) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) TransferByAdmin(user common.Address, to common.Address, id *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.TransferByAdmin(&_ERC721NFTCustom.TransactOpts, user, to, id)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0x8eb17dfe.
//
// Solidity: function transferByAdmin(address user, address to, uint256 id) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) TransferByAdmin(user common.Address, to common.Address, id *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.TransferByAdmin(&_ERC721NFTCustom.TransactOpts, user, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.TransferFrom(&_ERC721NFTCustom.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.TransferFrom(&_ERC721NFTCustom.TransactOpts, from, to, tokenId)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string newUri) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactor) UpdateTokenURI(opts *bind.TransactOpts, tokenId *big.Int, newUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.contract.Transact(opts, "updateTokenURI", tokenId, newUri)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string newUri) returns()
func (_ERC721NFTCustom *ERC721NFTCustomBulkTransactor) UpdateTokenURI(opts *bind.TransactOpts, tokenId *big.Int, newUri string) types.UnsignedTransaction {
	return _ERC721NFTCustom.contract.GenUnsignedTransaction(opts, "updateTokenURI", tokenId, newUri)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string newUri) returns()
func (_ERC721NFTCustom *ERC721NFTCustomSession) UpdateTokenURI(tokenId *big.Int, newUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.UpdateTokenURI(&_ERC721NFTCustom.TransactOpts, tokenId, newUri)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string newUri) returns()
func (_ERC721NFTCustom *ERC721NFTCustomTransactorSession) UpdateTokenURI(tokenId *big.Int, newUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC721NFTCustom.Contract.UpdateTokenURI(&_ERC721NFTCustom.TransactOpts, tokenId, newUri)
}

// ERC721NFTCustomApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721NFTCustom contract.
type ERC721NFTCustomApprovalIterator struct {
	Event *ERC721NFTCustomApproval // Event containing the contract specifics and raw log

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
func (it *ERC721NFTCustomApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721NFTCustomApproval)
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
		it.Event = new(ERC721NFTCustomApproval)
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
func (it *ERC721NFTCustomApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721NFTCustomApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721NFTCustomApproval represents a Approval event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// ERC721NFTCustomApprovalOrChainReorg represents a Approval subscription event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomApprovalOrChainReorg struct {
	Event      *ERC721NFTCustomApproval
	ChainReorg *types.ChainReorg
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721NFTCustomApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, err := _ERC721NFTCustom.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomApprovalIterator{contract: _ERC721NFTCustom.contract, event: "Approval", logs: logs}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721NFTCustomApprovalOrChainReorg, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721NFTCustom.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721NFTCustomApprovalOrChainReorg)
				event.Event = new(ERC721NFTCustomApproval)

				if log.ChainReorg == nil {
					if err := _ERC721NFTCustom.contract.UnpackLog(event.Event, "Approval", *log.Log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) ParseApproval(log types.Log) (*ERC721NFTCustomApproval, error) {
	event := new(ERC721NFTCustomApproval)
	if err := _ERC721NFTCustom.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721NFTCustomApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721NFTCustom contract.
type ERC721NFTCustomApprovalForAllIterator struct {
	Event *ERC721NFTCustomApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721NFTCustomApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721NFTCustomApprovalForAll)
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
		it.Event = new(ERC721NFTCustomApprovalForAll)
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
func (it *ERC721NFTCustomApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721NFTCustomApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721NFTCustomApprovalForAll represents a ApprovalForAll event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// ERC721NFTCustomApprovalForAllOrChainReorg represents a ApprovalForAll subscription event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomApprovalForAllOrChainReorg struct {
	Event      *ERC721NFTCustomApprovalForAll
	ChainReorg *types.ChainReorg
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721NFTCustomApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, err := _ERC721NFTCustom.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomApprovalForAllIterator{contract: _ERC721NFTCustom.contract, event: "ApprovalForAll", logs: logs}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721NFTCustomApprovalForAllOrChainReorg, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721NFTCustom.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721NFTCustomApprovalForAllOrChainReorg)
				event.Event = new(ERC721NFTCustomApprovalForAll)

				if log.ChainReorg == nil {
					if err := _ERC721NFTCustom.contract.UnpackLog(event.Event, "ApprovalForAll", *log.Log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) ParseApprovalForAll(log types.Log) (*ERC721NFTCustomApprovalForAll, error) {
	event := new(ERC721NFTCustomApprovalForAll)
	if err := _ERC721NFTCustom.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721NFTCustomInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC721NFTCustom contract.
type ERC721NFTCustomInitializedIterator struct {
	Event *ERC721NFTCustomInitialized // Event containing the contract specifics and raw log

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
func (it *ERC721NFTCustomInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721NFTCustomInitialized)
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
		it.Event = new(ERC721NFTCustomInitialized)
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
func (it *ERC721NFTCustomInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721NFTCustomInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721NFTCustomInitialized represents a Initialized event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// ERC721NFTCustomInitializedOrChainReorg represents a Initialized subscription event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomInitializedOrChainReorg struct {
	Event      *ERC721NFTCustomInitialized
	ChainReorg *types.ChainReorg
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC721NFTCustomInitializedIterator, error) {

	logs, err := _ERC721NFTCustom.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomInitializedIterator{contract: _ERC721NFTCustom.contract, event: "Initialized", logs: logs}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC721NFTCustomInitializedOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _ERC721NFTCustom.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721NFTCustomInitializedOrChainReorg)
				event.Event = new(ERC721NFTCustomInitialized)

				if log.ChainReorg == nil {
					if err := _ERC721NFTCustom.contract.UnpackLog(event.Event, "Initialized", *log.Log); err != nil {
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
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) ParseInitialized(log types.Log) (*ERC721NFTCustomInitialized, error) {
	event := new(ERC721NFTCustomInitialized)
	if err := _ERC721NFTCustom.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721NFTCustomPermanentURIIterator is returned from FilterPermanentURI and is used to iterate over the raw logs and unpacked data for PermanentURI events raised by the ERC721NFTCustom contract.
type ERC721NFTCustomPermanentURIIterator struct {
	Event *ERC721NFTCustomPermanentURI // Event containing the contract specifics and raw log

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
func (it *ERC721NFTCustomPermanentURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721NFTCustomPermanentURI)
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
		it.Event = new(ERC721NFTCustomPermanentURI)
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
func (it *ERC721NFTCustomPermanentURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721NFTCustomPermanentURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721NFTCustomPermanentURI represents a PermanentURI event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomPermanentURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// ERC721NFTCustomPermanentURIOrChainReorg represents a PermanentURI subscription event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomPermanentURIOrChainReorg struct {
	Event      *ERC721NFTCustomPermanentURI
	ChainReorg *types.ChainReorg
}

// FilterPermanentURI is a free log retrieval operation binding the contract event 0xa109ba539900bf1b633f956d63c96fc89b814c7287f7aa50a9216d0b55657207.
//
// Solidity: event PermanentURI(string value, uint256 indexed id)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) FilterPermanentURI(opts *bind.FilterOpts, id []*big.Int) (*ERC721NFTCustomPermanentURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, err := _ERC721NFTCustom.contract.FilterLogs(opts, "PermanentURI", idRule)
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomPermanentURIIterator{contract: _ERC721NFTCustom.contract, event: "PermanentURI", logs: logs}, nil
}

// WatchPermanentURI is a free log subscription operation binding the contract event 0xa109ba539900bf1b633f956d63c96fc89b814c7287f7aa50a9216d0b55657207.
//
// Solidity: event PermanentURI(string value, uint256 indexed id)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) WatchPermanentURI(opts *bind.WatchOpts, sink chan<- *ERC721NFTCustomPermanentURIOrChainReorg, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC721NFTCustom.contract.WatchLogs(opts, "PermanentURI", idRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721NFTCustomPermanentURIOrChainReorg)
				event.Event = new(ERC721NFTCustomPermanentURI)

				if log.ChainReorg == nil {
					if err := _ERC721NFTCustom.contract.UnpackLog(event.Event, "PermanentURI", *log.Log); err != nil {
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
// Solidity: event PermanentURI(string value, uint256 indexed id)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) ParsePermanentURI(log types.Log) (*ERC721NFTCustomPermanentURI, error) {
	event := new(ERC721NFTCustomPermanentURI)
	if err := _ERC721NFTCustom.contract.UnpackLog(event, "PermanentURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721NFTCustomPermanentURIGlobalIterator is returned from FilterPermanentURIGlobal and is used to iterate over the raw logs and unpacked data for PermanentURIGlobal events raised by the ERC721NFTCustom contract.
type ERC721NFTCustomPermanentURIGlobalIterator struct {
	Event *ERC721NFTCustomPermanentURIGlobal // Event containing the contract specifics and raw log

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
func (it *ERC721NFTCustomPermanentURIGlobalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721NFTCustomPermanentURIGlobal)
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
		it.Event = new(ERC721NFTCustomPermanentURIGlobal)
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
func (it *ERC721NFTCustomPermanentURIGlobalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721NFTCustomPermanentURIGlobalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721NFTCustomPermanentURIGlobal represents a PermanentURIGlobal event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomPermanentURIGlobal struct {
	Raw types.Log // Blockchain specific contextual infos
}

// ERC721NFTCustomPermanentURIGlobalOrChainReorg represents a PermanentURIGlobal subscription event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomPermanentURIGlobalOrChainReorg struct {
	Event      *ERC721NFTCustomPermanentURIGlobal
	ChainReorg *types.ChainReorg
}

// FilterPermanentURIGlobal is a free log retrieval operation binding the contract event 0xb59f45df38ec0d34114b1248c38a29cdbccbf3e745ae3ef310ac66199a4ceccf.
//
// Solidity: event PermanentURIGlobal()
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) FilterPermanentURIGlobal(opts *bind.FilterOpts) (*ERC721NFTCustomPermanentURIGlobalIterator, error) {

	logs, err := _ERC721NFTCustom.contract.FilterLogs(opts, "PermanentURIGlobal")
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomPermanentURIGlobalIterator{contract: _ERC721NFTCustom.contract, event: "PermanentURIGlobal", logs: logs}, nil
}

// WatchPermanentURIGlobal is a free log subscription operation binding the contract event 0xb59f45df38ec0d34114b1248c38a29cdbccbf3e745ae3ef310ac66199a4ceccf.
//
// Solidity: event PermanentURIGlobal()
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) WatchPermanentURIGlobal(opts *bind.WatchOpts, sink chan<- *ERC721NFTCustomPermanentURIGlobalOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _ERC721NFTCustom.contract.WatchLogs(opts, "PermanentURIGlobal")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721NFTCustomPermanentURIGlobalOrChainReorg)
				event.Event = new(ERC721NFTCustomPermanentURIGlobal)

				if log.ChainReorg == nil {
					if err := _ERC721NFTCustom.contract.UnpackLog(event.Event, "PermanentURIGlobal", *log.Log); err != nil {
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
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) ParsePermanentURIGlobal(log types.Log) (*ERC721NFTCustomPermanentURIGlobal, error) {
	event := new(ERC721NFTCustomPermanentURIGlobal)
	if err := _ERC721NFTCustom.contract.UnpackLog(event, "PermanentURIGlobal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721NFTCustomRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ERC721NFTCustom contract.
type ERC721NFTCustomRoleAdminChangedIterator struct {
	Event *ERC721NFTCustomRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC721NFTCustomRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721NFTCustomRoleAdminChanged)
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
		it.Event = new(ERC721NFTCustomRoleAdminChanged)
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
func (it *ERC721NFTCustomRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721NFTCustomRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721NFTCustomRoleAdminChanged represents a RoleAdminChanged event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// ERC721NFTCustomRoleAdminChangedOrChainReorg represents a RoleAdminChanged subscription event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomRoleAdminChangedOrChainReorg struct {
	Event      *ERC721NFTCustomRoleAdminChanged
	ChainReorg *types.ChainReorg
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ERC721NFTCustomRoleAdminChangedIterator, error) {

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

	logs, err := _ERC721NFTCustom.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomRoleAdminChangedIterator{contract: _ERC721NFTCustom.contract, event: "RoleAdminChanged", logs: logs}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC721NFTCustomRoleAdminChangedOrChainReorg, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ERC721NFTCustom.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721NFTCustomRoleAdminChangedOrChainReorg)
				event.Event = new(ERC721NFTCustomRoleAdminChanged)

				if log.ChainReorg == nil {
					if err := _ERC721NFTCustom.contract.UnpackLog(event.Event, "RoleAdminChanged", *log.Log); err != nil {
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
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) ParseRoleAdminChanged(log types.Log) (*ERC721NFTCustomRoleAdminChanged, error) {
	event := new(ERC721NFTCustomRoleAdminChanged)
	if err := _ERC721NFTCustom.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721NFTCustomRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ERC721NFTCustom contract.
type ERC721NFTCustomRoleGrantedIterator struct {
	Event *ERC721NFTCustomRoleGranted // Event containing the contract specifics and raw log

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
func (it *ERC721NFTCustomRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721NFTCustomRoleGranted)
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
		it.Event = new(ERC721NFTCustomRoleGranted)
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
func (it *ERC721NFTCustomRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721NFTCustomRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721NFTCustomRoleGranted represents a RoleGranted event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// ERC721NFTCustomRoleGrantedOrChainReorg represents a RoleGranted subscription event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomRoleGrantedOrChainReorg struct {
	Event      *ERC721NFTCustomRoleGranted
	ChainReorg *types.ChainReorg
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC721NFTCustomRoleGrantedIterator, error) {

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

	logs, err := _ERC721NFTCustom.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomRoleGrantedIterator{contract: _ERC721NFTCustom.contract, event: "RoleGranted", logs: logs}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ERC721NFTCustomRoleGrantedOrChainReorg, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC721NFTCustom.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721NFTCustomRoleGrantedOrChainReorg)
				event.Event = new(ERC721NFTCustomRoleGranted)

				if log.ChainReorg == nil {
					if err := _ERC721NFTCustom.contract.UnpackLog(event.Event, "RoleGranted", *log.Log); err != nil {
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
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) ParseRoleGranted(log types.Log) (*ERC721NFTCustomRoleGranted, error) {
	event := new(ERC721NFTCustomRoleGranted)
	if err := _ERC721NFTCustom.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721NFTCustomRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ERC721NFTCustom contract.
type ERC721NFTCustomRoleRevokedIterator struct {
	Event *ERC721NFTCustomRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ERC721NFTCustomRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721NFTCustomRoleRevoked)
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
		it.Event = new(ERC721NFTCustomRoleRevoked)
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
func (it *ERC721NFTCustomRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721NFTCustomRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721NFTCustomRoleRevoked represents a RoleRevoked event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// ERC721NFTCustomRoleRevokedOrChainReorg represents a RoleRevoked subscription event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomRoleRevokedOrChainReorg struct {
	Event      *ERC721NFTCustomRoleRevoked
	ChainReorg *types.ChainReorg
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC721NFTCustomRoleRevokedIterator, error) {

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

	logs, err := _ERC721NFTCustom.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomRoleRevokedIterator{contract: _ERC721NFTCustom.contract, event: "RoleRevoked", logs: logs}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ERC721NFTCustomRoleRevokedOrChainReorg, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC721NFTCustom.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721NFTCustomRoleRevokedOrChainReorg)
				event.Event = new(ERC721NFTCustomRoleRevoked)

				if log.ChainReorg == nil {
					if err := _ERC721NFTCustom.contract.UnpackLog(event.Event, "RoleRevoked", *log.Log); err != nil {
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
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) ParseRoleRevoked(log types.Log) (*ERC721NFTCustomRoleRevoked, error) {
	event := new(ERC721NFTCustomRoleRevoked)
	if err := _ERC721NFTCustom.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721NFTCustomRoyaltyUpdatedIterator is returned from FilterRoyaltyUpdated and is used to iterate over the raw logs and unpacked data for RoyaltyUpdated events raised by the ERC721NFTCustom contract.
type ERC721NFTCustomRoyaltyUpdatedIterator struct {
	Event *ERC721NFTCustomRoyaltyUpdated // Event containing the contract specifics and raw log

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
func (it *ERC721NFTCustomRoyaltyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721NFTCustomRoyaltyUpdated)
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
		it.Event = new(ERC721NFTCustomRoyaltyUpdated)
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
func (it *ERC721NFTCustomRoyaltyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721NFTCustomRoyaltyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721NFTCustomRoyaltyUpdated represents a RoyaltyUpdated event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomRoyaltyUpdated struct {
	RoyaltiesBps     *big.Int
	RoyaltiesAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// ERC721NFTCustomRoyaltyUpdatedOrChainReorg represents a RoyaltyUpdated subscription event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomRoyaltyUpdatedOrChainReorg struct {
	Event      *ERC721NFTCustomRoyaltyUpdated
	ChainReorg *types.ChainReorg
}

// FilterRoyaltyUpdated is a free log retrieval operation binding the contract event 0x13ceafb5e8ec39102bd452913a5c00a05f3b060d636d3a567c1c80e8b4321fe7.
//
// Solidity: event RoyaltyUpdated(uint256 royaltiesBps, address royaltiesAddress)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) FilterRoyaltyUpdated(opts *bind.FilterOpts) (*ERC721NFTCustomRoyaltyUpdatedIterator, error) {

	logs, err := _ERC721NFTCustom.contract.FilterLogs(opts, "RoyaltyUpdated")
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomRoyaltyUpdatedIterator{contract: _ERC721NFTCustom.contract, event: "RoyaltyUpdated", logs: logs}, nil
}

// WatchRoyaltyUpdated is a free log subscription operation binding the contract event 0x13ceafb5e8ec39102bd452913a5c00a05f3b060d636d3a567c1c80e8b4321fe7.
//
// Solidity: event RoyaltyUpdated(uint256 royaltiesBps, address royaltiesAddress)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) WatchRoyaltyUpdated(opts *bind.WatchOpts, sink chan<- *ERC721NFTCustomRoyaltyUpdatedOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _ERC721NFTCustom.contract.WatchLogs(opts, "RoyaltyUpdated")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721NFTCustomRoyaltyUpdatedOrChainReorg)
				event.Event = new(ERC721NFTCustomRoyaltyUpdated)

				if log.ChainReorg == nil {
					if err := _ERC721NFTCustom.contract.UnpackLog(event.Event, "RoyaltyUpdated", *log.Log); err != nil {
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
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) ParseRoyaltyUpdated(log types.Log) (*ERC721NFTCustomRoyaltyUpdated, error) {
	event := new(ERC721NFTCustomRoyaltyUpdated)
	if err := _ERC721NFTCustom.contract.UnpackLog(event, "RoyaltyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721NFTCustomTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721NFTCustom contract.
type ERC721NFTCustomTransferIterator struct {
	Event *ERC721NFTCustomTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721NFTCustomTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721NFTCustomTransfer)
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
		it.Event = new(ERC721NFTCustomTransfer)
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
func (it *ERC721NFTCustomTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721NFTCustomTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721NFTCustomTransfer represents a Transfer event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// ERC721NFTCustomTransferOrChainReorg represents a Transfer subscription event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomTransferOrChainReorg struct {
	Event      *ERC721NFTCustomTransfer
	ChainReorg *types.ChainReorg
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721NFTCustomTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, err := _ERC721NFTCustom.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomTransferIterator{contract: _ERC721NFTCustom.contract, event: "Transfer", logs: logs}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721NFTCustomTransferOrChainReorg, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721NFTCustom.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721NFTCustomTransferOrChainReorg)
				event.Event = new(ERC721NFTCustomTransfer)

				if log.ChainReorg == nil {
					if err := _ERC721NFTCustom.contract.UnpackLog(event.Event, "Transfer", *log.Log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) ParseTransfer(log types.Log) (*ERC721NFTCustomTransfer, error) {
	event := new(ERC721NFTCustomTransfer)
	if err := _ERC721NFTCustom.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721NFTCustomTransferCooldownTimeChangedIterator is returned from FilterTransferCooldownTimeChanged and is used to iterate over the raw logs and unpacked data for TransferCooldownTimeChanged events raised by the ERC721NFTCustom contract.
type ERC721NFTCustomTransferCooldownTimeChangedIterator struct {
	Event *ERC721NFTCustomTransferCooldownTimeChanged // Event containing the contract specifics and raw log

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
func (it *ERC721NFTCustomTransferCooldownTimeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721NFTCustomTransferCooldownTimeChanged)
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
		it.Event = new(ERC721NFTCustomTransferCooldownTimeChanged)
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
func (it *ERC721NFTCustomTransferCooldownTimeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721NFTCustomTransferCooldownTimeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721NFTCustomTransferCooldownTimeChanged represents a TransferCooldownTimeChanged event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomTransferCooldownTimeChanged struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// ERC721NFTCustomTransferCooldownTimeChangedOrChainReorg represents a TransferCooldownTimeChanged subscription event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomTransferCooldownTimeChangedOrChainReorg struct {
	Event      *ERC721NFTCustomTransferCooldownTimeChanged
	ChainReorg *types.ChainReorg
}

// FilterTransferCooldownTimeChanged is a free log retrieval operation binding the contract event 0x58279edb83db7d31273be58bda5721e9bfdf45b9ca02fc555c75eac8eae2f235.
//
// Solidity: event TransferCooldownTimeChanged(uint256 value)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) FilterTransferCooldownTimeChanged(opts *bind.FilterOpts) (*ERC721NFTCustomTransferCooldownTimeChangedIterator, error) {

	logs, err := _ERC721NFTCustom.contract.FilterLogs(opts, "TransferCooldownTimeChanged")
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomTransferCooldownTimeChangedIterator{contract: _ERC721NFTCustom.contract, event: "TransferCooldownTimeChanged", logs: logs}, nil
}

// WatchTransferCooldownTimeChanged is a free log subscription operation binding the contract event 0x58279edb83db7d31273be58bda5721e9bfdf45b9ca02fc555c75eac8eae2f235.
//
// Solidity: event TransferCooldownTimeChanged(uint256 value)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) WatchTransferCooldownTimeChanged(opts *bind.WatchOpts, sink chan<- *ERC721NFTCustomTransferCooldownTimeChangedOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _ERC721NFTCustom.contract.WatchLogs(opts, "TransferCooldownTimeChanged")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721NFTCustomTransferCooldownTimeChangedOrChainReorg)
				event.Event = new(ERC721NFTCustomTransferCooldownTimeChanged)

				if log.ChainReorg == nil {
					if err := _ERC721NFTCustom.contract.UnpackLog(event.Event, "TransferCooldownTimeChanged", *log.Log); err != nil {
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

// ParseTransferCooldownTimeChanged is a log parse operation binding the contract event 0x58279edb83db7d31273be58bda5721e9bfdf45b9ca02fc555c75eac8eae2f235.
//
// Solidity: event TransferCooldownTimeChanged(uint256 value)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) ParseTransferCooldownTimeChanged(log types.Log) (*ERC721NFTCustomTransferCooldownTimeChanged, error) {
	event := new(ERC721NFTCustomTransferCooldownTimeChanged)
	if err := _ERC721NFTCustom.contract.UnpackLog(event, "TransferCooldownTimeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721NFTCustomTransferableChangedIterator is returned from FilterTransferableChanged and is used to iterate over the raw logs and unpacked data for TransferableChanged events raised by the ERC721NFTCustom contract.
type ERC721NFTCustomTransferableChangedIterator struct {
	Event *ERC721NFTCustomTransferableChanged // Event containing the contract specifics and raw log

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
func (it *ERC721NFTCustomTransferableChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721NFTCustomTransferableChanged)
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
		it.Event = new(ERC721NFTCustomTransferableChanged)
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
func (it *ERC721NFTCustomTransferableChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721NFTCustomTransferableChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721NFTCustomTransferableChanged represents a TransferableChanged event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomTransferableChanged struct {
	TransferableByAdmin bool
	TransferableByUser  bool
	Raw                 types.Log // Blockchain specific contextual infos
}

// ERC721NFTCustomTransferableChangedOrChainReorg represents a TransferableChanged subscription event raised by the ERC721NFTCustom contract.
type ERC721NFTCustomTransferableChangedOrChainReorg struct {
	Event      *ERC721NFTCustomTransferableChanged
	ChainReorg *types.ChainReorg
}

// FilterTransferableChanged is a free log retrieval operation binding the contract event 0xab67e71edde643937fe4eedd295209d1ea4a844ff0f9203a60e651dd3e46687f.
//
// Solidity: event TransferableChanged(bool transferableByAdmin, bool transferableByUser)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) FilterTransferableChanged(opts *bind.FilterOpts) (*ERC721NFTCustomTransferableChangedIterator, error) {

	logs, err := _ERC721NFTCustom.contract.FilterLogs(opts, "TransferableChanged")
	if err != nil {
		return nil, err
	}
	return &ERC721NFTCustomTransferableChangedIterator{contract: _ERC721NFTCustom.contract, event: "TransferableChanged", logs: logs}, nil
}

// WatchTransferableChanged is a free log subscription operation binding the contract event 0xab67e71edde643937fe4eedd295209d1ea4a844ff0f9203a60e651dd3e46687f.
//
// Solidity: event TransferableChanged(bool transferableByAdmin, bool transferableByUser)
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) WatchTransferableChanged(opts *bind.WatchOpts, sink chan<- *ERC721NFTCustomTransferableChangedOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _ERC721NFTCustom.contract.WatchLogs(opts, "TransferableChanged")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721NFTCustomTransferableChangedOrChainReorg)
				event.Event = new(ERC721NFTCustomTransferableChanged)

				if log.ChainReorg == nil {
					if err := _ERC721NFTCustom.contract.UnpackLog(event.Event, "TransferableChanged", *log.Log); err != nil {
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
func (_ERC721NFTCustom *ERC721NFTCustomFilterer) ParseTransferableChanged(log types.Log) (*ERC721NFTCustomTransferableChanged, error) {
	event := new(ERC721NFTCustomTransferableChanged)
	if err := _ERC721NFTCustom.contract.UnpackLog(event, "TransferableChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
