package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"github.com/houjichao/gin-learn/entity"
	"github.com/houjichao/gin-learn/library/orm"
	"github.com/houjichao/gin-learn/route"
	"gopkg.in/ini.v1"
	"log"
)

var global int

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

	//自定义校验方法注册
	// 3、将我们自定义的校验方法注册到 validator中  ---- 自定义验证这里没生效，待排查
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 这里的 key 和 fn 可以不一样最终在 struct 使用的是 key
		v.RegisterValidation("NotNullAndAdmin", entity.NameNotNullAndAdmin)
	}

	// 接口绑定
	route.NewAndSetup(r)

	r.Run(":8000")
}
