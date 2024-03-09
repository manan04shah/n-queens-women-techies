package routes

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func ReturnNews(c *fiber.Ctx) error {
	// Get the search keyword from the query string
	query := c.Query("searchKeyword")

	// If the search keyword is empty, search for female harassment news
	if query == "" {
		query = "female%20harass"
	}

	log.Printf("Search keyword: %v", query)

	//Get API key from env
	godotenv.Load(".env")
	apiKey := os.Getenv("NEWS_API_KEY")

	// Make a GET request to the News API
	resp, err := http.Get("https://newsapi.org/v2/everything?q=" + query + "&apiKey=" + apiKey)
	if err != nil {
		log.Printf("Error: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Print the body
	log.Printf("Response body: %v", string(body))

	// Send the response body to the client
	return c.SendString(string(body))
}
