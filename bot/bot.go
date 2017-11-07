package bot

import (
	"fmt"

	"github.com/bobheadxi/calories/facebook"
)

type Bot struct {
	api *facebook.API
}

func New(api *facebook.API) *Bot {
	b := &Bot{
		api: api,
	}

	b.api.SetHandlers(b.TestMessageReceivedAndReply)
	return b
}

// TestMessageReceivedAndReply : Test that bot receives messages and replies.
// DEPRECATE ASAP - replace with Bot handlers or something
func (b *Bot) TestMessageReceivedAndReply(event facebook.Event, opts facebook.MessageOpts, msg facebook.ReceivedMessage) {
	resp, err := b.api.SendTextMessage(opts.Sender.ID, fmt.Sprintf("Hello! %s", msg.Text))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", opts)
}
