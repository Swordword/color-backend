package routers

import (
	"colorist/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

var origin = "www.baidu.com"

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors())

	v1Group := r.Group("v1")
	{
		v1Group.POST("/color", controller.CreateAColor)
		v1Group.GET("/color/list", controller.GetColorList)
		v1Group.DELETE("/color/:id", controller.DeleteAColor)
		v1Group.POST("/star/:id", controller.StarAColor)
		v1Group.DELETE("/star/:id", controller.CancelStar)
	}

	return r
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
