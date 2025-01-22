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
	app.Get("/", handlers.GetAllProducts(db))

	// Product routes
	app.Post("/products", handlers.CreateProduct(db))
	app.Delete("/products/:id", handlers.DeleteProduct(db))
}
