package utils

import (
	"errors"
	"fmt"
	_ "image/gif"
	_ "image/png"
	"math/big"
	"runtime/debug"
	"strings"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
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
	_, err := cfxaddress.NewFromBase32(addr)
	return err
}

func SimpleAddress(addr string) string {
	tmp := strings.Split(addr, ":")
	return tmp[0] + ":" + tmp[1][:3] + "..." + tmp[1][:3]
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

func TodayDateStr() string {
	now := time.Now()
	return fmt.Sprintf("%04d-%02d-%02d", now.Year(), now.Month(), now.Day())
}

func TomorrowBegin() time.Time {
	t := time.Now().Add(time.Hour * 24)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func Retry(count int, interval time.Duration, fn func() error) error {
	for i := 0; i < count; i++ {
		err := fn()
		if err != nil {
			logrus.WithError(err).WithField("stack", string(debug.Stack())).WithField("retry cnt", i).Info("retry function error")
			if i == count-1 {
				return err
			} else {
				time.Sleep(interval)
				continue
			}
		}
		return nil
	}
	return nil
}

func Bool2Int(val bool) int {
	if val {
		return 1
	}
	return 0
}

func MapSlice[T any, R any](items []T, mapFunc func(item T) (R, error)) ([]R, error) {
	var result []R
	for _, item := range items {
		r, err := mapFunc(item)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}
