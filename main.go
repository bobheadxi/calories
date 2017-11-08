package main

import (
	"log"
	"net/http"

	"github.com/bobheadxi/calories/bot"
	"github.com/bobheadxi/calories/config"
	"github.com/bobheadxi/calories/facebook"
)

var b = &bot.Bot{}

func main() {
	config := config.GetenvConfig()
	if config.Port == "" {
		log.Fatal("$PORT must be set")
	}

	// Attach API instance to Bot
	api := facebook.New(config)
	b.SetAPI(api)

	// Requests made to our /webhook endpont will be handled by API module
	http.HandleFunc("/webhook", b.API.Handler)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
