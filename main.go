package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.New()
	// 处理404请求
	r.NoRoute(func(c *gin.Context) {
		//获取标头信息的Accept信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			//如果是html的话
			c.String(http.StatusNotFound, "页面返回404")
		} else {
			// 默认返回JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认URL和请求方法是否正确。",
			})
		}
	})

	//运行服务
	r.Run(":8000")
}
