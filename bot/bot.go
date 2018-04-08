/*
Package bot contains implementation of the Calories bot.
The bot's Message handlers and functions are described in this package.
*/
package bot

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/bobheadxi/calories/facebook"
	"github.com/bobheadxi/calories/server"
	"github.com/jdkato/prose/tag"
	"github.com/jdkato/prose/tokenize"
)

// Bot : The Calories bot of the app.
type Bot struct {
	api       facebook.APILayer
	server    server.ServerLayer
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
func New(api facebook.APILayer, sv server.ServerLayer) *Bot {
	b := Bot{
		api:    api,
		server: sv,
	}
	b.api.SetHandlers(b.MessageHandler, b.PostbackHandler)

	// Add new command keywords here
	commands := map[string]Handler{
		"help": b.help,
		"test": b.test,
		"update": b.update,
	}
	b.commands = commands

	// Add new postback events here
	postbacks := map[string]Handler{
		"INIT_NEW_USER": b.initUser,
		"TEST_1":        b.testOne,
		"TEST_2":        b.testTwo,
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
	go b.scheduler()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// MessageHandler : Handles messages
func (b *Bot) MessageHandler(event facebook.Event, sender facebook.Sender, msg facebook.ReceivedMessage) error {
	context := &Context{
		facebookID: sender.ID,
		timestamp:  event.Time,
		content:    msg.Text,
	}

	handler := b.help
	words := tokenize.NewTreebankWordTokenizer().Tokenize(strings.ToLower(context.content))
	tagger := tag.NewPerceptronTagger()
	for _, tok := range tagger.Tag(words) {
		// 'VB', 'VBG', 'VBN', 'VBP', 'VBZ' are verb types
		fmt.Println(tok.Text, tok.Tag)
		if tok.Tag == "VB" {
			h, found := b.commands[tok.Text]
			if found {
				handler = h
			}
		}
	}
	return handler(context)
}

// PostbackHandler : Handles postbacks
func (b *Bot) PostbackHandler(event facebook.Event, sender facebook.Sender, postback facebook.Postback) error {
	context := &Context{
		facebookID: sender.ID,
		timestamp:  event.Time,
		content:    postback.Payload,
	}
	handler, found := b.postbacks[postback.Payload]
	if found {
		return handler(context)
	}
	return nil
}
