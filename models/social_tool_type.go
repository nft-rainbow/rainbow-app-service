package models

import (
	"encoding/json"
	"errors"
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

func (t SocialToolType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *SocialToolType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	val, ok := ParseSocialToolType(str)
	if !ok {
		return errors.New("unkown social_tool_type")
	}
	*t = *val

	return nil
}

func (t SocialToolType) String() string {
	v, ok := socialTypeValue2StrMap[t]
	if ok {
		return v
	}
	return "UNKNOWN"
}

func ParseSocialToolType(str string) (*SocialToolType, bool) {
	v, ok := socialTypeStr2ValueMap[str]
	return &v, ok
}
