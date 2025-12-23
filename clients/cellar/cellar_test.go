package cellar

import (
	"fmt"
	"testing"

	"github.com/nft-rainbow/rainbow-app-service/config"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/stretchr/testify/assert"
)

func TestGetAccount(t *testing.T) {
	config.InitByFile("../../config.yaml")
	c, err := NewCellarClient(enums.CHAIN_CONFLUX)
	assert.NoError(t, err)
	resp, err := c.GetAccount("31b408afd5394a7dbe59831bfc21f764")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestGetOrCreateAccount(t *testing.T) {
	config.InitByFile("../../config.yaml")
	c, err := NewCellarClient(enums.CHAIN_CONFLUX)
	assert.NoError(t, err)
	resp, err := c.GetOrCreateAccount("13983211056")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp)
}
