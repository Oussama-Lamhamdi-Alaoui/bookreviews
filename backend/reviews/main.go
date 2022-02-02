package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Review struct {
	Id     uint   `json:"id"`
	BookId uint   `json:"book_id"`
	Text   string `json:"text"`
}

func main() {
	// Connect to SQLite Database
	db, err := gorm.Open(sqlite.Open("db_reviews.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	// Migrate Book model
	db.AutoMigrate(Review{})

	// Init & Run Fiber app
	app := fiber.New()

	// Chrome workaround
	app.Use(cors.New())

	app.Get("/api/books/:id/reviews", func(c *fiber.Ctx) error {
		var reviews []Review

		db.Find(&reviews, "book_id = ?", c.Params("id"))

		return c.JSON(reviews)
	})

	app.Post("/api/reviews", func(c *fiber.Ctx) error {
		var review Review

		if err := c.BodyParser(&review); err != nil {
			return err
		}

		db.Create(&review)

		return c.JSON(review)
	})

	app.Listen(":3010")
}
