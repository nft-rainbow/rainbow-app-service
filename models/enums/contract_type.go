package enums

import (
	"errors"
	"fmt"
)

type ContractType uint

const (
	ContractType_ERC721 ContractType = iota + 1
	ContractType_ERC1155
)

var (
	ContractTypeValue2StrMap map[ContractType]string
	ContractTypeStr2ValueMap map[string]ContractType
)

var (
	ErrUnkownContractType = errors.New("unknown contract type")
)

func init() {
	ContractTypeValue2StrMap = map[ContractType]string{
		ContractType_ERC721:  "erc721",
		ContractType_ERC1155: "erc1155",
	}

	ContractTypeStr2ValueMap = make(map[string]ContractType)
	for k, v := range ContractTypeValue2StrMap {
		ContractTypeStr2ValueMap[v] = k
	}
}

func (t ContractType) String() string {
	v, ok := ContractTypeValue2StrMap[t]
	if ok {
		return v
	}
	return "unkown"
}

func (t ContractType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *ContractType) UnmarshalText(data []byte) error {
	v, ok := ContractTypeStr2ValueMap[string(data)]
	if ok {
		*t = v
		return nil
	}
	return fmt.Errorf("unknown contract type %v", string(data))
}

func ParseContractType(str string) (*ContractType, error) {
	v, ok := ContractTypeStr2ValueMap[str]
	if !ok {
		return nil, fmt.Errorf("unknown contract type %v", str)
	}
	return &v, nil
}
