package ai

import (
	"context"
	"fmt"

	"github.com/Saisathvik94/codemaxx/internal/config"
	"github.com/Saisathvik94/codemaxx/internal/models"
)

type Request struct {
	Provider string
	Prompt   string
}

type Response struct {
	Content  string
	Provider string
}

func Generate(ctx context.Context, req Request) (Response, error) {

	var providerName = req.Provider

	if providerName == "" {
		providerName = config.GetDefaultProvider()
	}

	provider, err := models.GetProvider(providerName)
	if err != nil {
		return Response{}, fmt.Errorf("unsupported provider %s", providerName)
	}

	content, err := provider.Generate(ctx, req.Prompt)
	if err != nil {
		return Response{}, err
	}

	return Response{
		Content:  content,
		Provider: providerName,
	}, nil
}
