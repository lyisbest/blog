package middleware

import (
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func Wrapper(f func(ctx *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := f(ctx)
		if err == nil {
			if data == nil {
				utils.ResponseNormal(ctx)
				return
			}
			utils.ResponseWithData(ctx, data)
			return
		} else {
			utils.EndWithError(ctx, err)
		}
	}
}
