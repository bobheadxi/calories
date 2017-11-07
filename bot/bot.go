package bot

import (
	"fmt"

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
	b.API.SendTextMessage(opts.Sender.ID, fmt.Sprintf("Hello! %s", msg.Text))
	fmt.Printf("%+v", opts)
}
