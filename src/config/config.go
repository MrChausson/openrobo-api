package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

var (
	// OpenAI API URL
	OpenAIURL string
	// OpenAI API Key
	OpenAIKey string
	// Unique bearer token for auth
	Token string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		// Log the error and continue. The .env file is not required, and environment variables can still be set manually.
		log.Println("Error loading .env file")
	}

	OpenAIURL = os.Getenv("OPENAIURL")
	OpenAIKey = os.Getenv("OPENAIKEY")
	Token = os.Getenv("TOKEN")
}
