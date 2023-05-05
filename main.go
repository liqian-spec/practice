package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/practice/bootstrap"
	btsConfig "github.com/practice/config"
	"github.com/practice/pkg/config"
	"github.com/practice/pkg/sms"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如--env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// 初始化 Logger
	bootstrap.SetupLogger()
	// 初始化 DB
	bootstrap.SetupDB()
	// 初始化 Redis
	bootstrap.SetupRedis()

	// new 一个 Gin Engine 实例
	router := gin.New()
	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	sms.NewSMS().Send("15931012793", sms.Message{
		Template: config.GetString("sms.aliyun.template_code"),
		Data:     map[string]string{"code": "123456"},
	})

	// 运行服务
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
