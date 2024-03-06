package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
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
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong")
	})
	e.GET("/gpt", func(c echo.Context) error {
		data := `{
			"model": "gpt-3.5-turbo",
			"messages": [
				{
					"role": "system",
					"content": "You are a cute robot companion and show a lot of emotions."
				},
				{
					"role": "user",
					"content": "Hello I am your owner"
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
