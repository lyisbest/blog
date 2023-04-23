//go:build wireinject
// +build wireinject

package main

import (
	"blog/apps/blog/controller"
	"blog/apps/blog/repository"
	"blog/apps/blog/service"
	controller3 "blog/apps/sys/controller"
	controller2 "blog/apps/user/controller"
	repository2 "blog/apps/user/repository"
	service2 "blog/apps/user/service"
	"github.com/google/wire"
)

var set = wire.NewSet(controller.NewBlogController, service.NewBlogService, repository.NewBlogRepository)
var userSet = wire.NewSet(controller2.NewLoginController, service2.NewLoginService, repository2.NewUserRepository)
var sysSet = wire.NewSet(controller3.NewSysController)
var appSet = wire.NewSet(NewApplication)

func InitApp() *Application {
	wire.Build(set, sysSet, userSet, appSet)
	return &Application{}
}

type Application struct {
	*controller.BlogController
	*controller2.LoginController
	*controller3.SysController
}

func NewApplication(blogController *controller.BlogController, loginController *controller2.LoginController, sysController *controller3.SysController) *Application {
	return &Application{
		BlogController:  blogController,
		LoginController: loginController,
		SysController:   sysController,
	}
}
