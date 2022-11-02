package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/houjichao/gin-learn/route"
)

func main() {

	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()

	r := gin.Default()
	fmt.Printf("%T\n",r)
	// 接口绑定
	route.NewAndSetup(r)
	r.Run(":8000")
}
