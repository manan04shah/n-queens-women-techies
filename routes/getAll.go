package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manan04shah/n-queens-backend/database"
	"github.com/manan04shah/n-queens-backend/models"
)

func GetAllEmployees(c *fiber.Ctx) error {
	var employees []models.Employee
	result := database.Database.Db.Find(&employees)
	if result.Error != nil {
		return c.Status(500).SendString("Error finding employees")
	}
	return c.Status(200).JSON(employees)
}

func GetAllHR(c *fiber.Ctx) error {
	var hr []models.HR
	result := database.Database.Db.Find(&hr)
	if result.Error != nil {
		return c.Status(500).SendString("Error finding HR")
	}
	return c.Status(200).JSON(hr)
}

func GetAllReports(c *fiber.Ctx) error {
	var reports []models.Report
	result := database.Database.Db.Find(&reports)
	if result.Error != nil {
		return c.Status(500).SendString("Error finding reports")
	}
	return c.Status(200).JSON(reports)
}
