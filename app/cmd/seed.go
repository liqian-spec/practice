package cmd

import (
	"github.com/practice/database/seeders"
	"github.com/practice/pkg/console"
	"github.com/practice/pkg/seed"
	"github.com/spf13/cobra"
)

var CmdDBSeed = &cobra.Command{
	Use:   "seed",
	Short: "Insert fake data to the database",
	Run:   runSeeders,
	Args:  cobra.MaximumNArgs(1),
}

func runSeeders(cmd *cobra.Command, args []string) {
	seeders.Initialize()
	if len(args) > 0 {
		name := args[0]
		seeder := seed.GetSeeder(name)
		if len(seeder.Name) > 0 {
			seed.RunSeeder(name)
		} else {
			console.Error("Seeder not found: " + name)
		}
	} else {
		seed.RunAll()
		console.Success("Done seeding.")
	}
}
