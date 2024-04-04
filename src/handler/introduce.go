package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"openrobo-api/src/config"

	"github.com/forPelevin/gomoji"
	"github.com/labstack/echo/v4"
)

func IntroduceHandler(c echo.Context) error {
	if c.Request().Header.Get("Authorization") != "Bearer "+config.Token {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	smiley := c.QueryParam("smiley")

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
	req, err := http.NewRequest("POST", config.OpenAIURL, bytes.NewBuffer([]byte(data)))

	// If there is an error with our request
	if err != nil {
		log.Println(err)
	}

	// Add the API key and content type as headers
	req.Header.Add("Authorization", "Bearer "+config.OpenAIKey)
	req.Header.Add("Content-Type", "application/json")

	// Send the request via a client
	client := &http.Client{}
	resp, err := client.Do(req)

	// Handle the error
	if err != nil {
		log.Println(err)
	}

	// Read the response body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	// Unmarshal the JSON response
	var response config.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		// log the erorr wthout exiting
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Error while getting data from delta OpenAI")
	}

	// Get the AI's message content
	aiMessage := response.Choices[0].Message.Content

	if smiley != "yes" {
		//remove all smileys from the message
		aiMessage = gomoji.RemoveEmojis(aiMessage)
	}

	return c.String(http.StatusOK, aiMessage)
}
