package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// HTTPError implements ClientError interface.
type HTTPError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  int    `json:"-"`
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%d %s, %s", e.Status, http.StatusText(e.Status), e.Message)
}

// ResponseBody returns JSON response body.
func (e *HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

func NewHTTPError(status int, code int, message string) error {
	return &HTTPError{
		Message: message,
		Code:    code,
		Status:  status,
	}
}

func ParseHttpError(status int, body []byte) (*HTTPError, error) {
	var h HTTPError
	if err := json.Unmarshal(body, &h); err != nil {
		return nil, err
	}

	if h.Code == 0 && h.Message == "" {
		return nil, errors.New("invalid http error")
	}

	h.Status = status
	return &h, nil
}
