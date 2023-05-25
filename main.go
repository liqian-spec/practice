package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liqian-spec/practice/bootstrap"
	"github.com/liqian-spec/practice/pkg/config"

	btsConfig "github.com/liqian-spec/practice/config"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	router := gin.New()

	bootstrap.SetupDB()

	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
