package enums

import (
	"errors"
	"fmt"
)

type ActivityType uint

const (
	ACTIVITY_BLINDBOX ActivityType = iota + 1
	ACTIVITY_SINGLE
	// ACTIVITY_SINGLE_ID_ORDER
	// ACTIVITY_POAP
)

var (
	activityTypeValue2StrMap map[ActivityType]string
	activityTypeStr2ValueMap map[string]ActivityType
)

var (
	ErrUnkownActivityType = errors.New("unknown activity type")
)

func init() {
	activityTypeValue2StrMap = map[ActivityType]string{
		ACTIVITY_BLINDBOX: "blind_box",
		ACTIVITY_SINGLE:   "single",
		// ACTIVITY_SINGLE_ID_ORDER: "single_id_order",
		// ACTIVITY_POAP: "poap",
	}

	activityTypeStr2ValueMap = make(map[string]ActivityType)
	for k, v := range activityTypeValue2StrMap {
		activityTypeStr2ValueMap[v] = k
	}
}

func (t ActivityType) String() string {
	v, ok := activityTypeValue2StrMap[t]
	if ok {
		return v
	}
	return "unkown"
}

func (t ActivityType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *ActivityType) UnmarshalText(data []byte) error {
	v, ok := activityTypeStr2ValueMap[string(data)]
	if ok {
		*t = v
		return nil
	}
	return fmt.Errorf("unknown activity type %v", string(data))
}

func ParseActivityType(str string) (*ActivityType, error) {
	v, ok := activityTypeStr2ValueMap[str]
	if !ok {
		return nil, fmt.Errorf("unknown activity type %v", str)
	}
	return &v, nil
}
