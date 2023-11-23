package config

import (
	"errors"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port         string
	GptToken     string
	GptURL       string
	GptMaxTokens int
	GptModel     string
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
	gptMaxTokens, err := strconv.Atoi(os.Getenv("GPT_MAX_TOKENS"))
	if err != nil {
		return nil, errors.Join(errors.New("error while parsing gptMaxTokens: "), err)
	}
	gptModel := os.Getenv("GPT_MODEL")
	if gptModel == "" {
		return nil, errors.New("unset url")
	}
	return &Config{
		GptModel:     gptModel,
		GptMaxTokens: gptMaxTokens,
		Port:         port,
		GptToken:     gptToken,
		GptURL:       gptURL,
	}, nil
}
