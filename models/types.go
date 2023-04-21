package models

type Pagination struct {
	Page  int `json:"page" form:"page" default:"1"`
	Limit int `json:"limit" form:"limit" default:"10"`
}

func (p Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}
