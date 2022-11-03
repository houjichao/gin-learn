package entity

type StuQuery struct {
	Page Page `json:"page"`
	ID   int `json:"id"`
	Name string `json:"name"`
}
