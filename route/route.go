package route

import (
	"github.com/gin-gonic/gin"
	"github.com/houjichao/gin-learn/service"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewAndSetup(r *gin.Engine) {
	/*router := gin.New()

	demoGroup := router.Group("/demo")
	{
		demoGroup.POST("/loginJSON", service.Demo)
	}*/
	r.POST("/demo/loginJSON", service.Demo)

	r.POST("/stu/create", service.Create)

	r.GET("/stu/queryById/:id", service.Query)

}
