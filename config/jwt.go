package config

import "github.com/practice/pkg/config"

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{

			// 过期时间，单位分钟
			"expire_time": config.Env("JWT_EXPIRE_TIME", 120),

			"max_refresh_time": config.Env("JWT_MAX_REFRESH_TIME", 86400),

			"debug_expire_time": 86400,
		}
	})
}