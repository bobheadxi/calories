/*
Package bot contains implementation of the Calories bot.
The bot's Message handlers and functions are described in this package.
*/
package bot

import (
	"log"
	"strconv"

	"github.com/bobheadxi/calories/facebook"
	"github.com/bobheadxi/calories/server"
)

// Bot : The Calories bot of the app.
type Bot struct {
	API    *facebook.API
	Server *server.Server
}

// SetupAPI : Assigns an instance of facebook.API to bot and
// sets up appropriate handlers in the API
func (b *Bot) SetupAPI(api *facebook.API) {
	b.API = api
	b.API.MessageHandler = b.TestMessageReceivedAndReply
}

// TestMessageReceivedAndReply : Tests that bot receives messages and replies.
// DEPRECATE ASAP - replace with Bot handlers or something
func (b *Bot) TestMessageReceivedAndReply(event facebook.Event, sender facebook.Sender, msg facebook.ReceivedMessage) {
	b.API.SendTextMessage(sender.ID, "Hello!")
	output, err := b.Server.InsertDataExample(sender.ID, msg.Text)
	if err != nil {
		b.API.SendTextMessage(sender.ID, "Dang data insertion example failed")
	}
	b.API.SendTextMessage(sender.ID, string(output))

	err = b.Server.AddUser(server.User{ID: sender.ID, MaxCal: 69})
	if err != nil {
		b.API.SendTextMessage(sender.ID, "Dang user insertion failed")
	}
	b.API.SendTextMessage(sender.ID, "inserted")

	err = b.Server.AddEntry(server.Entry{ID: sender.ID, Time: event.Time, Item: msg.Text, Calories: 100})
	if err != nil {
		b.API.SendTextMessage(sender.ID, "Dang entry insertion failed")
	}
	b.API.SendTextMessage(sender.ID, "inserted entry")

	user, err := b.Server.GetUser(sender.ID)
	if err != nil {
		b.API.SendTextMessage(sender.ID, "No user found")
	}
	b.API.SendTextMessage(sender.ID, "Your maxcal is "+strconv.Itoa(user.MaxCal))

	entries, err := b.Server.GetEntries(sender.ID)
	if err != nil {
		b.API.SendTextMessage(sender.ID, "No entries")
	}
	for _, entry := range *entries {
		str := "One entry: " + entry.ID + "," + entry.Item + "," + strconv.FormatInt(entry.Time, 10) + "," + strconv.Itoa(entry.Calories)
		b.API.SendTextMessage(sender.ID, str)
	}

	log.Printf("Event: %+v", event)   // {ID:2066945410258565 Time:1510063491984}
	log.Printf("Sender: %+v", sender) // {ID:1657077300989984}
	log.Printf("Msg: %+v", msg)       // {ID:mid.$cAAcNE7mWyw1lyBGR51flsxJvj8_- Text:hello Seq:1028142}
}
