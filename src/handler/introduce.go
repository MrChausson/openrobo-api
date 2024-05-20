package handler

import (
	"log"
	"net/http"
	"openrobo-api/src/aiclient"
	"openrobo-api/src/config"

	"github.com/forPelevin/gomoji"
	"github.com/labstack/echo/v4"
)

func IntroduceHandler(c echo.Context) error {
	if c.Request().Header.Get("Authorization") != "Bearer "+config.Token {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	smiley := c.QueryParam("smiley")

	//Let's reset the conversation
	aiclient.Reset()

	// Get the AI's message content
	aiMessage, err := aiclient.Ask("Hello I am your owner, can you present yourself and explain what you can do ? Please don't forget to keep answers very short")
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Error while getting data from delta OpenAI: "+err.Error())
	}

	if smiley != "yes" {
		//remove all smileys from the message
		aiMessage = gomoji.RemoveEmojis(aiMessage)
	}

	return c.String(http.StatusOK, aiMessage)
}
