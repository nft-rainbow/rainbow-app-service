package models

type H5Config struct {
	BaseModel
	ActivityId       string `gorm:"type:integer;index" json:"activity_id"`
	Link             string `gorm:"type:string" json:"link" binding:"required"`
	Title            string `gorm:"type:string" json:"title"`
	TitleSize        int32  `gorm:"type:integer" json:"title_size"`
	TitleColor       int32  `gorm:"type:integer" json:"title_color"`
	Content          string `gorm:"type:string" json:"content"`
	ContentSize      int32  `gorm:"type:integer" json:"content_size"`
	ContentColor     int32  `gorm:"type:integer" json:"content_color"`
	ClaimButtonColor string `gorm:"type:string" json:"claim_button_color"`
	ButtonWordColor  int32  `gorm:"type:string" json:"button_word_color"`
	LogoURL          string `gorm:"type:string" json:"logo_url"`
	PCPicURL         string `gorm:"type:string" json:"pc_picture_url"`
	MobilePicURL     string `gorm:"type:string" json:"mobile_picture_url"`
}
