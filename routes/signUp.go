package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/manan04shah/n-queens-backend/database"
	"github.com/manan04shah/n-queens-backend/models"
)

func CreateResponseHR(hr models.HR) models.HRResponse {
	return models.HRResponse{
		CompanyName:   hr.CompanyName,
		Email:         hr.Email,
		ContactNumber: hr.ContactNumber,
	}
}

func CreateResponseEmployee(employee models.Employee) models.EmployeeResponse {
	return models.EmployeeResponse{
		FirstName:     employee.FirstName,
		LastName:      employee.LastName,
		Email:         employee.Email,
		ContactNumber: employee.ContactNumber,
		CompanyID:     employee.CompanyID,
	}
}

func SignUpHR(c *fiber.Ctx) error {
	var newHR models.HR

	if err := c.BodyParser(&newHR); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid JSON",
		})
	}

	result := database.Database.Db.Create(&newHR)

	if result.Error != nil {
		log.Printf("Error creating HR: %v\n", result.Error)
		return c.Status(500).SendString("Error creating HR")
	}

	responseHR := CreateResponseHR(newHR)
	log.Printf("HR with ID %v created successfully\n", newHR.ID)

	return c.Status(200).JSON(responseHR)
}

func SignUpEmployee(c *fiber.Ctx) error {
	var newEmployee models.Employee

	if err := c.BodyParser(&newEmployee); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid JSON",
		})
	}

	result := database.Database.Db.Create(&newEmployee)

	if result.Error != nil {
		log.Printf("Error creating Employee: %v\n", result.Error)
		return c.Status(500).SendString("Error creating Employee")
	}

	//Append the employee to the employees array of HR model having that company ID
	var hr models.HR
	hrResult := database.Database.Db.Where("id = ?", newEmployee.CompanyID).First(&hr)
	if hrResult.Error != nil {
		log.Printf("Error finding HR: %v\n", hrResult.Error)
		return c.Status(500).SendString("Error finding HR")
	}
	log.Printf("HR found: %v\n", hr)

	hr.Employees = append(hr.Employees, newEmployee)
	log.Printf("%v", hr.Employees)
	hrResult = database.Database.Db.Save(&hr)
	if hrResult.Error != nil {
		log.Printf("Error saving HR: %v\n", hrResult.Error)
		return c.Status(500).SendString("Error saving HR")
	}

	newEmployeeResponse := CreateResponseEmployee(newEmployee)
	log.Printf("Employee with ID %v created successfully\n", newEmployee.ID)

	return c.Status(200).JSON(newEmployeeResponse)
}
