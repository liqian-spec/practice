package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/liqian-spec/practice/bootstrap"
	"github.com/liqian-spec/practice/pkg/config"
	"github.com/liqian-spec/practice/pkg/console"
	"github.com/liqian-spec/practice/pkg/logger"
	"github.com/spf13/cobra"
)

var CmdServer = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server,error:" + err.Error())
	}
}
