package models

import (
	"context"
	"fmt"
	"maps"
	"slices"
)

type Provider interface {
	Generate(ctx context.Context, prompt string) (string, error)
}

var providers map[string]Provider

func RegisterProvider(name string, provider Provider) {
	if providers == nil {
		providers = make(map[string]Provider)
	}
	providers[name] = provider
}

func ListProviders() []string {
	return slices.Collect(maps.Keys(providers))
}

func init() {
	providers = map[string]Provider{}
}

func GetProvider(name string) (Provider, error) {
	p, exists := providers[name]
	if !exists {
		return nil, fmt.Errorf("unsupported provider: %s", name)
	}
	return p, nil
}
