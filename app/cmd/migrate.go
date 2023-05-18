package cmd

import (
	"github.com/practice/database/migrations"
	"github.com/practice/pkg/migrate"
	"github.com/spf13/cobra"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unmigrated migrations",
	Run:   runUp,
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateRollback,
	)
}

func migrator() *migrate.Migrator {
	migrations.Initialize()
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}

var CmdMigrateRollback = &cobra.Command{
	Use:     "down",
	Aliases: []string{"rollback"},
	Short:   "Reverse the up command",
	Run:     runDown,
}

func runDown(cmd *cobra.Command, args []string) {
	migrator().Rollback()
}
