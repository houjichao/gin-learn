package entity

// 定义学生结构体
type Student struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Id   int    `gorm:"column:id" json:"id" uri:"id" gorm:"AUTO_INCREMENT"`
	Name string `gorm:"column:name" json:"name" binding:"required"`
	Age  int `gorm:"column:age" json:"age" binding:"required"`
}
