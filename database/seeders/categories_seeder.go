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

    seed.Add("SeedCategoriesTable", func(db *gorm.DB){

        categories := factories.MakeCategories(10)

        result := db.Table("categories").Create(&categories)

        if err := result.Error;err != nil{
            logger.LogIf(err)
            return
        }

        console.Success(fmt.Sprintf("Table [%v] %v rows seeded",result.Statement.Table, result.RowsAffected))
    })
}