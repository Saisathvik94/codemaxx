package ollama

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Saisathvik94/codemaxx/internal/models"
)

type OllamaProvider struct{}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model       string    `json:"model"`
	Messages    []message `json:"messages"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
	Stream      bool      `json:"stream"`
}

type Response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (p OllamaProvider) Generate(ctx context.Context, prompt string) (string, error) {

	reqBody := Request{
		Model: "qwen2.5-coder:3b",
		Messages: []message{
			{Role: "user", Content: prompt},
		},
		MaxTokens:   100,
		Temperature: 0.2,
		Stream:      false,
	}

	jsonData, err := json.Marshal(reqBody)

	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:11434/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		return "", fmt.Errorf("returned status %d: %s", response.StatusCode, string(bodyBytes))
	}

	var Resp Response

	if err := json.NewDecoder(response.Body).Decode(&Resp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(Resp.Choices) == 0 {
		return "", fmt.Errorf("No response from Ollama")
	}

	return Resp.Choices[0].Message.Content, nil
}

func init() {
	models.RegisterProvider("ollama", OllamaProvider{})
}
