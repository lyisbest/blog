package router

import (
	"blog/apps/login/controller"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) *gin.Engine {
	c := controller.LoginController{}
	route := r.Group("/login")
	{
		route.POST("", c.Login)
	}

	manageRoute := r.Group("/admin")
	manageRoute.Use(middleware.Auth())
	{
		manageRoute.POST("")
	}

	return r
}
