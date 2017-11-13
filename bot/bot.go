/*
Package bot contains implementation of the Calories bot.
The bot's Message handlers and functions are described in this package.
*/
package bot

import (
	"log"
	"net/http"

	"github.com/bobheadxi/calories/facebook"
	"github.com/bobheadxi/calories/server"
)

// Bot : The Calories bot of the app.
type Bot struct {
	api       *facebook.API
	server    *server.Server
	commands  map[string]Handler
	postbacks map[string]Handler
}

// Handler : A function that handles an event
type Handler func(*Context) error

// Context :
type Context struct {
	facebookID string
	timestamp  int64
	content    string
}

// New : Sets up and returns a Bot
func New(api *facebook.API, sv *server.Server) *Bot {
	b := Bot{
		api:    api,
		server: sv,
	}
	b.api.MessageHandler = b.MessageHandler
	b.api.PostbackHandler = b.PostbackHandler

	// Add new command keywords here
	commands := map[string]Handler{
		"help": b.help,
		"test": b.test,
	}
	b.commands = commands

	// Add new postback events here
	postbacks := map[string]Handler{
		"INIT_NEW_USER": b.initUser,
	}
	b.postbacks = postbacks
	return &b
}

// Run : Spins up the Calories bot
func (b *Bot) Run(port string) {
	log.Print("Spinning up the bot...")
	http.HandleFunc("/webhook", b.api.Handler)
	err := b.api.SetWelcomeScreen()
	if err != nil {
		log.Print(err)
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// MessageHandler : Handles messages
func (b *Bot) MessageHandler(event facebook.Event, sender facebook.Sender, msg facebook.ReceivedMessage) {
	context := &Context{
		facebookID: sender.ID,
		timestamp:  event.Time,
		content:    msg.Text,
	}
	// TODO : make command recognition smarter
	handler, found := b.commands[msg.Text]
	if !found {
		handler = b.help
	}
	err := handler(context)
	if err != nil {
		log.Print(err)
	}
}

// PostbackHandler : Handles postbacks
func (b *Bot) PostbackHandler(event facebook.Event, sender facebook.Sender, postback facebook.Postback) {
	context := &Context{
		facebookID: sender.ID,
		timestamp:  event.Time,
		content:    postback.Payload,
	}
	handler, found := b.postbacks[postback.Payload]
	if found {
		err := handler(context)
		if err != nil {
			log.Print(err)
		}
	}
}
