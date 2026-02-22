package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	DefaultProvider string            `json:"default_provider"`
	Keys            map[string]string `json:"keys"`
}

// getConfigPath returns the path to the config file
func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".codemaxx.json"), nil
}

// Load reads configuration from disk
func Load() (*Config, error) {
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	// If config doesn't exist → create default
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return &Config{
			DefaultProvider: "openai",
			Keys:            make(map[string]string),
		}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	// Safety: ensure Keys map is not nil
	if cfg.Keys == nil {
		cfg.Keys = make(map[string]string)
	}

	return &cfg, nil
}

// Save writes configuration to disk
func Save(cfg *Config) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// GetDefaultProvider returns the default provider
func GetDefaultProvider() string {
	cfg, err := Load()
	if err != nil {
		return "openai" // fallback
	}
	return cfg.DefaultProvider
}

// SetDefaultProvider updates the default provider
func SetDefaultProvider(provider string) error {
	cfg, err := Load()
	if err != nil {
		return err
	}

	cfg.DefaultProvider = provider
	return Save(cfg)
}
