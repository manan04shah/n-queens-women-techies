package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/manan04shah/n-queens-backend/database"
	"github.com/manan04shah/n-queens-backend/models"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHR(c *fiber.Ctx) error {
	var loginRequest LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid JSON",
		})
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   "Email and password are required",
		})
	}

	//Query the database for user with that email ID
	//If user is found, check if the password matches
	//If password matches, return the user details
	//If password does not match, return an error
	//If user is not found, return an error
	var user models.HR
	result := database.Database.Db.Where("email = ?", loginRequest.Email).First(&user)
	if result.RowsAffected == 0 {
		return c.Status(404).SendString("HR not found")
	}
	if result.Error != nil {
		return c.Status(500).SendString("Error finding HR")
	}

	if user.Password != loginRequest.Password {
		return c.Status(401).SendString("Incorrect password")
	}

	log.Printf("HR with email %v logged in successfully\n", loginRequest.Email)
	return c.Status(200).JSON(user)
}

func LoginEmployee(c *fiber.Ctx) error {
	var loginRequest LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid JSON",
		})
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   "Email and password are required",
		})
	}

	//Query the database for user with that email ID
	//If user is found, check if the password matches
	//If password matches, return the user details
	//If password does not match, return an error
	//If user is not found, return an error
	var user models.Employee
	result := database.Database.Db.Where("email = ?", loginRequest.Email).First(&user)
	if result.RowsAffected == 0 {
		return c.Status(404).SendString("Employee not found")
	}
	if result.Error != nil {
		return c.Status(500).SendString("Error finding Employee")
	}

	if user.Password != loginRequest.Password {
		return c.Status(401).SendString("Incorrect password")
	}

	log.Printf("Employee with email %v logged in successfully\n", loginRequest.Email)
	return c.Status(200).JSON(user)
}
