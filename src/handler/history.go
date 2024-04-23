package handler

import (
	"net/http"
	"openrobo-api/src/aiclient"
	"openrobo-api/src/config"

	"github.com/forPelevin/gomoji"
	"github.com/labstack/echo/v4"
)

func HistoryHandler(c echo.Context) error {
	if c.Request().Header.Get("Authorization") != "Bearer "+config.Token {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	smiley := c.QueryParam("smiley")

	// Get the AI's message content
	message_history := aiclient.GetMessages()

	if message_history == nil {
		message_history = []map[string]string{}
	}

	if smiley != "yes" {
		//remove all smileys from the messages
		for i, message := range message_history {
			message_history[i]["content"] = gomoji.RemoveEmojis(message["content"])
		}
	}

	return c.JSON(http.StatusOK, message_history)
}
