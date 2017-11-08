/*
Package config contains functions and data types to get
necessary parameters from outside the app.
*/
package config

import "os"

// EnvConfig : Holds the app's configuration variables
type EnvConfig struct {
	Port        string
	DatabaseURL string
	Token       string
	PageID      string
}

// GetenvConfig : Gets all configuration variables from ENV
func GetenvConfig() *EnvConfig {
	return &EnvConfig{
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Token:       os.Getenv("FB_TOKEN"),
		PageID:      os.Getenv("FB_PAGE_ID"),
	}
}
