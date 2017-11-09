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
	api    *facebook.API
	server *server.Server
}

// New : Sets up and returns a Bot
func New(api *facebook.API, sv *server.Server) *Bot {
	b := Bot{
		api:    api,
		server: sv,
	}
	b.api.MessageHandler = b.TestMessageReceivedAndReply
	return &b
}

// Run : Spins up the Calories bot
func (b *Bot) Run(port string) {
	http.HandleFunc("/webhook", b.api.Handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// TestMessageReceivedAndReply : Tests that bot receives messages and replies.
// DEPRECATE ASAP - replace with Bot handlers or something
func (b *Bot) TestMessageReceivedAndReply(event facebook.Event, sender facebook.Sender, msg facebook.ReceivedMessage) {
	b.api.SendTextMessage(sender.ID, "Hello!")
	output, err := b.server.InsertDataExample(sender.ID, msg.Text)
	if err != nil {
		b.api.SendTextMessage(sender.ID, "Dang data insertion failed")
		return
	}
	b.api.SendTextMessage(sender.ID, string(output))
	log.Printf("Event: %+v", event)   // {ID:2066945410258565 Time:1510063491984}
	log.Printf("Sender: %+v", sender) // {ID:1657077300989984}
	log.Printf("Msg: %+v", msg)       // {ID:mid.$cAAcNE7mWyw1lyBGR51flsxJvj8_- Text:hello Seq:1028142}
}
