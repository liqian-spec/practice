package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/practice/pkg/jwt"
	"github.com/practice/pkg/response"
)

func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(c.GetHeader("Authorization")) > 0 {

			_, err := jwt.NewJWT().ParserToken(c)
			if err == nil {
				response.Unauthorized(c, "请使用游客身份访问")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
