package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/houjichao/gin-learn/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func Create(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/learn?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}
	defer db.Close()

	//自动迁移
	//db.AutoMigrate(&entity.Student{})

	//关闭复数表名
	db.SingularTable(true)

	var stu entity.Student
	err1 := c.ShouldBindJSON(&stu)
	if err1 != nil {
		fmt.Println("参数绑定错误:", err)
	}
	//创建记录
	var result = db.Create(&stu)
	c.JSON(http.StatusOK, gin.H{"data": result})
	return
}

func Query(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/learn?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}
	defer db.Close()

	//自动迁移
	//db.AutoMigrate(&entity.Student{})

	//关闭复数表名
	db.SingularTable(true)

	var stu entity.Student
	err1 := c.ShouldBindUri(&stu)
	if err1 != nil {
		fmt.Println("参数绑定错误:", err)
	}
	db.First(&stu, stu.Id)
	c.JSON(http.StatusOK, gin.H{"data": stu})
	return
}
