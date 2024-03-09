package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/manan04shah/n-queens-backend/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the nQueens Backend")
}

func main() {
	app := fiber.New()

	app.Get("/", welcome)
	app.Get("/news", routes.ReturnNews)

	log.Fatal(app.Listen(":3000"))
}
