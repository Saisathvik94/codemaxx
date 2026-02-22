package keys

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Saisathvik94/codemaxx/internal/config"
)

// GetKey returns API key for provider
func GetKey(provider string) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		return "", err
	}

	key := cfg.Keys[provider]
	if key == "" {
		return "", errors.New("api key not set")
	}

	return key, nil
}

// SetKey stores API key for provider
func SetKey(provider, key string) error {

	if strings.TrimSpace(key) == "" {
		return fmt.Errorf("cannot save an empty API key for provider: %s", provider)
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	cfg.Keys[provider] = key
	return config.Save(cfg)
}
