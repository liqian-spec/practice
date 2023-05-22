package migrations

import (
	"database/sql"
	"github.com/practice/app/models"
	"github.com/practice/pkg/migrate"
	"gorm.io/gorm"
)

func init() {

	type User struct {
		models.BaseModel

		City         string `gorm:"type:varchar(10);"`
		Introduction string `gorm:"type:varchar(255);"`
		Avatar       string `gorm:"type:varchar(255);default:null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{})
	}

	migrate.Add("2023_05_22_142455_add_fields_to_user", up, down)
}
