package middleware

import (
	"gin-z/app/common/response"
	"gin-z/app/services"
	"gin-z/global"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.TokenFail(ctx)
			ctx.Abort()
			return
		}
		tokenStr = tokenStr[len(services.TokenType)+1:]

		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		if err != nil || services.JwtService.IsInBlacklist(tokenStr) {
			response.TokenFail(ctx)
			ctx.Abort()
			return
		}

		claims := token.Claims.(*services.CustomClaims)
		// Token 发布者校验
		if claims.Issuer != GuardName {
			response.TokenFail(ctx)
			ctx.Abort()
			return
		}

		// token 续签
		if claims.ExpiresAt-time.Now().Unix() < global.App.Config.Jwt.RefreshGracePeriod {
			lock := global.Lock("refresh_token_lock", global.App.Config.Jwt.JwtBlacklistGracePeriod)
			if lock.Get() {
				err, user := services.JwtService.GetUserInfo(GuardName, claims.Id)
				if err != nil {
					global.App.Log.Error(err.Error())
					lock.Release()
				} else {
					tokenData, _, _ := services.JwtService.CreateToken(GuardName, user)
					ctx.Header("new-token", tokenData.AccessToken)
					ctx.Header("new-expires-in", strconv.Itoa(tokenData.ExpiresIn))
					_ = services.JwtService.JoinBlackList(token)
				}
			}
		}

		ctx.Set("token", token)
		ctx.Set("id", claims.Id)
	}
}
