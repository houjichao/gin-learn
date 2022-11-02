package service

import (
	"fmt"
	"github.com/houjichao/gin-learn/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Create(stu entity.Student) {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/learn?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}
	defer db.Close()

	//自动迁移
	db.AutoMigrate(&entity.Student{})

	//关闭复数表名
	db.SingularTable(true)

	//创建记录
	db.Create(stu)
}
