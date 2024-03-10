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
	hrResult := database.Database.Db.First(&hr, newReport.HRID)
	if hrResult.Error != nil {
		return c.Status(500).SendString("Error finding HR: " + hrResult.Error.Error())
	}

	if hrResult.RowsAffected == 0 {
		return c.Status(404).SendString("HR not found")
	}

	// Append the new report to the Reports array of the HR
	hr.Reports = append(hr.Reports, newReport)

	// Save the updated HR back to the database
	updateResult := database.Database.Db.Save(&hr)
	if updateResult.Error != nil {
		return c.Status(500).SendString("Error updating HR: " + updateResult.Error.Error())
	}

	responseReport := CreateResponseReport(newReport)
	log.Printf("Report with ID %v created successfully\n", newReport.ID)

	return c.Status(200).JSON(responseReport)
}
