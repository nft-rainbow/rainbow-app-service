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

// ERC1155NFTABI is the input ABI used to generate the binding from.
const ERC1155NFTABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"}],\"name\":\"mintTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"}],\"name\":\"mintToBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newURI\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC1155NFT is an auto generated Go binding around an Conflux contract.
type ERC1155NFT struct {
	ERC1155NFTCaller     // Read-only binding to the contract
	ERC1155NFTTransactor // Write-only binding to the contract
	ERC1155NFTFilterer   // Log filterer for contract events
}

// ERC1155NFTCaller is an auto generated read-only Go binding around an Conflux contract.
type ERC1155NFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155NFTBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type ERC1155NFTBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155NFTTransactor is an auto generated write-only Go binding around an Conflux contract.
type ERC1155NFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155NFTBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type ERC1155NFTBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155NFTFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type ERC1155NFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155NFTSession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type ERC1155NFTSession struct {
	Contract     *ERC1155NFT       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1155NFTCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type ERC1155NFTCallerSession struct {
	Contract *ERC1155NFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC1155NFTTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type ERC1155NFTTransactorSession struct {
	Contract     *ERC1155NFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC1155NFTRaw is an auto generated low-level Go binding around an Conflux contract.
type ERC1155NFTRaw struct {
	Contract *ERC1155NFT // Generic contract binding to access the raw methods on
}

// ERC1155NFTCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type ERC1155NFTCallerRaw struct {
	Contract *ERC1155NFTCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1155NFTTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type ERC1155NFTTransactorRaw struct {
	Contract *ERC1155NFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155NFT creates a new instance of ERC1155NFT, bound to a specific deployed contract.
func NewERC1155NFT(address types.Address, backend bind.ContractBackend) (*ERC1155NFT, error) {
	contract, err := bindERC1155NFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFT{ERC1155NFTCaller: ERC1155NFTCaller{contract: contract}, ERC1155NFTTransactor: ERC1155NFTTransactor{contract: contract}, ERC1155NFTFilterer: ERC1155NFTFilterer{contract: contract}}, nil
}

// NewERC1155NFTCaller creates a new read-only instance of ERC1155NFT, bound to a specific deployed contract.
func NewERC1155NFTCaller(address types.Address, caller bind.ContractCaller) (*ERC1155NFTCaller, error) {
	contract, err := bindERC1155NFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTCaller{contract: contract}, nil
}

// NewERC1155NFTTransactor creates a new write-only instance of ERC1155NFT, bound to a specific deployed contract.
func NewERC1155NFTTransactor(address types.Address, transactor bind.ContractTransactor) (*ERC1155NFTTransactor, error) {
	contract, err := bindERC1155NFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTTransactor{contract: contract}, nil
}

// NewERC1155NFTFilterer creates a new log filterer instance of ERC1155NFT, bound to a specific deployed contract.
func NewERC1155NFTFilterer(address types.Address, filterer bind.ContractFilterer) (*ERC1155NFTFilterer, error) {
	contract, err := bindERC1155NFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTFilterer{contract: contract}, nil
}

// NewERC1155NFTCaller creates a new read-only instance of ERC1155NFT, bound to a specific deployed contract.
func NewERC1155NFTBulkCaller(address types.Address, caller bind.ContractCaller) (*ERC1155NFTBulkCaller, error) {
	contract, err := bindERC1155NFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTBulkCaller{contract: contract}, nil
}

// NewERC1155NFTBulkTransactor creates a new write-only instance of ERC1155NFT, bound to a specific deployed contract.
func NewERC1155NFTBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*ERC1155NFTBulkTransactor, error) {
	contract, err := bindERC1155NFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTBulkTransactor{contract: contract}, nil
}

// bindERC1155NFT binds a generic wrapper to an already deployed contract.
func bindERC1155NFT(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1155NFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155NFT *ERC1155NFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155NFT.Contract.ERC1155NFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155NFT *ERC1155NFTRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.ERC1155NFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155NFT *ERC1155NFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.ERC1155NFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155NFT *ERC1155NFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155NFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155NFT *ERC1155NFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155NFT *ERC1155NFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC1155NFT *ERC1155NFTCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	__err := _ERC1155NFT.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC1155NFT *ERC1155NFTBulkCaller) DEFAULTADMINROLE(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFT.contract.GenRequest(opts, "DEFAULT_ADMIN_ROLE")

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFT.contract.DecodeOutput(&out, rawOut, "DEFAULT_ADMIN_ROLE")
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
func (_ERC1155NFT *ERC1155NFTSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC1155NFT.Contract.DEFAULTADMINROLE(&_ERC1155NFT.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC1155NFT *ERC1155NFTCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC1155NFT.Contract.DEFAULTADMINROLE(&_ERC1155NFT.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC1155NFT *ERC1155NFTCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	__err := _ERC1155NFT.contract.Call(opts, &out, "MINTER_ROLE")

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC1155NFT *ERC1155NFTBulkCaller) MINTERROLE(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFT.contract.GenRequest(opts, "MINTER_ROLE")

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFT.contract.DecodeOutput(&out, rawOut, "MINTER_ROLE")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC1155NFT *ERC1155NFTSession) MINTERROLE() ([32]byte, error) {
	return _ERC1155NFT.Contract.MINTERROLE(&_ERC1155NFT.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC1155NFT *ERC1155NFTCallerSession) MINTERROLE() ([32]byte, error) {
	return _ERC1155NFT.Contract.MINTERROLE(&_ERC1155NFT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155NFT *ERC1155NFTCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	__err := _ERC1155NFT.contract.Call(opts, &out, "balanceOf", account, id)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155NFT *ERC1155NFTBulkCaller) BalanceOf(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, account common.Address, id *big.Int) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFT.contract.GenRequest(opts, "balanceOf", account, id)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFT.contract.DecodeOutput(&out, rawOut, "balanceOf")
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
func (_ERC1155NFT *ERC1155NFTSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155NFT.Contract.BalanceOf(&_ERC1155NFT.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155NFT *ERC1155NFTCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155NFT.Contract.BalanceOf(&_ERC1155NFT.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155NFT *ERC1155NFTCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	__err := _ERC1155NFT.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if __err != nil {
		return *new([]*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, __err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155NFT *ERC1155NFTBulkCaller) BalanceOfBatch(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) (*[]*big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFT.contract.GenRequest(opts, "balanceOfBatch", accounts, ids)

	out0 := new([]*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFT.contract.DecodeOutput(&out, rawOut, "balanceOfBatch")
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
func (_ERC1155NFT *ERC1155NFTSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ERC1155NFT.Contract.BalanceOfBatch(&_ERC1155NFT.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155NFT *ERC1155NFTCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ERC1155NFT.Contract.BalanceOfBatch(&_ERC1155NFT.CallOpts, accounts, ids)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC1155NFT *ERC1155NFTCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	__err := _ERC1155NFT.contract.Call(opts, &out, "getRoleAdmin", role)

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC1155NFT *ERC1155NFTBulkCaller) GetRoleAdmin(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, role [32]byte) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFT.contract.GenRequest(opts, "getRoleAdmin", role)

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFT.contract.DecodeOutput(&out, rawOut, "getRoleAdmin")
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
func (_ERC1155NFT *ERC1155NFTSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC1155NFT.Contract.GetRoleAdmin(&_ERC1155NFT.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC1155NFT *ERC1155NFTCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC1155NFT.Contract.GetRoleAdmin(&_ERC1155NFT.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC1155NFT *ERC1155NFTCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	__err := _ERC1155NFT.contract.Call(opts, &out, "hasRole", role, account)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC1155NFT *ERC1155NFTBulkCaller) HasRole(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, role [32]byte, account common.Address) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFT.contract.GenRequest(opts, "hasRole", role, account)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFT.contract.DecodeOutput(&out, rawOut, "hasRole")
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
func (_ERC1155NFT *ERC1155NFTSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC1155NFT.Contract.HasRole(&_ERC1155NFT.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC1155NFT *ERC1155NFTCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC1155NFT.Contract.HasRole(&_ERC1155NFT.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155NFT *ERC1155NFTCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	__err := _ERC1155NFT.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155NFT *ERC1155NFTBulkCaller) IsApprovedForAll(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, account common.Address, operator common.Address) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFT.contract.GenRequest(opts, "isApprovedForAll", account, operator)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFT.contract.DecodeOutput(&out, rawOut, "isApprovedForAll")
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
func (_ERC1155NFT *ERC1155NFTSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ERC1155NFT.Contract.IsApprovedForAll(&_ERC1155NFT.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155NFT *ERC1155NFTCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ERC1155NFT.Contract.IsApprovedForAll(&_ERC1155NFT.CallOpts, account, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC1155NFT *ERC1155NFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _ERC1155NFT.contract.Call(opts, &out, "name")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC1155NFT *ERC1155NFTBulkCaller) Name(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFT.contract.GenRequest(opts, "name")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFT.contract.DecodeOutput(&out, rawOut, "name")
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
func (_ERC1155NFT *ERC1155NFTSession) Name() (string, error) {
	return _ERC1155NFT.Contract.Name(&_ERC1155NFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC1155NFT *ERC1155NFTCallerSession) Name() (string, error) {
	return _ERC1155NFT.Contract.Name(&_ERC1155NFT.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155NFT *ERC1155NFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	__err := _ERC1155NFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155NFT *ERC1155NFTBulkCaller) SupportsInterface(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, interfaceId [4]byte) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFT.contract.GenRequest(opts, "supportsInterface", interfaceId)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFT.contract.DecodeOutput(&out, rawOut, "supportsInterface")
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
func (_ERC1155NFT *ERC1155NFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155NFT.Contract.SupportsInterface(&_ERC1155NFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155NFT *ERC1155NFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155NFT.Contract.SupportsInterface(&_ERC1155NFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC1155NFT *ERC1155NFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _ERC1155NFT.contract.Call(opts, &out, "symbol")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC1155NFT *ERC1155NFTBulkCaller) Symbol(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFT.contract.GenRequest(opts, "symbol")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFT.contract.DecodeOutput(&out, rawOut, "symbol")
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
func (_ERC1155NFT *ERC1155NFTSession) Symbol() (string, error) {
	return _ERC1155NFT.Contract.Symbol(&_ERC1155NFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC1155NFT *ERC1155NFTCallerSession) Symbol() (string, error) {
	return _ERC1155NFT.Contract.Symbol(&_ERC1155NFT.CallOpts)
}

// TokenSupply is a free data retrieval call binding the contract method 0x2693ebf2.
//
// Solidity: function tokenSupply(uint256 ) view returns(uint256)
func (_ERC1155NFT *ERC1155NFTCaller) TokenSupply(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	__err := _ERC1155NFT.contract.Call(opts, &out, "tokenSupply", arg0)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// TokenSupply is a free data retrieval call binding the contract method 0x2693ebf2.
//
// Solidity: function tokenSupply(uint256 ) view returns(uint256)
func (_ERC1155NFT *ERC1155NFTBulkCaller) TokenSupply(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, arg0 *big.Int) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFT.contract.GenRequest(opts, "tokenSupply", arg0)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFT.contract.DecodeOutput(&out, rawOut, "tokenSupply")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TokenSupply is a free data retrieval call binding the contract method 0x2693ebf2.
//
// Solidity: function tokenSupply(uint256 ) view returns(uint256)
func (_ERC1155NFT *ERC1155NFTSession) TokenSupply(arg0 *big.Int) (*big.Int, error) {
	return _ERC1155NFT.Contract.TokenSupply(&_ERC1155NFT.CallOpts, arg0)
}

// TokenSupply is a free data retrieval call binding the contract method 0x2693ebf2.
//
// Solidity: function tokenSupply(uint256 ) view returns(uint256)
func (_ERC1155NFT *ERC1155NFTCallerSession) TokenSupply(arg0 *big.Int) (*big.Int, error) {
	return _ERC1155NFT.Contract.TokenSupply(&_ERC1155NFT.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_ERC1155NFT *ERC1155NFTCaller) TotalSupply(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	__err := _ERC1155NFT.contract.Call(opts, &out, "totalSupply", _id)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_ERC1155NFT *ERC1155NFTBulkCaller) TotalSupply(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, _id *big.Int) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFT.contract.GenRequest(opts, "totalSupply", _id)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFT.contract.DecodeOutput(&out, rawOut, "totalSupply")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_ERC1155NFT *ERC1155NFTSession) TotalSupply(_id *big.Int) (*big.Int, error) {
	return _ERC1155NFT.Contract.TotalSupply(&_ERC1155NFT.CallOpts, _id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 _id) view returns(uint256)
func (_ERC1155NFT *ERC1155NFTCallerSession) TotalSupply(_id *big.Int) (*big.Int, error) {
	return _ERC1155NFT.Contract.TotalSupply(&_ERC1155NFT.CallOpts, _id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_ERC1155NFT *ERC1155NFTCaller) Uri(opts *bind.CallOpts, _id *big.Int) (string, error) {
	var out []interface{}
	__err := _ERC1155NFT.contract.Call(opts, &out, "uri", _id)

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_ERC1155NFT *ERC1155NFTBulkCaller) Uri(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, _id *big.Int) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _ERC1155NFT.contract.GenRequest(opts, "uri", _id)

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _ERC1155NFT.contract.DecodeOutput(&out, rawOut, "uri")
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
// Solidity: function uri(uint256 _id) view returns(string)
func (_ERC1155NFT *ERC1155NFTSession) Uri(_id *big.Int) (string, error) {
	return _ERC1155NFT.Contract.Uri(&_ERC1155NFT.CallOpts, _id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_ERC1155NFT *ERC1155NFTCallerSession) Uri(_id *big.Int) (string, error) {
	return _ERC1155NFT.Contract.Uri(&_ERC1155NFT.CallOpts, _id)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC1155NFT *ERC1155NFTTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC1155NFT *ERC1155NFTBulkTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) types.UnsignedTransaction {
	return _ERC1155NFT.contract.GenUnsignedTransaction(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC1155NFT *ERC1155NFTSession) GrantRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.GrantRole(&_ERC1155NFT.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC1155NFT *ERC1155NFTTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.GrantRole(&_ERC1155NFT.TransactOpts, role, account)
}

// MintTo is a paid mutator transaction binding the contract method 0x3dbd5b25.
//
// Solidity: function mintTo(address account, uint256 id, uint256 amount, string tokenUri) returns()
func (_ERC1155NFT *ERC1155NFTTransactor) MintTo(opts *bind.TransactOpts, account common.Address, id *big.Int, amount *big.Int, tokenUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.contract.Transact(opts, "mintTo", account, id, amount, tokenUri)
}

// MintTo is a paid mutator transaction binding the contract method 0x3dbd5b25.
//
// Solidity: function mintTo(address account, uint256 id, uint256 amount, string tokenUri) returns()
func (_ERC1155NFT *ERC1155NFTBulkTransactor) MintTo(opts *bind.TransactOpts, account common.Address, id *big.Int, amount *big.Int, tokenUri string) types.UnsignedTransaction {
	return _ERC1155NFT.contract.GenUnsignedTransaction(opts, "mintTo", account, id, amount, tokenUri)
}

// MintTo is a paid mutator transaction binding the contract method 0x3dbd5b25.
//
// Solidity: function mintTo(address account, uint256 id, uint256 amount, string tokenUri) returns()
func (_ERC1155NFT *ERC1155NFTSession) MintTo(account common.Address, id *big.Int, amount *big.Int, tokenUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.MintTo(&_ERC1155NFT.TransactOpts, account, id, amount, tokenUri)
}

// MintTo is a paid mutator transaction binding the contract method 0x3dbd5b25.
//
// Solidity: function mintTo(address account, uint256 id, uint256 amount, string tokenUri) returns()
func (_ERC1155NFT *ERC1155NFTTransactorSession) MintTo(account common.Address, id *big.Int, amount *big.Int, tokenUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.MintTo(&_ERC1155NFT.TransactOpts, account, id, amount, tokenUri)
}

// MintTo0 is a paid mutator transaction binding the contract method 0x9f6ed25f.
//
// Solidity: function mintTo(address account, uint256 id, string tokenUri) returns(uint256)
func (_ERC1155NFT *ERC1155NFTTransactor) MintTo0(opts *bind.TransactOpts, account common.Address, id *big.Int, tokenUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.contract.Transact(opts, "mintTo0", account, id, tokenUri)
}

// MintTo0 is a paid mutator transaction binding the contract method 0x9f6ed25f.
//
// Solidity: function mintTo(address account, uint256 id, string tokenUri) returns(uint256)
func (_ERC1155NFT *ERC1155NFTBulkTransactor) MintTo0(opts *bind.TransactOpts, account common.Address, id *big.Int, tokenUri string) types.UnsignedTransaction {
	return _ERC1155NFT.contract.GenUnsignedTransaction(opts, "mintTo0", account, id, tokenUri)
}

// MintTo0 is a paid mutator transaction binding the contract method 0x9f6ed25f.
//
// Solidity: function mintTo(address account, uint256 id, string tokenUri) returns(uint256)
func (_ERC1155NFT *ERC1155NFTSession) MintTo0(account common.Address, id *big.Int, tokenUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.MintTo0(&_ERC1155NFT.TransactOpts, account, id, tokenUri)
}

// MintTo0 is a paid mutator transaction binding the contract method 0x9f6ed25f.
//
// Solidity: function mintTo(address account, uint256 id, string tokenUri) returns(uint256)
func (_ERC1155NFT *ERC1155NFTTransactorSession) MintTo0(account common.Address, id *big.Int, tokenUri string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.MintTo0(&_ERC1155NFT.TransactOpts, account, id, tokenUri)
}

// MintToBatch is a paid mutator transaction binding the contract method 0x00c538b2.
//
// Solidity: function mintToBatch(address to, uint256[] ids, uint256[] amounts, string[] uris) returns()
func (_ERC1155NFT *ERC1155NFTTransactor) MintToBatch(opts *bind.TransactOpts, to common.Address, ids []*big.Int, amounts []*big.Int, uris []string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.contract.Transact(opts, "mintToBatch", to, ids, amounts, uris)
}

// MintToBatch is a paid mutator transaction binding the contract method 0x00c538b2.
//
// Solidity: function mintToBatch(address to, uint256[] ids, uint256[] amounts, string[] uris) returns()
func (_ERC1155NFT *ERC1155NFTBulkTransactor) MintToBatch(opts *bind.TransactOpts, to common.Address, ids []*big.Int, amounts []*big.Int, uris []string) types.UnsignedTransaction {
	return _ERC1155NFT.contract.GenUnsignedTransaction(opts, "mintToBatch", to, ids, amounts, uris)
}

// MintToBatch is a paid mutator transaction binding the contract method 0x00c538b2.
//
// Solidity: function mintToBatch(address to, uint256[] ids, uint256[] amounts, string[] uris) returns()
func (_ERC1155NFT *ERC1155NFTSession) MintToBatch(to common.Address, ids []*big.Int, amounts []*big.Int, uris []string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.MintToBatch(&_ERC1155NFT.TransactOpts, to, ids, amounts, uris)
}

// MintToBatch is a paid mutator transaction binding the contract method 0x00c538b2.
//
// Solidity: function mintToBatch(address to, uint256[] ids, uint256[] amounts, string[] uris) returns()
func (_ERC1155NFT *ERC1155NFTTransactorSession) MintToBatch(to common.Address, ids []*big.Int, amounts []*big.Int, uris []string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.MintToBatch(&_ERC1155NFT.TransactOpts, to, ids, amounts, uris)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC1155NFT *ERC1155NFTTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC1155NFT *ERC1155NFTBulkTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) types.UnsignedTransaction {
	return _ERC1155NFT.contract.GenUnsignedTransaction(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC1155NFT *ERC1155NFTSession) RenounceRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.RenounceRole(&_ERC1155NFT.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ERC1155NFT *ERC1155NFTTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.RenounceRole(&_ERC1155NFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC1155NFT *ERC1155NFTTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC1155NFT *ERC1155NFTBulkTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) types.UnsignedTransaction {
	return _ERC1155NFT.contract.GenUnsignedTransaction(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC1155NFT *ERC1155NFTSession) RevokeRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.RevokeRole(&_ERC1155NFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC1155NFT *ERC1155NFTTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.RevokeRole(&_ERC1155NFT.TransactOpts, role, account)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155NFT *ERC1155NFTTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155NFT *ERC1155NFTBulkTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) types.UnsignedTransaction {
	return _ERC1155NFT.contract.GenUnsignedTransaction(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155NFT *ERC1155NFTSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.SafeBatchTransferFrom(&_ERC1155NFT.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155NFT *ERC1155NFTTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.SafeBatchTransferFrom(&_ERC1155NFT.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155NFT *ERC1155NFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155NFT *ERC1155NFTBulkTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) types.UnsignedTransaction {
	return _ERC1155NFT.contract.GenUnsignedTransaction(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155NFT *ERC1155NFTSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.SafeTransferFrom(&_ERC1155NFT.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155NFT *ERC1155NFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.SafeTransferFrom(&_ERC1155NFT.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155NFT *ERC1155NFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155NFT *ERC1155NFTBulkTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) types.UnsignedTransaction {
	return _ERC1155NFT.contract.GenUnsignedTransaction(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155NFT *ERC1155NFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.SetApprovalForAll(&_ERC1155NFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155NFT *ERC1155NFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.SetApprovalForAll(&_ERC1155NFT.TransactOpts, operator, approved)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string _newURI) returns()
func (_ERC1155NFT *ERC1155NFTTransactor) SetURI(opts *bind.TransactOpts, _newURI string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.contract.Transact(opts, "setURI", _newURI)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string _newURI) returns()
func (_ERC1155NFT *ERC1155NFTBulkTransactor) SetURI(opts *bind.TransactOpts, _newURI string) types.UnsignedTransaction {
	return _ERC1155NFT.contract.GenUnsignedTransaction(opts, "setURI", _newURI)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string _newURI) returns()
func (_ERC1155NFT *ERC1155NFTSession) SetURI(_newURI string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.SetURI(&_ERC1155NFT.TransactOpts, _newURI)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string _newURI) returns()
func (_ERC1155NFT *ERC1155NFTTransactorSession) SetURI(_newURI string) (*types.UnsignedTransaction, *types.Hash, error) {
	return _ERC1155NFT.Contract.SetURI(&_ERC1155NFT.TransactOpts, _newURI)
}

// ERC1155NFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC1155NFT contract.
type ERC1155NFTApprovalForAllIterator struct {
	Event *ERC1155NFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC1155NFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTApprovalForAll)
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
		it.Event = new(ERC1155NFTApprovalForAll)
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
func (it *ERC1155NFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTApprovalForAll represents a ApprovalForAll event raised by the ERC1155NFT contract.
type ERC1155NFTApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// ERC1155NFTApprovalForAllOrChainReorg represents a ApprovalForAll subscription event raised by the ERC1155NFT contract.
type ERC1155NFTApprovalForAllOrChainReorg struct {
	Event      *ERC1155NFTApprovalForAll
	ChainReorg *types.ChainReorg
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155NFT *ERC1155NFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ERC1155NFTApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, err := _ERC1155NFT.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTApprovalForAllIterator{contract: _ERC1155NFT.contract, event: "ApprovalForAll", logs: logs}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155NFT *ERC1155NFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC1155NFTApprovalForAllOrChainReorg, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC1155NFT.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTApprovalForAllOrChainReorg)
				event.Event = new(ERC1155NFTApprovalForAll)

				if log.ChainReorg == nil {
					if err := _ERC1155NFT.contract.UnpackLog(event.Event, "ApprovalForAll", *log.Log); err != nil {
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
func (_ERC1155NFT *ERC1155NFTFilterer) ParseApprovalForAll(log types.Log) (*ERC1155NFTApprovalForAll, error) {
	event := new(ERC1155NFTApprovalForAll)
	if err := _ERC1155NFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ERC1155NFT contract.
type ERC1155NFTRoleAdminChangedIterator struct {
	Event *ERC1155NFTRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC1155NFTRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTRoleAdminChanged)
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
		it.Event = new(ERC1155NFTRoleAdminChanged)
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
func (it *ERC1155NFTRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTRoleAdminChanged represents a RoleAdminChanged event raised by the ERC1155NFT contract.
type ERC1155NFTRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// ERC1155NFTRoleAdminChangedOrChainReorg represents a RoleAdminChanged subscription event raised by the ERC1155NFT contract.
type ERC1155NFTRoleAdminChangedOrChainReorg struct {
	Event      *ERC1155NFTRoleAdminChanged
	ChainReorg *types.ChainReorg
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC1155NFT *ERC1155NFTFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ERC1155NFTRoleAdminChangedIterator, error) {

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

	logs, err := _ERC1155NFT.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTRoleAdminChangedIterator{contract: _ERC1155NFT.contract, event: "RoleAdminChanged", logs: logs}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC1155NFT *ERC1155NFTFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC1155NFTRoleAdminChangedOrChainReorg, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ERC1155NFT.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTRoleAdminChangedOrChainReorg)
				event.Event = new(ERC1155NFTRoleAdminChanged)

				if log.ChainReorg == nil {
					if err := _ERC1155NFT.contract.UnpackLog(event.Event, "RoleAdminChanged", *log.Log); err != nil {
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
func (_ERC1155NFT *ERC1155NFTFilterer) ParseRoleAdminChanged(log types.Log) (*ERC1155NFTRoleAdminChanged, error) {
	event := new(ERC1155NFTRoleAdminChanged)
	if err := _ERC1155NFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ERC1155NFT contract.
type ERC1155NFTRoleGrantedIterator struct {
	Event *ERC1155NFTRoleGranted // Event containing the contract specifics and raw log

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
func (it *ERC1155NFTRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTRoleGranted)
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
		it.Event = new(ERC1155NFTRoleGranted)
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
func (it *ERC1155NFTRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTRoleGranted represents a RoleGranted event raised by the ERC1155NFT contract.
type ERC1155NFTRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// ERC1155NFTRoleGrantedOrChainReorg represents a RoleGranted subscription event raised by the ERC1155NFT contract.
type ERC1155NFTRoleGrantedOrChainReorg struct {
	Event      *ERC1155NFTRoleGranted
	ChainReorg *types.ChainReorg
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC1155NFT *ERC1155NFTFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC1155NFTRoleGrantedIterator, error) {

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

	logs, err := _ERC1155NFT.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTRoleGrantedIterator{contract: _ERC1155NFT.contract, event: "RoleGranted", logs: logs}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC1155NFT *ERC1155NFTFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ERC1155NFTRoleGrantedOrChainReorg, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC1155NFT.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTRoleGrantedOrChainReorg)
				event.Event = new(ERC1155NFTRoleGranted)

				if log.ChainReorg == nil {
					if err := _ERC1155NFT.contract.UnpackLog(event.Event, "RoleGranted", *log.Log); err != nil {
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
func (_ERC1155NFT *ERC1155NFTFilterer) ParseRoleGranted(log types.Log) (*ERC1155NFTRoleGranted, error) {
	event := new(ERC1155NFTRoleGranted)
	if err := _ERC1155NFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ERC1155NFT contract.
type ERC1155NFTRoleRevokedIterator struct {
	Event *ERC1155NFTRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ERC1155NFTRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTRoleRevoked)
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
		it.Event = new(ERC1155NFTRoleRevoked)
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
func (it *ERC1155NFTRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTRoleRevoked represents a RoleRevoked event raised by the ERC1155NFT contract.
type ERC1155NFTRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// ERC1155NFTRoleRevokedOrChainReorg represents a RoleRevoked subscription event raised by the ERC1155NFT contract.
type ERC1155NFTRoleRevokedOrChainReorg struct {
	Event      *ERC1155NFTRoleRevoked
	ChainReorg *types.ChainReorg
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC1155NFT *ERC1155NFTFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC1155NFTRoleRevokedIterator, error) {

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

	logs, err := _ERC1155NFT.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTRoleRevokedIterator{contract: _ERC1155NFT.contract, event: "RoleRevoked", logs: logs}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC1155NFT *ERC1155NFTFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ERC1155NFTRoleRevokedOrChainReorg, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC1155NFT.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTRoleRevokedOrChainReorg)
				event.Event = new(ERC1155NFTRoleRevoked)

				if log.ChainReorg == nil {
					if err := _ERC1155NFT.contract.UnpackLog(event.Event, "RoleRevoked", *log.Log); err != nil {
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
func (_ERC1155NFT *ERC1155NFTFilterer) ParseRoleRevoked(log types.Log) (*ERC1155NFTRoleRevoked, error) {
	event := new(ERC1155NFTRoleRevoked)
	if err := _ERC1155NFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ERC1155NFT contract.
type ERC1155NFTTransferBatchIterator struct {
	Event *ERC1155NFTTransferBatch // Event containing the contract specifics and raw log

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
func (it *ERC1155NFTTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTTransferBatch)
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
		it.Event = new(ERC1155NFTTransferBatch)
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
func (it *ERC1155NFTTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTTransferBatch represents a TransferBatch event raised by the ERC1155NFT contract.
type ERC1155NFTTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// ERC1155NFTTransferBatchOrChainReorg represents a TransferBatch subscription event raised by the ERC1155NFT contract.
type ERC1155NFTTransferBatchOrChainReorg struct {
	Event      *ERC1155NFTTransferBatch
	ChainReorg *types.ChainReorg
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155NFT *ERC1155NFTFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ERC1155NFTTransferBatchIterator, error) {

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

	logs, err := _ERC1155NFT.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTTransferBatchIterator{contract: _ERC1155NFT.contract, event: "TransferBatch", logs: logs}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155NFT *ERC1155NFTFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ERC1155NFTTransferBatchOrChainReorg, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC1155NFT.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTTransferBatchOrChainReorg)
				event.Event = new(ERC1155NFTTransferBatch)

				if log.ChainReorg == nil {
					if err := _ERC1155NFT.contract.UnpackLog(event.Event, "TransferBatch", *log.Log); err != nil {
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
func (_ERC1155NFT *ERC1155NFTFilterer) ParseTransferBatch(log types.Log) (*ERC1155NFTTransferBatch, error) {
	event := new(ERC1155NFTTransferBatch)
	if err := _ERC1155NFT.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ERC1155NFT contract.
type ERC1155NFTTransferSingleIterator struct {
	Event *ERC1155NFTTransferSingle // Event containing the contract specifics and raw log

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
func (it *ERC1155NFTTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTTransferSingle)
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
		it.Event = new(ERC1155NFTTransferSingle)
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
func (it *ERC1155NFTTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTTransferSingle represents a TransferSingle event raised by the ERC1155NFT contract.
type ERC1155NFTTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// ERC1155NFTTransferSingleOrChainReorg represents a TransferSingle subscription event raised by the ERC1155NFT contract.
type ERC1155NFTTransferSingleOrChainReorg struct {
	Event      *ERC1155NFTTransferSingle
	ChainReorg *types.ChainReorg
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155NFT *ERC1155NFTFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ERC1155NFTTransferSingleIterator, error) {

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

	logs, err := _ERC1155NFT.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTTransferSingleIterator{contract: _ERC1155NFT.contract, event: "TransferSingle", logs: logs}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155NFT *ERC1155NFTFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ERC1155NFTTransferSingleOrChainReorg, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC1155NFT.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTTransferSingleOrChainReorg)
				event.Event = new(ERC1155NFTTransferSingle)

				if log.ChainReorg == nil {
					if err := _ERC1155NFT.contract.UnpackLog(event.Event, "TransferSingle", *log.Log); err != nil {
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
func (_ERC1155NFT *ERC1155NFTFilterer) ParseTransferSingle(log types.Log) (*ERC1155NFTTransferSingle, error) {
	event := new(ERC1155NFTTransferSingle)
	if err := _ERC1155NFT.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155NFTURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ERC1155NFT contract.
type ERC1155NFTURIIterator struct {
	Event *ERC1155NFTURI // Event containing the contract specifics and raw log

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
func (it *ERC1155NFTURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155NFTURI)
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
		it.Event = new(ERC1155NFTURI)
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
func (it *ERC1155NFTURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155NFTURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155NFTURI represents a URI event raised by the ERC1155NFT contract.
type ERC1155NFTURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// ERC1155NFTURIOrChainReorg represents a URI subscription event raised by the ERC1155NFT contract.
type ERC1155NFTURIOrChainReorg struct {
	Event      *ERC1155NFTURI
	ChainReorg *types.ChainReorg
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155NFT *ERC1155NFTFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ERC1155NFTURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, err := _ERC1155NFT.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155NFTURIIterator{contract: _ERC1155NFT.contract, event: "URI", logs: logs}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155NFT *ERC1155NFTFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ERC1155NFTURIOrChainReorg, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC1155NFT.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155NFTURIOrChainReorg)
				event.Event = new(ERC1155NFTURI)

				if log.ChainReorg == nil {
					if err := _ERC1155NFT.contract.UnpackLog(event.Event, "URI", *log.Log); err != nil {
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
func (_ERC1155NFT *ERC1155NFTFilterer) ParseURI(log types.Log) (*ERC1155NFTURI, error) {
	event := new(ERC1155NFTURI)
	if err := _ERC1155NFT.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
