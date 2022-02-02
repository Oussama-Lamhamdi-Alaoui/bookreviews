package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Model
type Publisher struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Established  uint   `json:"established"`
	Headquarters string `json:"headquarters"`
	Ceo          string `json:"ceo"`
}

func main() {
	// Connect to SQLite Database
	db, err := gorm.Open(sqlite.Open("db_publishers.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	// Migrate Publisher model
	db.AutoMigrate(Publisher{})

	app := fiber.New()

	// Chrome workaround
	app.Use(cors.New())

	// Endpoint to handle GET requests
	app.Get("/api/publishers", func(c *fiber.Ctx) error {
		var publishers []Publisher

		db.Find(&publishers)

		return c.JSON(publishers)
	})

	// Endpoint to handle POST requests
	app.Post("/api/publishers", func(c *fiber.Ctx) error {
		var publisher Publisher

		if err := c.BodyParser(&publisher); err != nil {
			return err
		}

		db.Create(&publisher)

		return c.JSON(publisher)
	})

	app.Listen(":3020")
}
