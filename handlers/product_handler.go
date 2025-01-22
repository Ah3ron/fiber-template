package handlers

import (
	"your-project/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		product := new(models.Product)
		if err := c.BodyParser(product); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		db.Create(product)

		// Return the updated product list
		var products []models.Product
		db.Find(&products)
		return c.Render("product_list", fiber.Map{
			"Products": products,
		})
	}
}

func GetAllProducts(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var products []models.Product
		db.Find(&products)
		return c.Render("product_list", fiber.Map{
			"Products": products,
		})
	}
}

func DeleteProduct(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		db.Delete(&models.Product{}, id)

		// Return the updated product list
		var products []models.Product
		db.Find(&products)
		return c.Render("product_list", fiber.Map{
			"Products": products,
		})
	}
}
