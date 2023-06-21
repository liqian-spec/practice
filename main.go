package main

import (
	"fmt"
	"github.com/liqian-spec/practice/app/cmd"
	"github.com/liqian-spec/practice/bootstrap"
	btsConfig "github.com/liqian-spec/practice/config"
	"github.com/liqian-spec/practice/pkg/config"
	"github.com/liqian-spec/practice/pkg/console"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	var rootCmd = &cobra.Command{
		Use:   "Practice",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command,you can use "-h" flag to see all subcommands`,

		PersistentPreRun: func(command *cobra.Command, args []string) {

			config.InitConfig(cmd.Env)

			bootstrap.SetupLogger()

			bootstrap.SetupDB()

			bootstrap.SetupRedis()
		},
	}

	rootCmd.AddCommand(
		cmd.CmdServer,
		cmd.CmdKey,
		cmd.CmdPlay,
	)

	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServer)

	cmd.RegisterGlobalFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v:%s", os.Args, err.Error()))
	}
}
