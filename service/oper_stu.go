package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/houjichao/gin-learn/entity"
	"github.com/houjichao/gin-learn/library/orm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func Create(c *gin.Context) {
	//自动迁移
	//orm.AutoMigrate(&entity.Student{})

	var stu entity.Student
	err1 := c.ShouldBindJSON(&stu)
	if err1 != nil {
		fmt.Println("参数绑定错误:", err1)
	}
	//创建记录
	var result = orm.DB.Create(&stu)
	c.JSON(http.StatusOK, gin.H{"data": result})
	return
}

func Update(c *gin.Context) {

	var stu entity.Student
	err1 := c.ShouldBindJSON(&stu)
	if err1 != nil {
		fmt.Println("参数绑定错误:", err1)
	}
	//更新操作
	//提示: 相当于根据主键id，更新所有模型字段值。
	orm.DB.Save(&stu)
	//更新单个字段值
	//orm.Model(&stu).Update("age", stu.Age)
	//其他的method等使用时再查api
	c.JSON(http.StatusOK, gin.H{"data": stu})
	return
}

func Query(c *gin.Context) {

	var stu entity.Student
	err1 := c.ShouldBindUri(&stu)
	if err1 != nil {
		fmt.Println("参数绑定错误:", err1)
	}
	orm.DB.First(&stu, stu.Id)
	c.JSON(http.StatusOK, gin.H{"data": stu})
	return
}
