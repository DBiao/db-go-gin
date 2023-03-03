package middleware

import (
	"db-go-gin/internal/app/dto/response"
	"db-go-gin/internal/global/statuscode"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/multi"
)

// JwtAuth token验证中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		claims, err := multi.AuthDriver.GetMultiClaims(token)
		if err != nil {
			response.Response(c, response.NewResponseMsg(statuscode.SystemUserAuthError, statuscode.GetText(statuscode.SystemUserAuthError), ""))
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
