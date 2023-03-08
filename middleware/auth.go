package middleware

import (
	"blog/apps/user/constant"
	"blog/apps/user/repository"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		isValidated := utils.ValidateToken(token)
		if !isValidated {
			utils.EndWithError(ctx, constant.TokenError)
			ctx.Abort()
			return
		}

		_, claim, err := utils.ParseToken(token)
		if err != nil {
			utils.EndWithError(ctx, err)
			ctx.Abort()
			return
		}

		rowNum, err := repository.NewUserRepository().GetUserByUserName(ctx, claim.Name)
		if err != nil {
			utils.EndWithError(ctx, err)
			ctx.Abort()
			return
		}

		if rowNum == 0 {
			utils.EndWithError(ctx, constant.TokenError)
			ctx.Abort()
			return
		}

		refreshToken, err := utils.RefreshToken(token)
		if err != nil {
			utils.EndWithError(ctx, err)
			ctx.Abort()
			return
		}

		ctx.Header("token", refreshToken)
		ctx.Set("username", claim.Name)
		ctx.Next()
	}
}
