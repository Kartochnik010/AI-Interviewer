package config

import (
	"errors"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port     string
	GptToken string
	GptURL   string
}

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	gptToken := os.Getenv("GPT_TOKEN")
	if gptToken == "" {
		return nil, errors.New("unset token")
	}

	gptURL := os.Getenv("GPT_URL")
	if gptURL == "" {
		return nil, errors.New("unset url")
	}

	return &Config{
		Port:     port,
		GptToken: gptToken,
		GptURL:   gptURL,
	}, nil
}
