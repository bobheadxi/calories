/*
Package config contains functions and data types to get
necessary parameters from outside the app.
*/
package config

import (
	"errors"
	"os"
)

// EnvConfig : Holds the app's configuration variables
type EnvConfig struct {
	Port        string
	DatabaseURL string
	Token       string
	PageID      string
}

// GetEnvConfig : Gets all configuration variables from ENV
func GetEnvConfig() (*EnvConfig, error) {
	cfg := &EnvConfig{
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Token:       os.Getenv("FB_TOKEN"),
		PageID:      os.Getenv("FB_PAGE_ID"),
	}
	if cfg.Port == "" || cfg.DatabaseURL == "" || cfg.Token == "" || cfg.PageID == "" {
		return nil, errors.New("All configuration variables in EnvConfig must be set in your deployment environment")
	}
	return cfg, nil
}
