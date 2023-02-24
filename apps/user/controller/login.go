package controller

import (
	"blog/apps/user/request"
	"blog/apps/user/service"
	"blog/constants"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginService service.LoginService
}

func NewLoginController(loginService service.LoginService) *LoginController {
	return &LoginController{loginService: loginService}
}

func (c LoginController) Login(ctx *gin.Context) (interface{}, error) {
	var loginRequest request.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		return nil, constants.ResolveError
	}
	err := c.loginService.Login(ctx, loginRequest.UserName, loginRequest.Password)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
