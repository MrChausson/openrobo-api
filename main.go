package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// static variables
const (
	// OpenAI API URL
	OpenAIURL = "https://api-endpoint.delta-net.ovh/v1/chat/completions"
	// OpenAI API Key
	OpenAIKey = "bce8141ce7661f6350a3c7ce29aced8f"
)

type Response struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func main() {
	token := os.Getenv("TOKEN")

	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong")
	})
	e.GET("/introduce", func(c echo.Context) error {
		if c.Request().Header.Get("Authorization") != "Bearer "+token {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		data := `{
			"model": "gpt-3.5-turbo",
			"messages": [
				{
					"role": "system",
					"content": "You are a cute robot companion and show a lot of emotions. End every message with a smiley face representing the emotion you are feeling."
				},
				{
					"role": "user",
					"content": "Hello I am your owner, can you present yourself and explain what you can do ?"
				}
			]
		}`

		// Create a new request using http
		req, err := http.NewRequest("POST", OpenAIURL, bytes.NewBuffer([]byte(data)))

		// If there is an error with our request
		if err != nil {
			log.Fatalln(err)
		}

		// Add the API key and content type as headers
		req.Header.Add("Authorization", "Bearer "+OpenAIKey)
		req.Header.Add("Content-Type", "application/json")

		// Send the request via a client
		client := &http.Client{}
		resp, err := client.Do(req)

		// Handle the error
		if err != nil {
			log.Fatalln(err)
		}

		// Read the response body
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		// Unmarshal the JSON response
		var response Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Fatalln(err)
		}

		// Get the AI's message content
		aiMessage := response.Choices[0].Message.Content

		return c.String(http.StatusOK, aiMessage)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
