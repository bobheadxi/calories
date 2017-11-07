package bot

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bobheadxi/calories/facebook"
)

type Bot struct {
	api *facebook.API
}

func (b *Bot) SetApi(api *facebook.API) {
	b.api = api
	b.api.SetHandlers(b.TestMessageReceivedAndReply)
}

// TestMessageReceivedAndReply : Test that bot receives messages and replies.
// DEPRECATE ASAP - replace with Bot handlers or something
func (b *Bot) TestMessageReceivedAndReply(event facebook.Event, opts facebook.MessageOpts, msg facebook.ReceivedMessage) {
	b.api.SendTextMessage(opts.Sender.ID, fmt.Sprintf("Hello! %s", msg.Text))
	fmt.Printf("%+v", opts)
}

// Handler : Listens for all HTTP requests and decides what to do with them
func (b *Bot) Handler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		query := req.URL.Query()
		verifyToken := query.Get("hub.verify_token")
		if verifyToken == b.api.Token {
			rw.WriteHeader(http.StatusOK)
			log.Println("RET:", query.Get("hub.challenge"))
			rw.Write([]byte(query.Get("hub.challenge")))
		} else {
			log.Println("GET Request Unauthorized")
		}
	} else if req.Method == "POST" {
		b.api.HandlePOST(rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
