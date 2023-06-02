package enums

import (
	"errors"
	"fmt"
)

type SnapshotStatus uint

const (
	SNAPSHOT_STATUS_INIT SnapshotStatus = iota
	SNAPSHOT_STATUS_PROCESSING
	SNAPSHOT_STATUS_SUCCESS
	SNAPSHOT_STATUS_FAILED
)

var (
	SnapshotStatusValue2StrMap map[SnapshotStatus]string
	SnapshotStatusStr2ValueMap map[string]SnapshotStatus
)

var (
	ErrUnkownSnapshotStatus = errors.New("unknown snapshot status")
)

func init() {
	SnapshotStatusValue2StrMap = map[SnapshotStatus]string{
		SNAPSHOT_STATUS_INIT:       "init",
		SNAPSHOT_STATUS_PROCESSING: "processing",
		SNAPSHOT_STATUS_SUCCESS:    "success",
		SNAPSHOT_STATUS_FAILED:     "failed",
	}

	SnapshotStatusStr2ValueMap = make(map[string]SnapshotStatus)
	for k, v := range SnapshotStatusValue2StrMap {
		SnapshotStatusStr2ValueMap[v] = k
	}
}

func (t SnapshotStatus) String() string {
	v, ok := SnapshotStatusValue2StrMap[t]
	if ok {
		return v
	}
	return "unkown"
}

func (t SnapshotStatus) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *SnapshotStatus) UnmarshalText(data []byte) error {
	v, ok := SnapshotStatusStr2ValueMap[string(data)]
	if ok {
		*t = v
		return nil
	}
	return fmt.Errorf("unknown snapshot status %v", string(data))
}

func ParseSnapshotStatus(str string) (*SnapshotStatus, error) {
	v, ok := SnapshotStatusStr2ValueMap[str]
	if !ok {
		return nil, fmt.Errorf("unknown snapshot status %v", str)
	}
	return &v, nil
}
