package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Model
type User struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Joined   uint   `json:"joined"`
}

func main() {
	// Connect to SQLite Database
	db, err := gorm.Open(sqlite.Open("db_users.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	// Migrate Library model
	db.AutoMigrate(User{})

	app := fiber.New()

	// Chrome workaround
	app.Use(cors.New())

	// Endpoint to handle GET requests
	app.Get("/api/users", func(c *fiber.Ctx) error {
		var users []User

		db.Find(&users)

		return c.JSON(users)
	})

	// Endpoint to handle POST requests
	app.Post("/api/users", func(c *fiber.Ctx) error {
		var user User

		if err := c.BodyParser(&user); err != nil {
			return err
		}

		db.Create(&user)

		return c.JSON(user)
	})

	app.Listen(":3030")
}
