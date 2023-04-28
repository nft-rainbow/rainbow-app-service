package enums

import (
	"errors"
	"fmt"
)

type SocialToolType uint

const (
	SOCIAL_TOOL_DISCORD SocialToolType = iota + 1
	SOCIAL_TOOL_DODO
)

var (
	socialTypeValue2StrMap map[SocialToolType]string
	socialTypeStr2ValueMap map[string]SocialToolType
)

var (
	ErrUnkownSocialType = errors.New("unknown trade type")
)

func init() {
	socialTypeValue2StrMap = map[SocialToolType]string{
		SOCIAL_TOOL_DISCORD: "discord",
		SOCIAL_TOOL_DODO:    "dodo",
	}

	socialTypeStr2ValueMap = make(map[string]SocialToolType)
	for k, v := range socialTypeValue2StrMap {
		socialTypeStr2ValueMap[v] = k
	}
}

func (s SocialToolType) String() string {
	v, ok := socialTypeValue2StrMap[s]
	if ok {
		return v
	}
	return "unkown"
}

func (s SocialToolType) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *SocialToolType) UnmarshalText(data []byte) error {
	v, ok := socialTypeStr2ValueMap[string(data)]
	if ok {
		*s = v
		return nil
	}
	return fmt.Errorf("unknown social tool type %v", string(data))
}

func ParseSocialToolType(str string) (*SocialToolType, error) {
	v, ok := socialTypeStr2ValueMap[str]
	if !ok {
		return nil, fmt.Errorf("unknown social tool type %v", str)
	}
	return &v, nil
}
