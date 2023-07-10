package bootstrap

import (
	"fmt"

	"github.com/liqian-spec/practice/pkg/config"
	"github.com/liqian-spec/practice/pkg/redis"
)

func SetupRedis() {

	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
