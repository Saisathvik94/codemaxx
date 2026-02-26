package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Saisathvik94/codemaxx/internal/keys"
	"github.com/Saisathvik94/codemaxx/internal/models"
)

type OpenAIProvider struct{}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model       string    `json:"model"`
	Messages    []message `json:"messages"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
}

type Response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (p OpenAIProvider) Generate(ctx context.Context, prompt string) (string, error) {
	key, err := keys.GetKey("openai")

	if err != nil {
		return "", fmt.Errorf("OpenAI Key is not added")
	}

	reqBody := Request{
		Model: "gpt-4o-mini",
		Messages: []message{
			{Role: "user", Content: prompt},
		},
		MaxTokens:   512,
		Temperature: 0.2,
	}

	jsonData, err := json.Marshal(reqBody)

	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", key))

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		return "", fmt.Errorf("OpenAI API returned status %d: %s", response.StatusCode, string(bodyBytes))
	}

	var Resp Response

	if err := json.NewDecoder(response.Body).Decode(&Resp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(Resp.Choices) == 0 {
		return "", fmt.Errorf("No response from OpenAI")
	}

	return Resp.Choices[0].Message.Content, nil
}

func init() {
	models.RegisterProvider("openai", OpenAIProvider{})
}
