package models

import "gorm.io/gorm"

// Activity -- ActivityBindContract 1vs1
type ActivityBindContract struct {
	BaseModel
	// POAPActivityConfigID uint   `gorm:"column:poap_activity_config_id" json:"poap_activity_config_id"`
	ContractID      int32  `gorm:"uniqueIndex" json:"contract_id"` // rainbow-api contract id
	ContractAddress string `gorm:"type:string" json:"contract_address"`
	ContractType    int32  `gorm:"type:int" json:"contract_type"`
	ChainId         int32  `gorm:"type:int" json:"chain_id"`
	ChainType       int32  `gorm:"type:int" json:"chain_type"`
}

func FindContractByRawId(contractId uint) (*ActivityBindContract, error) {
	var contract ActivityBindContract
	if err := GetDB().Model(&contract).Where("contract_id=?", contractId).First(&contract).Error; err != nil {
		return nil, err
	}
	return &contract, nil
}

func FindContract(cond ActivityBindContract) ([]*ActivityBindContract, error) {
	var contracts []*ActivityBindContract
	if err := GetDB().Where(&cond).Find(&contracts).Error; err != nil {
		return nil, err
	}
	return contracts, nil
}

func FirstContract(cond ActivityBindContract) (*ActivityBindContract, error) {
	var contract *ActivityBindContract
	if err := GetDB().Where(&cond).First(&contract).Error; err != nil {
		return nil, err
	}
	return contract, nil
}

func UpdateOrCreateContract(contractId, contractType, chainId, chainType uint, contractAddress string) (*ActivityBindContract, error) {
	var contract ActivityBindContract
	c, err := FindContractByRawId(contractId)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	} else {
		contract = *c
	}

	contract.ContractID = int32(contractId)
	contract.ContractType = int32(contractType)
	contract.ChainId = int32(chainId)
	contract.ChainType = int32(chainType)
	contract.ContractAddress = contractAddress
	if err := GetDB().Save(&contract).Error; err != nil {
		return nil, err
	}
	return &contract, nil
}
