package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Model
type Author struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Birth       uint   `json:"birth"`
	Nationality string `json:"nationality"`
}

func main() {
	// Connect to SQLite Database
	db, err := gorm.Open(sqlite.Open("db_authors.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	// Migrate Author model
	db.AutoMigrate(Author{})

	app := fiber.New()

	// Chrome workaround
	app.Use(cors.New())

	// Endpoint to handle GET requests
	app.Get("/api/authors", func(c *fiber.Ctx) error {
		var authors []Author

		db.Find(&authors)

		return c.JSON(authors)
	})

	// Endpoint to handle POST requests
	app.Post("/api/authors", func(c *fiber.Ctx) error {
		var author Author

		if err := c.BodyParser(&author); err != nil {
			return err
		}

		db.Create(&author)

		return c.JSON(author)
	})

	app.Listen(":3025")
}
