package seeders

import "github.com/liqian-spec/practice/pkg/seed"

func Initialize() {

	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
