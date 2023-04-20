package models

type Pagination struct {
	Page  int `json:"page" form:"page" default:"1"`
	Limit int `json:"limit" form:"limit" default:"10"`
}
