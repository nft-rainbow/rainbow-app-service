package models

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/nft-rainbow/rainbow-app-service/models/enums"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalWalletType(t *testing.T) {
	var w *enums.WalletType
	err := json.Unmarshal([]byte("\"cellar\""), &w)
	assert.NoError(t, err)
	assert.Equal(t, enums.WALLET_CELLAR, *w)
	fmt.Println(*w)

	type AddWalletUserReq struct {
		Wallet  enums.WalletType `json:"wallet"`
		Code    string           `json:"code"`
		Phone   string           `json:"phone"`
		Address string           `json:"address"`
	}
	var req *AddWalletUserReq
	err = json.Unmarshal([]byte(`{
		"wallet":"cellar",
		"phone": "17011112222",
		"address": "cfx:aamgvyzht7h1zxdghb9ee9w26wrz8rd3gj837392dp"
	}`), &req)
	assert.NoError(t, err)
	assert.Equal(t, enums.WALLET_CELLAR, req.Wallet)
	fmt.Println(*req)
}
