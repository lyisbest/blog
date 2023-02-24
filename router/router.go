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
		route.POST("/login", middleware.Wrapper(c.Login))
	}

	manageRoute := r.Group("/admin")
	manageRoute.Use(middleware.Auth())
	{
		manageRoute.POST("/blog/create", middleware.Wrapper(blogcontroller.CreateBlog))
		manageRoute.POST("/blog/update", middleware.Wrapper(blogcontroller.UpdateBlog))
		manageRoute.POST("/blog/delete", middleware.Wrapper(blogcontroller.DeleteBlog))
	}
	generalRoute := r.Group("")
	{
		generalRoute.GET("/blog/get", middleware.Wrapper(blogcontroller.GetBlog))
		generalRoute.GET("/blog/list", middleware.Wrapper(blogcontroller.ListBlogWithPagination))
	}

	return r
}
