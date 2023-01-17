package utils

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
)

var (
	U256Max = MustParseStrToBig("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
)

func Bytes2Hex(data []byte) string {
	return "0x" + common.Bytes2Hex(data)
}

func CheckCfxAddress(chain string, addr string) (*cfxaddress.Address, error) {
	chainType, chainId, err := ChainInfoByName(chain)
	if err != nil {
		return nil, err
	}
	if chainType != CHAIN_TYPE_CFX {
		return nil, errors.New("not cfx chain")
	}
	addrItem, err := cfxaddress.NewFromBase32(addr)
	if err != nil {
		return nil, err
	}
	if addrItem.GetNetworkID() != uint32(chainId) {
		return nil, fmt.Errorf("invalid conflux network address, want %v, got %v", addrItem.GetNetworkID(), uint32(chainId))
	}
	return &addrItem, nil
}

func IsCfxAddress(addr string) error {
	tmp := strings.Split(addr, ":")
	if tmp[0] == "cfx" {
		_, err := CheckCfxAddress(CONFLUX, addr)
		if err != nil {
			return err
		}
	}else if tmp[0] == "cfxtest" {
		_, err := CheckCfxAddress(CONFLUX_TEST, addr)
		if err != nil {
			return err
		}
	}else {
		return fmt.Errorf("Invalid conflux network")
	}
	return nil
}

func CurrentMonthStr() string {
	now := time.Now()
	return fmt.Sprintf("%d-%d", now.Year(), now.Month())
}

func UintPtrToBig(val *uint) *big.Int {
	var result *big.Int
	if val != nil {
		result = big.NewInt(int64(*val))
	}
	return result
}

func Uint64PtrToBig(val *uint64) *big.Int {
	var result *big.Int
	if val != nil {
		result = new(big.Int).SetUint64(*val)
	}
	return result
}

func UintPtrToUint(val *uint) uint {
	result := uint(0)
	if val != nil {
		result = uint(*val)
	}
	return result
}

func Uint64Ptr(val uint64) *uint64 {
	return &val
}

func UintPtr(val uint) *uint {
	return &val
}

func MustParseStrToBig(s string) *big.Int {
	val, ok := new(big.Int).SetString(s, 0)
	if !ok {
		panic(fmt.Sprintf("failed to parse %s as big int", s))
	}
	return val
}

func InUint256(val *big.Int) bool {
	return val.BitLen() <= 256
}

func MustNewBigIntByString(val string) *big.Int {
	b, _ := new(big.Int).SetString(val, 0)
	return b
}
