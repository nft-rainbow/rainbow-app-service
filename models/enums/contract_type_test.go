package enums

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalContractType(t *testing.T) {
	// var ct ContractType
	// err := json.Unmarshal([]byte(`"erc721"`), &ct)
	// assert.NoError(t, err)
	// fmt.Println(ct)

	// err = json.Unmarshal([]byte(``), &ct)
	// assert.NoError(t, err)
	// fmt.Println(ct)

	tmp := struct {
		ContractType ContractType `json:"contract_type"`
	}{}
	err := json.Unmarshal([]byte(`{"contract_type":""}`), &tmp)
	assert.NoError(t, err)
	fmt.Println(tmp)
}
