package entity

type Page struct {
	PageIndex       int    `json:"page_index"`
	PageSize        int    `json:"page_size" validate:"max=10000"`
}
