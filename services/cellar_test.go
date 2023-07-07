package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrCreateAccount(t *testing.T) {
	c := NewCellarClient()
	resp, err := c.getOrCreateAccount("13983211056")
	assert.NoError(t, err)
	assert.NotNil(t, resp.Data)
}
