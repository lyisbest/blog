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

		if token == "" {
			utils.EndWithError(ctx, constant.TokenError)
			ctx.Abort()
			return
		}

		//验证token
		isValidated := utils.ValidateToken(token)
		if !isValidated {
			utils.EndWithError(ctx, constant.TokenError)
			ctx.Abort()
			return
		}

		//先去查询该token是否已被使用过，如果使用过。表明该token有可能被窃取。
		err := utils.GetTokenUsed(token)
		if err != nil {
			utils.EndWithError(ctx, err)
			ctx.Abort()
			return
		}

		//解析token，用于识别用户
		_, claim, err := utils.ParseToken(token)
		if err != nil {
			utils.EndWithError(ctx, err)
			ctx.Abort()
			return
		}

		//识别用户
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

		//生成刷新后的token
		refreshToken, err := utils.RefreshToken(token)
		if err != nil {
			utils.EndWithError(ctx, err)
			ctx.Abort()
			return
		}

		//标记已经被使用过的token
		err = utils.MarkTokenUsed(token)
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
