package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/liqian-spec/practice/pkg/response"
)

func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(c.Request.Header["User-Agent"]) == 0 {
			response.BadRequest(c, errors.New("User-Agent 标头未找到"), "请求必须附带 User-Agent 标头")
			return
		}

		c.Next()
	}
}
