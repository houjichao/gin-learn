package service

import (
	"github.com/gin-gonic/gin"
	"github.com/houjichao/gin-learn/entity"
	"net/http"
)

func Demo(c *gin.Context) {
	// 声明接收的变量
	var json entity.Login
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if json.User != "root" || json.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}
