package enums

import (
	"errors"
	"fmt"
)

type SourceType uint

const (
	SOURCE_TYPE_ADDRESS SourceType = iota + 1
	SOURCE_TYPE_PHONE
	SOURCE_TYPE_DODO
)

var (
	sourceTypeValue2StrMap map[SourceType]string
	sourceTypeStr2ValueMap map[string]SourceType
)

var (
	ErrUnkownSourceType = errors.New("unknown source type")
)

func init() {
	sourceTypeValue2StrMap = map[SourceType]string{
		SOURCE_TYPE_ADDRESS: "address",
		SOURCE_TYPE_PHONE:   "phone",
	}

	sourceTypeStr2ValueMap = make(map[string]SourceType)
	for k, v := range sourceTypeValue2StrMap {
		sourceTypeStr2ValueMap[v] = k
	}
}

func (s SourceType) String() string {
	v, ok := sourceTypeValue2StrMap[s]
	if ok {
		return v
	}
	return "unkown"
}

func (s SourceType) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *SourceType) UnmarshalText(data []byte) error {
	v, ok := sourceTypeStr2ValueMap[string(data)]
	if ok {
		*s = v
		return nil
	}
	return fmt.Errorf("unknown social tool type %v", string(data))
}

func ParseSourceType(str string) (*SourceType, error) {
	v, ok := sourceTypeStr2ValueMap[str]
	if !ok {
		return nil, fmt.Errorf("unknown social tool type %v", str)
	}
	return &v, nil
}
