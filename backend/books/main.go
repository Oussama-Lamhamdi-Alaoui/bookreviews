package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Model
type Book struct {
	Id      uint     `json:"id"`
	Title   string   `json:"title"`
	Genre   string   `json:"genre"`
	Reviews []Review `json:"reviews" gorm:"-" default:"[]"`
}

type Review struct {
	Id     uint   `json:"id"`
	BookId uint   `json:"book_id"`
	Text   string `json:"text"`
}

func main() {
	// Connect to SQLite Database
	db, err := gorm.Open(sqlite.Open("db_books.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	// Migrate Book model
	db.AutoMigrate(Book{})

	// Init & Run Fiber app
	app := fiber.New()

	// Chrome workaround
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Bonjour Madame 👋!")
	})

	// Endpoint to handle GET requests
	app.Get("/api/books", func(c *fiber.Ctx) error {
		var books []Book

		db.Find(&books)

		for i, book := range books {
			response, err := http.Get(fmt.Sprintf("http://localhost:3010/api/books/%d/reviews", book.Id))

			if err != nil {
				return err
			}

			var reviews []Review

			json.NewDecoder(response.Body).Decode(&reviews)

			books[i].Reviews = reviews
		}

		return c.JSON(books)
	})

	// Endpoint to handle POST requests
	app.Post("/api/books", func(c *fiber.Ctx) error {
		var book Book

		if err := c.BodyParser(&book); err != nil {
			return err
		}

		db.Create(&book)

		return c.JSON(book)
	})

	app.Listen(":3005")
}
