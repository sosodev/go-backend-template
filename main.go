package main

import (
	"log"
	"net/http"

	"github.com/go-backend-template/models"

	"github.com/go-backend-template/utilities"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func migrateDatabase() {
	db, err := utilities.GetDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Example{})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

    migrateDatabase()

	e.POST("/example", func(c echo.Context) error {
		example := &models.Example{}
		if err := c.Bind(example); err != nil {
			return err
		}

		db, err := utilities.GetDatabaseConnection() 
		if err != nil {
			return err
		}
        defer db.Close()

        db.Create(example)
        return c.JSON(http.StatusCreated, example)
	})

	e.GET("/example", func(c echo.Context) error {
        examples := []models.Example{}
        db, err := utilities.GetDatabaseConnection()
        if err != nil {
            return err
        }
        defer db.Close()

        db.Limit(5).Find(&examples)

        return c.JSON(http.StatusOK, examples)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

