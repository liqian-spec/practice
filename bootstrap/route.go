package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/liqian-spec/practice/app/http/middlewares"
	"github.com/liqian-spec/practice/routes"
	"net/http"
	"strings"
)

func SetupRoute(router *gin.Engine) {

	registerGlobalMiddleware(router)

	routes.RegisterAPIRoutes(router)

	setup404Handler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.ForceUA(),
	)
}

func setup404Handler(router *gin.Engine) {

	router.NoRoute(func(c *gin.Context) {

		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确.",
			})
		}
	})
}
