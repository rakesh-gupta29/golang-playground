package main

import (
	"app/models"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Row struct {
	Title string `json:"title"`
}

const api = "https://restcountries.com/v3.1/name/"

func Fetch(path string) string {
	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)

}

func main() {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/search", func(c *fiber.Ctx) error {

		apiURL := api + c.Query("query")

		body := Fetch(apiURL)

		data := []models.Country{}
		json.Unmarshal([]byte(body), &data)

		return c.Render("results", fiber.Map{
			"Results":       data,
			"DetailsAPIURL": "https://restcountries.com/v3.1/alpha",
		})

	})

	app.Listen("localhost:8000")
}
