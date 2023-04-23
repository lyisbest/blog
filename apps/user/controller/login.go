package controller

import (
	"blog/apps/user/request"
	"blog/apps/user/service"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginService *service.LoginService
}

func NewLoginController(loginService *service.LoginService) *LoginController {
	return &LoginController{loginService: loginService}
}

func (c LoginController) Login(ctx *gin.Context, loginRequest request.LoginRequest) (interface{}, error) {
	err := c.loginService.Login(ctx, loginRequest.UserName, loginRequest.Password)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
