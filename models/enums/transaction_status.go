package enums

import (
	"errors"
	"fmt"
)

// const (
// 	STATUS_INIT = iota
// 	STATUS_SUCCESS
// 	STATUS_FAIL
// )

type TransactionStatus uint

const (
	TRANSACTION_STATUS_INIT TransactionStatus = iota
	TRANSACTION_STATUS_SUCCESS
	TRANSACTION_STATUS_FAILED
)

var (
	TransactionStatusValue2StrMap map[TransactionStatus]string
	TransactionStatusStr2ValueMap map[string]TransactionStatus
)

var (
	ErrUnkownTransactionStatus = errors.New("unknown transaction status")
)

func init() {
	TransactionStatusValue2StrMap = map[TransactionStatus]string{
		TRANSACTION_STATUS_INIT:    "init",
		TRANSACTION_STATUS_SUCCESS: "success",
		TRANSACTION_STATUS_FAILED:  "failed",
	}

	TransactionStatusStr2ValueMap = make(map[string]TransactionStatus)
	for k, v := range TransactionStatusValue2StrMap {
		TransactionStatusStr2ValueMap[v] = k
	}
}

func (t TransactionStatus) String() string {
	v, ok := TransactionStatusValue2StrMap[t]
	if ok {
		return v
	}
	return "unkown"
}

func (t TransactionStatus) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *TransactionStatus) UnmarshalText(data []byte) error {
	v, ok := TransactionStatusStr2ValueMap[string(data)]
	if ok {
		*t = v
		return nil
	}
	return fmt.Errorf("unknown transaction status %v", string(data))
}

func ParseTransactionStatus(str string) (*TransactionStatus, error) {
	v, ok := TransactionStatusStr2ValueMap[str]
	if !ok {
		return nil, fmt.Errorf("unknown transaction status %v", str)
	}
	return &v, nil
}
