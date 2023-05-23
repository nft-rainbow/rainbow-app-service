package models

import (
	"github.com/pkg/errors"
	"gorm.io/datatypes"
)

type (
	NftConfigUpdatePart struct {
		ImageURL           string                                 `gorm:"type:string" json:"image_url"`
		Name               string                                 `gorm:"type:string" json:"name"`
		Probability        float32                                `gorm:"type:float" json:"probability"`
		MetadataAttributes datatypes.JSONSlice[MetadataAttribute] `gorm:"type:json" json:"metadata_attributes"`
	}

	NFTConfig struct {
		BaseModel
		ActivityID uint
		NftConfigUpdatePart
	}

	MetadataAttribute struct {
		TraitType   string `gorm:"type:varchar(256)"  json:"trait_type"`
		DisplayType string `gorm:"type:varchar(256)"  json:"display_type,omitempty"`
		Value       string `gorm:"type:varchar(256)"  json:"value"`
	}
)

func CheckNftConfigBelongToUser(nftConfig *NFTConfig, userId uint) error {
	var activity Activity
	err := GetDB().First(&activity, nftConfig.ActivityID).Error
	if err != nil {
		return err
	}

	if userId != activity.RainbowUserId {
		return errors.New("the nft config not belongs to the user")
	}
	return nil
}

func FindNftConfigById(id uint) (*NFTConfig, error) {
	var nftConfig NFTConfig
	err := GetDB().First(&nftConfig, id).Error
	if err != nil {
		return nil, err
	}
	return &nftConfig, nil
}
