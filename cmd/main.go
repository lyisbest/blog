package main

import (
	"blog/configuration"
	"blog/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func main() {

	err := configuration.InitMySQL()
	if err != nil {
		panic("failed to connect database")
	}

	r := InitGin()
	err = r.Run(":8080")
	if err != nil {
		fmt.Println("failed to start framework")
	}
}

func InitGin() *gin.Engine {
	gin.DisableConsoleColor()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	r = router.SetRouters(r)
	return r
}
