package controller

import (
	"blog/apps/login/constant"
	"blog/apps/login/request"
	"blog/apps/login/service"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginService service.LoginService
}

func NewLoginController(loginService service.LoginService) *LoginController {
	return &LoginController{loginService: loginService}
}

func (c LoginController) Login(ctx *gin.Context) {
	var loginRequest request.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		utils.ResponseWithError(ctx, constant.ResolveError)
		return
	}
	err := c.loginService.Login(ctx, loginRequest.UserName, loginRequest.Password)
	if err != nil {
		utils.ResponseWithError(ctx, constant.LogFailError)
		return
	}
	utils.ResponseNormal(ctx)
}
