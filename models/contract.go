package models

import "gorm.io/gorm"

// Activity -- Contract 1vs1
type Contract struct {
	BaseModel
	// POAPActivityConfigID uint   `gorm:"column:poap_activity_config_id" json:"poap_activity_config_id"`
	ContractRawID   int32  `gorm:"uniqueIndex" json:"contract_raw_id"` // rainbow-api contract id
	ContractAddress string `gorm:"type:string" json:"contract_address"`
	ContractType    int32  `gorm:"type:int" json:"contract_type"`
	ChainId         int32  `gorm:"type:int" json:"chain_id"`
	ChainType       int32  `gorm:"type:int" json:"chain_type"`
}

func FindContractByRawId(contractRawId uint) (*Contract, error) {
	var contract Contract
	if err := GetDB().Model(&contract).Where("contract_raw_id=?", contractRawId).First(&contract).Error; err != nil {
		return nil, err
	}
	return &contract, nil
}

func FindContract(cond Contract) ([]*Contract, error) {
	var contracts []*Contract
	if err := GetDB().Where(&cond).Find(&contracts).Error; err != nil {
		return nil, err
	}
	return contracts, nil
}

func FirstContract(cond Contract) (*Contract, error) {
	var contract *Contract
	if err := GetDB().Where(&cond).First(&contract).Error; err != nil {
		return nil, err
	}
	return contract, nil
}

func UpdateOrCreateContract(contractId, contractType, chainId, chainType uint, contractAddress string) (*Contract, error) {
	var contract Contract
	c, err := FindContractByRawId(contractId)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	} else {
		contract = *c
	}

	contract.ContractRawID = int32(contractId)
	contract.ContractType = int32(contractType)
	contract.ChainId = int32(chainId)
	contract.ChainType = int32(chainType)
	contract.ContractAddress = contractAddress
	if err := GetDB().Save(&contract).Error; err != nil {
		return nil, err
	}
	return &contract, nil
}
