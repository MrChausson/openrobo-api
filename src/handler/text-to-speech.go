package handler

import (
	"io/ioutil"
	"net/http"
	"openrobo-api/src/config"
	"os"
	"path/filepath"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
	"github.com/labstack/echo/v4"
)

func TextToSpeechHandler(c echo.Context) error {
	if c.Request().Header.Get("Authorization") != "Bearer "+config.Token {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	text := c.QueryParam("text")

	speech := htgotts.Speech{Folder: "audio", Language: voices.English}
	speech.Speak(text)

	// Get the most recent file
	files, err := ioutil.ReadDir("audio")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error reading directory")
	}

	var newestFile os.FileInfo
	for _, file := range files {
		if file.Mode().IsRegular() && filepath.Ext(file.Name()) == ".mp3" {
			if newestFile == nil || file.ModTime().After(newestFile.ModTime()) {
				newestFile = file
			}
		}
	}

	if newestFile == nil {
		return c.String(http.StatusInternalServerError, "No MP3 files found")
	}

	return c.Attachment(filepath.Join("audio", newestFile.Name()), "speech.mp3")
}
