package main

import (
	"MagicSpider/web/route"
	"fmt"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()
func init()  {
	route.LoadHomeRoute(r)
	r.LoadHTMLGlob("static/templates/**/*")
}
func main() {

	if err := r.Run(":81"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}