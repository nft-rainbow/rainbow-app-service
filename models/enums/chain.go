package enums

import (
	"errors"
	"fmt"
)

type Chain uint

const (
	CHAIN_CONFLUX Chain = iota + 1
	CHAIN_CONFLUX_TEST
)

var (
	ChainValue2StrMap map[Chain]string
	ChainStr2ValueMap map[string]Chain
)

var (
	ErrUnkownChain = errors.New("unknown chain")
)

func init() {
	ChainValue2StrMap = map[Chain]string{
		CHAIN_CONFLUX:      "conflux",
		CHAIN_CONFLUX_TEST: "conflux_test",
	}

	ChainStr2ValueMap = make(map[string]Chain)
	for k, v := range ChainValue2StrMap {
		ChainStr2ValueMap[v] = k
	}
}

func (t Chain) String() string {
	v, ok := ChainValue2StrMap[t]
	if ok {
		return v
	}
	return "unkown"
}

func (t Chain) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *Chain) UnmarshalText(data []byte) error {
	v, ok := ChainStr2ValueMap[string(data)]
	if ok {
		*t = v
		return nil
	}
	return fmt.Errorf("unknown chain %v", string(data))
}

func ParseChain(str string) (*Chain, error) {
	v, ok := ChainStr2ValueMap[str]
	if !ok {
		return nil, fmt.Errorf("unknown chain %v", str)
	}
	return &v, nil
}
