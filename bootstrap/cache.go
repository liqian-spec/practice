package bootstrap

import (
	"fmt"
	"github.com/practice/pkg/cache"
	"github.com/practice/pkg/config"
)

func SetupCache() {

	rds := cache.NewRedisStore(fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
	)

	cache.InitWithCacheStore(rds)
}
