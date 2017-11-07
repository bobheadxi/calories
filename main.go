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

	api := facebook.New(config)
	b.SetApi(api)
	http.HandleFunc("/webhook", b.Handler)
}
