package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manan04shah/n-queens-backend/database"
	"github.com/manan04shah/n-queens-backend/models"
)

func GetHR(c *fiber.Ctx) error {
	email := c.Query("email")
	var hr models.HR
	result := database.Database.Db.Where("email = ?", email).First(&hr)
	if result.RowsAffected == 0 {
		return c.Status(404).SendString("HR not found")
	}
	if result.Error != nil {
		return c.Status(500).SendString("Error finding HR")
	}
	return c.Status(200).JSON(hr)
}
