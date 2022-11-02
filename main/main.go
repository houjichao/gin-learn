package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/houjichao/gin-learn/library/orm"
	"github.com/houjichao/gin-learn/route"
	"gopkg.in/ini.v1"
	"log"
)

func main() {

	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()

	r := gin.Default()
	fmt.Printf("%T\n", r)
	// 数据库初始化
	// 读取配置文件
	conf, err := ini.Load("./my.ini")
	if err != nil {
		log.Fatal("配置文件读取失败, err = ", err)
	}

	db := orm.InitDB(conf)
	defer db.Close()

	// 接口绑定
	route.NewAndSetup(r)
	r.Run(":8000")
}
