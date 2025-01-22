package main

import (
	"log"
	"your-project/config"
	"your-project/database"
	"your-project/models"
	"your-project/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the database
	db := database.InitDB(cfg.DBURL)
	db.AutoMigrate(&models.User{})

	// Initialize the Fiber app with HTML templates
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Set up routes
	routes.SetupRoutes(app, db)

	// Start the server
	log.Fatal(app.Listen(":" + cfg.Port))
}
