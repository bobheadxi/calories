package config

import (
	"os"
	"testing"
)

// TestGetEnvConfig : test Config var fetching in Travis
func TestGetEnvConfig(t *testing.T) {
	travis := os.Getenv("TRAVIS")
	if travis == "true" {
		cfg, err := GetEnvConfig()
		if err != nil {
			t.Errorf("Get environment config failed: " + err.Error())
		}
		if cfg.Token != "123456" {
			t.Errorf("Travis environment config incorrect")
		}
	}
}
