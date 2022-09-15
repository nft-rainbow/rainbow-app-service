package models

type Metadata struct {
	BaseModel `swaggerignore:"true"`
	AppId        uint         `gorm:"index" json:"app_id" swaggerignore:"true"`
	Name         string       `gorm:"type:varchar(256)" json:"name" binding:"required"`
	Description  string       `gorm:"type:varchar(256)" json:"description" binding:"required"`
	ExternalLink string       `gorm:"type:varchar(256)" json:"external_link" swaggo:"false"`
	Image        string       `gorm:"type:varchar(256)"  json:"image" binding:"required"`
	Attributes   []Attributes `gorm:"foreignkey:MetadataId" json:"attributes" swaggo:"false"`
	NftAddress   string       `gorm:"type:varchar(256)"  json:"nft_address" swaggo:"false"`
	ID           string       `gorm:"column:meta_data_id;type:varchar(256)" json:"metadata_id" swaggerignore:"true"`
	URI          string       `json:"uri" swaggerignore:"true"`
}

type Attributes struct {
	BaseModel `swaggerignore:"true"`
	Name        string `gorm:"type:varchar(256)"  json:"attribute_name"`
	TraitType   string `gorm:"type:varchar(256)"  json:"trait_type"`
	DisplayType string `gorm:"type:varchar(256)"  json:"display_type"`
	Value       string `gorm:"type:varchar(256)"  json:"value"`
	MetadataId  uint
}

type CreateMetadataResponse struct {
	metadata    Metadata
	MetadataURI string `json:"uri"`
	Message string `json:"message"`
}