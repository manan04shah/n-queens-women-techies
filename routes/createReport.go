package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/manan04shah/n-queens-backend/database"
	"github.com/manan04shah/n-queens-backend/models"
)

func CreateResponseReport(report models.Report) models.ReportResponse {
	return models.ReportResponse{
		ReportAgainst:     report.ReportAgainst,
		ReportDescription: report.ReportDescription,
		ProofURL:          report.ProofURL,
		HRID:              report.HRID,
	}
}

func CreateReport(c *fiber.Ctx) error {
	var newReport models.Report

	if err := c.BodyParser(&newReport); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid JSON",
		})
	}

	result := database.Database.Db.Create(&newReport)

	if result.Error != nil {
		return c.Status(500).SendString("Error creating report")
	}

	// Retrieve the HR from the database using HRID
	var hr models.HR
	hrResult := database.Database.Db.Where("id = ?", newReport.HRID).First(&hr)

	if hrResult.Error != nil {
		return c.Status(500).SendString("Error retrieving HR")
	}
	log.Printf("HR with ID %v retrieved successfully\n", newReport.HRID)

	hr.Reports = append(hr.Reports, newReport)
	database.Database.Db.Save(&hr).Debug()
	database.Database.Db.Model(&hr).Association("Reports").Append(&newReport)

	responseReport := CreateResponseReport(newReport)
	log.Printf("Report with ID %v created successfully\n", newReport.ID)

	return c.Status(200).JSON(responseReport)
}
