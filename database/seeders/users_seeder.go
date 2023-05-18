package seeders

import (
	"fmt"
	"github.com/practice/database/factories"
	"github.com/practice/pkg/console"
	"github.com/practice/pkg/logger"
	"github.com/practice/pkg/seed"
	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedUsersTable", func(db *gorm.DB) {

		users := factories.Makeusers(10)

		result := db.Table("users").Create(&users)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
