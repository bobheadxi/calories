package main

import (
	"log"
	"net/http"

	"github.com/bobheadxi/calories/config"
	"github.com/bobheadxi/calories/facebook"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.GetenvConfig()
	if config.Port == "" {
		log.Fatal("$PORT must be set")
	}

	api := facebook.New(config)
	router := gin.New()
	http.HandleFunc("/webhook", api.Handler)
}
