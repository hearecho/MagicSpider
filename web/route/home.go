package route

import (
	"MagicSpider/web/controller"
	"github.com/gin-gonic/gin"
)

func LoadHomeRoute(e *gin.Engine)  {
	e.GET("/",controller.IndexHandler)
	e.GET("/index",controller.HomeHandler)
}
