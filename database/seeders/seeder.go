package seeders

import "github.com/practice/pkg/seed"

func Initialize() {

	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
