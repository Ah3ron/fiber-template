package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	// Initialize the database
	db := initDB()

	// Auto migrate the schema
	db.AutoMigrate(&User{})

	// Initialize the Fiber app with HTML templates
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static files
	app.Static("/static", "./static")

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	// Htmx routes
	app.Get("/api/data", func(c *fiber.Ctx) error {
		return c.SendString("<p>Data loaded via HTMX!</p>")
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}

func initDB() *gorm.DB {
	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}
