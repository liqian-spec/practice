package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/liqian-spec/practice/pkg/config"
	"github.com/liqian-spec/practice/pkg/logger"
	"github.com/liqian-spec/practice/pkg/redis"
	limiterlib "github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

func GetKeyIP(c *gin.Context) string {
	return c.ClientIP()
}

func GetKeyRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.FullPath()) + c.ClientIP()
}

func CheckRate(c *gin.Context, key string, formatted string) (limiterlib.Context, error) {

	var context limiterlib.Context
	rate, err := limiterlib.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	store, err := sredis.NewStoreWithOptions(redis.Redis.Client, limiterlib.StoreOptions{
		Prefix: config.GetString("app.name") + ":limiter",
	})
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	limiterObj := limiterlib.New(store, rate)

	if c.GetBool("limiter-once") {
		return limiterObj.Peek(c, key)
	} else {
		c.Set("limiter-once", true)

		return limiterObj.Get(c, key)
	}
}

func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName
}
