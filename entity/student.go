package entity

import (
	"github.com/go-playground/validator"
	"time"
)

// 定义学生结构体
type Student struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Id   int    `gorm:"column:id" json:"id" uri:"id" gorm:"AUTO_INCREMENT"`
	Name string `gorm:"column:name" json:"name" validate:"NotNullAndAdmin"`
	//结构体验证
	//不能为空且大于6
	Age int `gorm:"column:age" json:"age" binding:"required,gt=6"`
}

//如果表名和结构体名有不同，则通过实现该结构体的TableName方法指定表名
// TableName 表名字
func (Student) TableName() string {
	return "t_student"
}

// 1、自定义的校验方法
var NameNotNullAndAdmin validator.Func = func(fl validator.FieldLevel) bool {

	if value, ok := fl.Field().Interface().(string); ok {
		// 字段不能为空，并且不等于  admin
		return value != "" && !("admin" == value)
	}

	return true
}

var BookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}
