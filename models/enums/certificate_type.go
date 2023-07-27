package enums

import (
	"errors"
	"fmt"
)

type CertificateType uint

const (
	CERTIFICATE_ADDRESS CertificateType = iota + 1
	CERTIFICATE_PHONE
	CERTIFICATE_DODO
	CERTIFICATE_CONTRACT
	CERTIFICATE_GASLESS
)

var (
	certificateTypeValue2StrMap map[CertificateType]string
	certificateTypeStr2ValueMap map[string]CertificateType
)

var (
	ErrUnkownCertificateType = errors.New("unknown certificate type")
)

func init() {
	certificateTypeValue2StrMap = map[CertificateType]string{
		CERTIFICATE_ADDRESS:  "address",
		CERTIFICATE_PHONE:    "phone",
		CERTIFICATE_DODO:     "dodo",
		CERTIFICATE_CONTRACT: "contract",
		CERTIFICATE_GASLESS:  "gasless",
	}

	certificateTypeStr2ValueMap = make(map[string]CertificateType)
	for k, v := range certificateTypeValue2StrMap {
		certificateTypeStr2ValueMap[v] = k
	}
}

func (t CertificateType) String() string {
	v, ok := certificateTypeValue2StrMap[t]
	if ok {
		return v
	}
	return "unkown"
}

func (t CertificateType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *CertificateType) UnmarshalText(data []byte) error {
	v, ok := certificateTypeStr2ValueMap[string(data)]
	if ok {
		*t = v
		return nil
	}
	return fmt.Errorf("unknown certificate type %v", string(data))
}

func ParseCertificateType(str string) (*CertificateType, error) {
	v, ok := certificateTypeStr2ValueMap[str]
	if !ok {
		return nil, fmt.Errorf("unknown certificate type %v", str)
	}
	return &v, nil
}
