package utils

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
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

func DrawLogo(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	src, _, err := image.Decode(response.Body)
	if err != nil {
		panic(err)
	}

	// 创建一个绘制对象
	rgba := image.NewRGBA(src.Bounds())
	draw.Draw(rgba, rgba.Bounds(), src, image.Point{0, 0}, draw.Src)

	// 在绘制对象上绘制文字
	point := fixed.Point26_6{fixed.Int26_6(rgba.Bounds().Min.X + 30*64), fixed.Int26_6(rgba.Bounds().Min.Y + 30*64)}
	d := &font.Drawer{
		Dst:  rgba,
		Src:  image.Black,
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString("Watermark")

	// 将绘制对象保存为新的图片
	newImg, err := os.Create("new.jpeg")
	if err != nil {
		panic(err)
	}
	defer newImg.Close()

	err = jpeg.Encode(newImg, rgba, nil)
	if err != nil {
		panic(err)
	}
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
