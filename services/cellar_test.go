package services

import (
	"testing"

	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/stretchr/testify/assert"
)

func TestGetAccount(t *testing.T) {
	c := NewCellarClient(enums.CHAIN_CONFLUX)
	resp, err := c.getAccount("13983211056")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestGetOrCreateAccount(t *testing.T) {
	c := NewCellarClient(enums.CHAIN_CONFLUX)
	resp, err := c.getOrCreateAccount("13983211056")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
