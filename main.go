package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/manan04shah/n-queens-backend/database"
	"github.com/manan04shah/n-queens-backend/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the nQueens Backend")
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	// CORS middleware configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3001", // Allow all origins (change to specific origins if needed)
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Get("/", welcome)
	app.Get("/news", routes.ReturnNews)
	//Sign Up Routes
	app.Post("/signup/hr", routes.SignUpHR)
	app.Post("/signup/employee", routes.SignUpEmployee)
	//Login Routes
	app.Post("/login/hr", routes.LoginHR)
	app.Post("/login/employee", routes.LoginEmployee)
	//Report Routes
	app.Post("/report/create", routes.CreateReport)
	//Get Routes
	app.Get("/get/employee", routes.GetEmployee)
	app.Get("/get/hr", routes.GetHR)
	app.Get("/get/all/employees", routes.GetAllEmployees)
	app.Get("/get/all/hr", routes.GetAllHR)
	app.Get("/get/all/reports", routes.GetAllReports)

	log.Fatal(app.Listen(":3000"))
}
