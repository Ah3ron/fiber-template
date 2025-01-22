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
		return c.JSON(product)
	}
}

func GetAllProducts(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var products []models.Product
		db.Find(&products)
		return c.JSON(products)
	}
}

func GetProductByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var product models.Product
		db.First(&product, id)
		return c.JSON(product)
	}
}

func UpdateProduct(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var product models.Product
		db.First(&product, id)
		if err := c.BodyParser(&product); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		db.Save(&product)
		return c.JSON(product)
	}
}

func DeleteProduct(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		db.Delete(&models.Product{}, id)
		return c.SendString("Product deleted")
	}
}
