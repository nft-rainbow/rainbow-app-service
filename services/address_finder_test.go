package services

import (
	"fmt"
	"testing"

	"github.com/nft-rainbow/rainbow-app-service/models"
	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/stretchr/testify/assert"
)

func TestFindAddressByPhones(t *testing.T) {
	initConfig()
	models.ConnectDB()

	af := AddressFinder{SourceType: enums.SOURCE_TYPE_PHONE}
	exists, unexist, err := af.Find(enums.CHAIN_CONFLUX_TEST, []string{"17011112223", "18656303977"})
	assert.NoError(t, err)

	fmt.Printf("exists %v,\nunexist %v", exists, unexist)
}
