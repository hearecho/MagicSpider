package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Age int
	Name string
}
func IndexHandler(c *gin.Context)  {
	u := User{
		Age:  12,
		Name: "echo",
	}
	c.JSON(http.StatusOK,gin.H{
		"title":"测试",
		"data":u,
	})
}

func HomeHandler(c *gin.Context)  {
	u := User{
		Age:  12,
		Name: "echo",
	}
	c.HTML(http.StatusOK,"home.html",gin.H{
		"title":"测试",
		"data":u,
	})
}
