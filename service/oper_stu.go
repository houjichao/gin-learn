package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/houjichao/gin-learn/entity"
	"github.com/houjichao/gin-learn/library/orm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strings"
)

func Create(c *gin.Context) {
	//自动迁移
	//orm.AutoMigrate(&entity.Student{})

	var stu entity.Student
	err1 := c.ShouldBindJSON(&stu)
	if err1 != nil {
		fmt.Println("参数绑定错误:", err1)
		c.JSON(http.StatusInternalServerError, err1)
		return
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

func PageQuery(c *gin.Context) {
	rsp := PageQueryInter(c)
	c.JSON(http.StatusOK, gin.H{"data": rsp})
}

func PageQueryInter(c *gin.Context) map[string]interface{} {
	var stuQuery entity.StuQuery
	fmt.Println(stuQuery)
	err1 := c.ShouldBindJSON(&stuQuery)
	if err1 != nil {
		fmt.Println("参数绑定错误:", err1)
	}
	var whereTpl []string
	var args []interface{}
	if stuQuery.Name != "" {
		whereTpl = append(whereTpl, " name =? ")
		args = append(args, stuQuery.Name)
	}

	if stuQuery.ID > 0 {
		whereTpl = append(whereTpl, " id > ? ")
		args = append(args, stuQuery.ID)
	}

	where := entity.Where{
		Query: fmt.Sprintf(" %s", strings.Join(whereTpl, " and ")),
		Args:  args,
	}

	offset := (stuQuery.Page.PageIndex - 1) * stuQuery.Page.PageSize
	limit := stuQuery.Page.PageSize

	db := orm.DB
	var count int
	fmt.Println(entity.Student{}.TableName())
	db.Debug().Table(entity.Student{}.TableName()).Where(where.Query, where.Args...).Count(&count)

	db = db.Debug().Table(entity.Student{}.TableName()).Where(where.Query, where.Args...)
	if offset >= 0 {
		db = db.Offset(offset)
	}

	if limit >= 0 {
		db = db.Limit(limit)
	}

	var list []entity.Student
	err := db.Debug().Find(&list).Error

	if err != nil {
		fmt.Println("查询失败")
	}

	rsp := map[string]interface{}{
		"total": count,
		"list":  list,
	}

	return rsp
}
