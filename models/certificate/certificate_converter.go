package certificate

import "encoding/json"

type CertificateConverter[T any] struct {
}

func (c *CertificateConverter[T]) Convert(source any) (T, error) {
	j, err := json.Marshal(source)
	if err != nil {
		return *new(T), err
	}
	var dest T
	err = json.Unmarshal(j, &dest)
	if err != nil {
		return *new(T), err
	}
	return dest, nil
}

func (c *CertificateConverter[T]) ConvertSlice(source []any) ([]T, error) {
	var result []T
	for _, s := range source {
		out, err := c.Convert(s)
		if err != nil {
			return nil, err
		}
		result = append(result, out)
	}
	return result, nil
}

func (c *CertificateConverter[T]) ConvertSliceBack(source []T) []any {
	var result []any
	for _, s := range source {
		result = append(result, s)
	}
	return result
}
