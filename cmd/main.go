package main

import (
	"blog/apps/login/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	c := controller.LoginController{}
	r := gin.Default()
	route := r.Group("/login")
	{
		route.GET("/login", c.Login)
	}
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("framework failed!")
	}
}
