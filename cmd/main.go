package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	logpkg "log"
	"net/http"
	"os"
)

func main() {
	err := serve()
	if err != nil {
		logINFO.Fatalln(err, nil)
	}
}

const (
	port = "8080"
)
const (
	// apiKey  = "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"
	apiKey  = "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"
	url     = "https://api.openai.com/v1/chat/completions"
	debug   = true
	davinci = "text-davinci-003"
	turbo   = "gpt-3.5-turbo"
)

var (
	logINFO  = logpkg.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	logERROR = logpkg.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
)

func GetRequestBodyString(model string, msg string) (string, error) {
	switch model {
	case "text-davinci-003":
		return fmt.Sprintf(`{
			"model": "text-davinci-003",
			"prompt": "%s",
			"max_tokens": 7,
			"temperature": 0,
			"top_p": 1,
			"n": 1,
			"stream": false,
			"logprobs": null,
			"stop": "\n"
	  	}`, msg), nil

	case "gpt-3.5-turbo":
		return fmt.Sprintf(`{
			"model": "gpt-3.5-turbo",
			"messages": [
				{"role": "user", "content": "%s"}
				],
			"temperature": 0.2 
		}`, msg), nil

	}
	return "", errors.New("no such gpt model")
}

// POST request to gpts url
func SendRequest(ctx context.Context, url string, client http.Client, requestBody string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

type gptRes struct {
	ID      string `json:"id,omitempty"`
	Object  string `json:"object,omitempty"`
	Created int    `json:"created,omitempty"`
	Model   string `json:"model,omitempty"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens,omitempty"`
		CompletionTokens int `json:"completion_tokens,omitempty"`
		TotalTokens      int `json:"total_tokens,omitempty"`
	} `json:"usage,omitempty"`
	Choices []struct {
		Message struct {
			Role    string `json:"role,omitempty"`
			Content string `json:"content,omitempty"`
		} `json:"message,omitempty"`
		FinishReason string `json:"finish_reason,omitempty"`
		Index        int    `json:"index,omitempty"`
	} `json:"choices,omitempty"`
}

func PrintDebug(m string) {
	if debug {
		logINFO.Println(m)
	}
}
