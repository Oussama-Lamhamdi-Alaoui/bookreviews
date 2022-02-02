package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Model
type Library struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Established uint   `json:"established"`
	Location    string `json:"location"`
	Director    string `json:"director"`
}

func main() {
	// Connect to SQLite Database
	db, err := gorm.Open(sqlite.Open("db_libraries.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	// Migrate Library model
	db.AutoMigrate(Library{})

	app := fiber.New()

	// Chrome workaround
	app.Use(cors.New())

	// Endpoint to handle GET requests
	app.Get("/api/libraries", func(c *fiber.Ctx) error {
		var libraries []Library

		db.Find(&libraries)

		return c.JSON(libraries)
	})

	// Endpoint to handle POST requests
	app.Post("/api/libraries", func(c *fiber.Ctx) error {
		var library Library

		if err := c.BodyParser(&library); err != nil {
			return err
		}

		db.Create(&library)

		return c.JSON(library)
	})

	app.Listen(":3015")
}
