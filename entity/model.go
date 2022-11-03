package entity

// Where gorm Where条件
type Where struct {
	Query interface{}
	Args  []interface{}
}