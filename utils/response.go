package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Data interface{}

type response struct {
	Code    int
	Message string
	Data    Data
}

var successResponse = response{
	Code:    0,
	Message: "success",
}

func Response(ctx *gin.Context, httpCode int, returnCode int, message string) {
	ctx.JSON(httpCode, response{Code: returnCode, Message: message})
}

func ResponseWithError(ctx *gin.Context, error BlogError) {
	ctx.JSON(http.StatusOK, response{Code: error.Code(), Message: error.Message()})
}

func ResponseNormal(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, successResponse)
}

func ResponseWithData(ctx *gin.Context, data Data) {
	ctx.JSON(http.StatusOK, response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}
