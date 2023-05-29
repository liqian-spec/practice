package config

import "github.com/liqian-spec/practice/pkg/config"

func init() {
	config.Add("mail", func() map[string]interface{} {
		return map[string]interface{}{

			"smtp": map[string]interface{}{
				"host":     config.Env("MAIL_HOST", "localhost"),
				"port":     config.Env("MAIL_PORT", "1025"),
				"username": config.Env("MAIL_USERNAME", ""),
				"password": config.Env("MAIL_PASSWORD", ""),
			},

			"from": map[string]interface{}{
				"address": config.Env("MAIL_FROM_ADDRESS", "practice@example.com"),
				"name":    config.Env("MAIL_FROM_NAME", "Practice"),
			},
		}
	})
}
