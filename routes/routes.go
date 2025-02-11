package routes

import (
	"your-project/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Serve static files
	app.Static("/static", "./static")

	// Render the index page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	// Product routes
	app.Get("/products/list", handlers.GetAllProducts(db))
	app.Post("/products", handlers.CreateProduct(db))
	app.Delete("/products/:id", handlers.DeleteProduct(db))
}
