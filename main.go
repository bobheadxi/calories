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
	log.Println("Starting app!")
	config := config.GetenvConfig()
	if config.Port == "" {
		log.Fatal("$PORT must be set")
	}

	api := facebook.New(config)
	b.API = api
	http.HandleFunc("/webhook", b.API.Handler)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
