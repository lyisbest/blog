package main

import (
	"blog/configuration"
	"blog/middleware"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	err := configuration.InitMySQL()
	if err != nil {
		panic("failed to connect mysql")
	}

	err = configuration.InitRedis()
	if err != nil {
		panic("failed to connect redis")
	}

	app := InitApp()

	r := InitGin(app)
	//err = r.Run(":8080")
	//if err != nil {
	//	fmt.Println("failed to start framework")
	//}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server listen err:%s", err)
		}
	}()

	// 在此阻塞
	<-quit

	ctx, channel := context.WithTimeout(context.Background(), 50*time.Second)

	defer channel()
	//关机
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown error")
	}

	fmt.Println("server exiting...")
}

func InitGin(application *Application) *gin.Engine {
	gin.DisableConsoleColor()

	// 记录到文件。
	f, err := os.Create("./log/gin.log")
	if err != nil {
		fmt.Printf("Log file generate failed! Error: %v\n", err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	r = SetRouters(r, application)
	return r
}

func SetRouters(r *gin.Engine, application *Application) *gin.Engine {
	route := r.Group("/user")
	{
		route.POST("/login", middleware.PostWrapper(application.Login))
	}

	manageRoute := r.Group("/admin")
	manageRoute.Use(middleware.Auth())
	{
		manageRoute.POST("/blog/create", middleware.PostWrapper(application.CreateBlog))
		manageRoute.POST("/blog/update", middleware.PostWrapper(application.UpdateBlog))
		manageRoute.POST("/blog/delete", middleware.PostWrapper(application.DeleteBlog))
	}
	generalRoute := r.Group("")
	{
		generalRoute.GET("/ping", middleware.GetWrapper(application.Ping))
		generalRoute.GET("/blog/get", middleware.GetWrapper(application.GetBlog))
		generalRoute.GET("/blog/list", middleware.GetWrapper(application.ListBlogWithPagination))
		generalRoute.POST("/update/image", application.UpdateImage)
	}

	return r
}
