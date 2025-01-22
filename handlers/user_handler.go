package handlers

import (
	"your-project/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(models.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		db.Create(user)
		return c.JSON(user)
	}
}

func GetAllUsers(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var users []models.User
		db.Find(&users)
		return c.JSON(users)
	}
}

func GetUserByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user models.User
		db.First(&user, id)
		return c.JSON(user)
	}
}

func UpdateUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user models.User
		db.First(&user, id)
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		db.Save(&user)
		return c.JSON(user)
	}
}

func DeleteUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		db.Delete(&models.User{}, id)
		return c.SendString("User deleted")
	}
}
