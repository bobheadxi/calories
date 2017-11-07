package config

import "os"

// EnvConfig contains all configuration stuff from the environment
type EnvConfig struct {
	Port   string
	Token  string
	PageID string
}

// GetenvConfig gets all configuration stuff from the environment
func GetenvConfig() *EnvConfig {
	return &EnvConfig{
		Port:   os.Getenv("PORT"),
		Token:  os.Getenv("FB_TOKEN"),
		PageID: os.Getenv("FB_PAGE_ID"),
	}
}
