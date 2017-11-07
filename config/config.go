package config

import "os"

// EnvConfig contains all configuration stuff from the environment
type EnvConfig struct {
	Port        string
	VerifyToken string
	AccessToken string
	AppSecret   string
	PageID      string
}

// GetenvConfig gets all configuration stuff from the environment
func GetenvConfig() *EnvConfig {
	return &EnvConfig{
		Port:        os.Getenv("PORT"),
		VerifyToken: os.Getenv("FB_VERIFY_TOKEN"),
		AccessToken: os.Getenv("FB_ACCESS_TOKEN"),
		AppSecret:   os.Getenv("FB_APP_SECRET"),
		PageID:      os.Getenv("FB_PAGE_ID"),
	}
}
