package middleware

import (
	"blog/apps/user/constant"
	"blog/apps/user/repository"
	"blog/constants"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("user_cookie")
		if err != nil {
			utils.ResponseWithError(ctx, constant.CookieResolveError)
			ctx.Abort()
			return
		}
		rowNum, err := repository.NewUserRepository().GetUserByUserName(ctx, cookie.Value)
		if err != nil {
			utils.ResponseWithError(ctx, constants.QueryError)
			ctx.Abort()
			return
		}
		if rowNum == 0 {
			utils.ResponseWithError(ctx, constant.CookieError)
			ctx.Abort()
			return
		}
		ctx.SetCookie(cookie.Name, cookie.Value, 1000, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		ctx.Set("username", cookie.Value)
		ctx.Next()
	}
}
