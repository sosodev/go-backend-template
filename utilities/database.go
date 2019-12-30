package utilities

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDatabaseConnection() (*gorm.DB, error) {
	return gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable")
}
