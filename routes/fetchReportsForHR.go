package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/manan04shah/n-queens-backend/database"
	"github.com/manan04shah/n-queens-backend/models"
)

func FetchReportsForHR(c *fiber.Ctx) error {
	companyID := c.Query("company_id")

	var reports []models.Report
	result := database.Database.Db.Where("hr_id = ?", companyID).Find(&reports)

	if result.Error != nil {
		return c.Status(500).SendString("Error fetching reports")
	}

	if result.RowsAffected == 0 {
		return c.Status(404).SendString("No reports found for the specified HR")
	}

	log.Printf("Fetched %v reports for HR with ID %v\n", len(reports), companyID)
	return c.Status(200).JSON(reports)
}
