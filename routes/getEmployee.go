package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manan04shah/n-queens-backend/database"
	"github.com/manan04shah/n-queens-backend/models"
)

func GetEmployee(c *fiber.Ctx) error {
	email := c.Query("email")
	var employee models.Employee
	result := database.Database.Db.Where("email = ?", email).First(&employee)
	if result.RowsAffected == 0 {
		return c.Status(404).SendString("Employee not found")
	}
	if result.Error != nil {
		return c.Status(500).SendString("Error finding Employee")
	}
	return c.Status(200).JSON(employee)
}
