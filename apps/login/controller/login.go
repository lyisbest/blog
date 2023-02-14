package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
}

func (c LoginController) Login(ctx *gin.Context) {
	userName := ctx.Query("username")
	password := ctx.Query("password")

	fmt.Println(userName, password)
}
