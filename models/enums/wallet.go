package enums

import "fmt"

type WalletType uint

const (
	WALLET_ANYWEB WalletType = iota + 1
	WALLET_CELLAR
)

var (
	walletValue2StrMap map[WalletType]string
	walletStr2ValueMap map[string]WalletType
)

func init() {
	walletValue2StrMap = map[WalletType]string{
		WALLET_CELLAR: "cellar",
		WALLET_ANYWEB: "anyweb",
	}

	walletStr2ValueMap = make(map[string]WalletType)
	for k, v := range walletValue2StrMap {
		walletStr2ValueMap[v] = k
	}
}

func (t WalletType) String() string {
	v, ok := walletValue2StrMap[t]
	if ok {
		return v
	}
	return "unknown"
}

func (u *WalletType) UnmarshalText(data []byte) error {
	v, ok := walletStr2ValueMap[string(data)]
	if ok {
		*u = v
		return nil
	}
	return fmt.Errorf("unknown wallet type %v", string(data))
}

func ParseWalletType(str string) (*WalletType, bool) {
	v, ok := walletStr2ValueMap[str]
	return &v, ok
}
