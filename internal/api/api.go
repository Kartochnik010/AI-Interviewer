package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"kolesagpt/config"
	"kolesagpt/internal/models"
	"log"
	"net/http"
	"strings"
)

func PromptGPT(client *http.Client, cfg *config.Config, userdata *models.User) (*models.GptResponse, error) {
	gptRequest := &models.GptRequest{
		Model:       "gpt-4",
		Messages:    userdata.Messages,
		Temperature: 0.2,
	}
	return SendRequest(client, gptRequest, cfg)
}

func SendRequest(client *http.Client, gptRequest *models.GptRequest, cfg *config.Config) (*models.GptResponse, error) {
	b, err := json.Marshal(gptRequest)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", cfg.GptURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", strings.Trim(cfg.GptToken, "\"")))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	log.Println("responce body:" + string(body))
	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}

	res := &models.GptResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
