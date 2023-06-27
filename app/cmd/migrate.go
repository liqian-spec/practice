package cmd

import (
	"github.com/liqian-spec/practice/database/migrations"
	"github.com/liqian-spec/practice/pkg/migrate"
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
	)
}

func migrator() *migrate.Migrator {
	migrations.Initialize()
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}
