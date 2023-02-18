package main

import (
	"blog/configuration"
	"blog/router"
	"fmt"
	"github.com/gin-gonic/gin"
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
	r := gin.Default()
	r = router.SetRouters(r)
	return r
}
