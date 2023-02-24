package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Data interface{}

type Response struct {
	Code    int
	Message string
	Data    Data
}

var successResponse = Response{
	Code:    0,
	Message: "success",
}

//	func Response(ctx *gin.Context, httpCode int, returnCode int, message string) {
//		ctx.JSON(httpCode, Response{Code: returnCode, Message: message})
//	}
func ResponseWithError(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, Response{Code: code, Message: message})
}

func ResponseAbnormalWithError(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusInternalServerError, Response{Code: code, Message: message})
}

func ResponseNormal(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, successResponse)
}

func ResponseWithData(ctx *gin.Context, data Data) {
	ctx.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func EndWithError(ctx *gin.Context, err error) {
	code, message := ResolveError(err)
	if code == http.StatusInternalServerError {
		ResponseAbnormalWithError(ctx, code, message)
		return
	} else {
		ResponseWithError(ctx, code, message)
		return
	}
}
