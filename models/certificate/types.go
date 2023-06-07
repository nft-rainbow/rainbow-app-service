package certificate

import "github.com/nft-rainbow/rainbow-app-service/models/enums"

type CertificatesQueryResult[T any] struct {
	Items           []T                   `json:"items"`
	Count           int64                 `json:"count"`
	CertificateType enums.CertificateType `json:"certificate_type" swaggertype:"string"`
}

func (c *CertificatesQueryResult[T]) ToAny() *CertificatesQueryResult[any] {
	itemsInAny := new(CertificateConverter[T]).ConvertSliceBack(c.Items)
	return &CertificatesQueryResult[any]{
		Items:           itemsInAny,
		Count:           c.Count,
		CertificateType: c.CertificateType,
	}
}

func (c *CertificatesQueryResult[T]) FromAny(input *CertificatesQueryResult[any]) error {
	itemsInType, err := new(CertificateConverter[T]).ConvertSlice(input.Items)
	if err != nil {
		return err
	}
	c.Items = itemsInType
	return nil
}
