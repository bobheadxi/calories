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
	errorMessage := "The following variables must be set in your deployment environment:"
	errored := false
	if cfg.Port == "" {
		errorMessage += " $PORT"
		errored = true
	}
	if cfg.DatabaseURL == "" {
		errorMessage += " $DATAASE_URL"
		errored = true
	}
	if cfg.Token == "" {
		errorMessage += " $FB_TOKEN"
		errored = true
	}
	if cfg.PageID == "" {
		errorMessage += " $PAGE_ID"
		errored = true
	}
	if errored {
		return nil, errors.New(errorMessage)
	}
	return cfg, nil
}
