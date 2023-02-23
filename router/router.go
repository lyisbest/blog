package router

import (
	controller2 "blog/apps/blog/controller"
	"blog/apps/user/controller"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) *gin.Engine {
	c := controller.LoginController{}
	blogcontroller := controller2.BlogController{}
	route := r.Group("/user")
	{
		route.POST("/login", c.Login)
	}

	manageRoute := r.Group("/admin")
	manageRoute.Use(middleware.Auth())
	{
		manageRoute.POST("/blog/create", blogcontroller.CreateBlog)
		manageRoute.POST("/blog/update", blogcontroller.UpdateBlog)
		manageRoute.POST("/blog/delete", blogcontroller.DeleteBlog)
	}
	generalRoute := r.Group("")
	{
		generalRoute.GET("/blog/get", blogcontroller.GetBlog)
		generalRoute.GET("/blog/list", blogcontroller.ListBlogWithPagination)
	}

	return r
}
