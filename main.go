package main

import (
	"log"
	"net/http"

	"github.com/bobheadxi/calories/config"
	"github.com/bobheadxi/calories/facebook"
)

func main() {
	config := config.GetenvConfig()
	if config.Port == "" {
		log.Fatal("$PORT must be set")
	}

	api := facebook.New(config)
	http.HandleFunc("/webhook", api.Handler)
}
