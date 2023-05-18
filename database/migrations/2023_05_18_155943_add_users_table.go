package migrations

import (
    "database/sql"
    "github.com/practice/app/models"
    "github.com/practice/pkg/migrate"
    "gorm.io/gorm"
)

func init(){

    type User struct {
        models.BaseModel

        Name    string `gorm:"type:varchar(255);not null;index"`
        Email    string `gorm:"type:varchar(255);index;default:null"`
        Phone    string `gorm:"type:varchar(255);index;default:null"`
        Password string `gorm:"type:varchar(255)"`

        models.CommonTimestampsField
    }

    up := func(migrator gorm.Migrator,DB *sql.DB) {
        migrator.AutoMigrate(&User{})
    }

    down := func(migrator gorm.Migrator,DB *sql.DB) {
        migrator.DropTable(&User{})
    }

    migrate.Add("2023_05_18_155943_add_users_table",up,down)
}