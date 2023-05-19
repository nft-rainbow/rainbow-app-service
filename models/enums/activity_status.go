package enums

import (
	"errors"
	"fmt"
)

type ActivityStatus uint

const (
	ACTIVITY_STATUS_UNSTART ActivityStatus = iota + 1
	ACTIVITY_STATUS_ONGOING
	ACTIVITY_SINGLE_END
)

var (
	activityStatusValue2StrMap map[ActivityStatus]string
	activityStatusStr2ValueMap map[string]ActivityStatus
)

var (
	ErrUnkownactivityStatus = errors.New("unknown activity status")
)

func init() {
	activityStatusValue2StrMap = map[ActivityStatus]string{
		ACTIVITY_STATUS_UNSTART: "unstart",
		ACTIVITY_STATUS_ONGOING: "ongoing",
		ACTIVITY_SINGLE_END:     "end",
	}

	activityStatusStr2ValueMap = make(map[string]ActivityStatus)
	for k, v := range activityStatusValue2StrMap {
		activityStatusStr2ValueMap[v] = k
	}
}

func (t ActivityStatus) String() string {
	v, ok := activityStatusValue2StrMap[t]
	if ok {
		return v
	}
	return "unkown"
}

func (t ActivityStatus) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *ActivityStatus) UnmarshalText(data []byte) error {
	v, ok := activityStatusStr2ValueMap[string(data)]
	if ok {
		*t = v
		return nil
	}
	return fmt.Errorf("unknown activity status %v", string(data))
}

func ParseActivityStatus(str string) (*ActivityStatus, error) {
	v, ok := activityStatusStr2ValueMap[str]
	if !ok {
		return nil, fmt.Errorf("unknown activity status %v", str)
	}
	return &v, nil
}
