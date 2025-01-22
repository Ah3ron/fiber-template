package routes

import (
	"your-project/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Serve static files
	app.Static("/static", "./static")

	// HTML routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	// HTMX routes
	app.Get("/api/data", func(c *fiber.Ctx) error {
		return c.SendString("<p>Data loaded via HTMX!</p>")
	})

	// User routes
	userRoutes := app.Group("/users")
	userRoutes.Post("/", handlers.CreateUser(db))
	userRoutes.Get("/", handlers.GetAllUsers(db))
	userRoutes.Get("/:id", handlers.GetUserByID(db))
	userRoutes.Put("/:id", handlers.UpdateUser(db))
	userRoutes.Delete("/:id", handlers.DeleteUser(db))
}
