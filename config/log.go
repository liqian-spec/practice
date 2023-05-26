package config

import "github.com/liqian-spec/practice/pkg/config"

func init() {
	config.Add("log", func() map[string]interface{} {
		return map[string]interface{}{

			"level":      config.Env("LOG_LEVEL", "debug"),
			"type":       config.Env("LOG_TYPE", "single"),
			"filename":   config.Env("LOG_NAME", "storage/logs/logs.log"),
			"max_size":   config.Env("LOG_MAX_SIZE", 64),
			"max_backup": config.Env("LOG_MAX_BACKUP", 5),
			"max_age":    config.Env("LOG_MAX_AGE", 30),
			"compress":   config.Env("LOG_COMPRESS", true),
		}
	})
}
