package tests

// Tests for the bot package

import (
	"testing"

	"github.com/bobheadxi/calories/bot"
	"github.com/bobheadxi/calories/facebook"
	"github.com/bobheadxi/calories/server"
)

// TestNewBot : Test Bot instantiation
func TestNewBot(t *testing.T) {
	api := facebook.API{}
	ser := server.Server{}
	b := bot.New(&api, &ser)
	if b == nil {
		t.Errorf("Bot instantiation failed")
	}

	// TODO
}
