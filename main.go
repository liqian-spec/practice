package main

import (
	"fmt"
	"github.com/practice/app/cmd"
	"github.com/practice/app/cmd/make"
	"github.com/practice/bootstrap"
	btsConfig "github.com/practice/config"
	"github.com/practice/pkg/config"
	"github.com/practice/pkg/console"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	var rootCmd = &cobra.Command{
		Use:   "Practice",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command,you can use "-h" flag to see all subcommands`,

		PersistentPreRun: func(command *cobra.Command, args []string) {

			config.InitConfig(cmd.Env)
			// 初始化 Logger
			bootstrap.SetupLogger()
			// 初始化 DB
			bootstrap.SetupDB()
			// 初始化 Redis
			bootstrap.SetupRedis()
			// 初始化缓存
			bootstrap.SetupCache()
		},
	}

	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make.CmdMake,
		cmd.CmdMigrate,
		cmd.CmdDBSeed,
		cmd.CmdCache,
	)

	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	cmd.RegisterGlobalFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v:%s", os.Args, err.Error()))
	}

}
