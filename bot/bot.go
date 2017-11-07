package bot

import (
	"fmt"
	"log"

	"github.com/bobheadxi/calories/facebook"
)

type Bot struct {
	API *facebook.API
}

func (b *Bot) SetApi(api *facebook.API) {
	b.API = api
	b.API.MessageHandler = b.TestMessageReceivedAndReply
}

// TestMessageReceivedAndReply : Test that bot receives messages and replies.
// DEPRECATE ASAP - replace with Bot handlers or something
func (b *Bot) TestMessageReceivedAndReply(event facebook.Event, opts facebook.MessageOpts, msg facebook.ReceivedMessage) {
	log.Println("Test receiver has received message")
	b.API.SendTextMessage(opts.Sender.ID, "Hello!")
	fmt.Printf("%+v", opts)
}
