package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type NewsReturnObject struct {
	Index       int       `json:"index"`
	Source      string    `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	ImageURL    string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
}

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
	// log.Printf("Response body: %v", string(body))

	// Parse the JSON response into a slice of NewsReturnObject
	var response struct {
		Articles []struct {
			Source struct {
				Name string `json:"name"`
			} `json:"source"`
			Author      string    `json:"author"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			URL         string    `json:"url"`
			ImageURL    string    `json:"urlToImage"`
			PublishedAt time.Time `json:"publishedAt"`
		} `json:"articles"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Error parsing JSON response: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Convert the parsed articles to NewsReturnObject
	var newsList []NewsReturnObject
	for i, article := range response.Articles {
		news := NewsReturnObject{
			Index:       i,
			Title:       article.Title,
			Description: article.Description,
			URL:         article.URL,
			ImageURL:    article.ImageURL,
			PublishedAt: article.PublishedAt,
			Source:      article.Source.Name,
			Author:      article.Author,
		}
		newsList = append(newsList, news)
	}

	// Convert the newsList to JSON
	body, err = json.Marshal(newsList)
	if err != nil {
		log.Printf("Error marshalling newsList: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	//Print the newslist one by one, with column name
	for i, news := range newsList {
		log.Printf("Index: %v\n", i)
		log.Printf("Title: %v\n", news.Title)
		log.Printf("Description: %v\n", news.Description)
		log.Printf("URL: %v\n", news.URL)
		log.Printf("Image URL: %v\n", news.ImageURL)
		log.Printf("Published At: %v\n", news.PublishedAt)
		log.Printf("Source: %v\n", news.Source)
		log.Printf("Author: %v\n\n\n", news.Author)
	}

	// Return the JSON response
	return c.Send(body)
}
