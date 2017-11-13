package main

import (
	"log"

	"github.com/bobheadxi/calories/bot"
	"github.com/bobheadxi/calories/config"
	"github.com/bobheadxi/calories/facebook"
	"github.com/bobheadxi/calories/server"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.GetEnvConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Set up Postgres connection
	server := server.New(config)

	// Set up Facebook API module
	api := facebook.New(config)

	// Start Bot on port
	bot := bot.New(api, server)
	bot.Run(config.Port)
}
