package utils

import (
	"bytes"
	"encoding/json"

	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func SendHttp[TPayload any, TResp any](method string, url string, payload TPayload, headers map[string]string) (*TResp, error) {
	payloadJ, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(payloadJ))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		httpErr, err := ParseHttpError(resp.StatusCode, body)
		if err != nil {
			return nil, errors.Errorf("%s, %s", resp.Status, string(body))
		}

		return nil, httpErr
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	logrus.WithField("resp.Body", string(body)).Trace("Received HTTP response")

	var cr TResp
	if err := json.Unmarshal(body, &cr); err != nil {
		return nil, err
	}

	return &cr, nil
}
