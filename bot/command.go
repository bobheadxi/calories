package bot

import (
	"strconv"

	"github.com/bobheadxi/calories/facebook"
	"github.com/bobheadxi/calories/server"
)

func (b *Bot) help(c *Context) error {
	return b.api.SendTextMessage(c.facebookID, "Hello! I'm pretty useless and can't really do anything right now, sorry :(")
}

// test : test new functions here - deprecate when possible
func (b *Bot) test(c *Context) error {
	e := server.Entry{
		ID:       c.facebookID,
		Time:     c.timestamp,
		Item:     c.content,
		Calories: 5,
	}
	err := b.server.AddEntry(e)
	if err != nil {
		b.api.SendTextMessage(c.facebookID, "No new entry for you: "+err.Error())
	}
	response, err := b.server.SumCalories(c.facebookID)
	if err != nil {
		b.api.SendTextMessage(c.facebookID, "No calories for you: "+err.Error())
	}
	b.api.SendTextMessage(c.facebookID, "your total calories are "+strconv.Itoa(response))
	users, err := b.server.GetUsersInTimezone(-8)
	if err != nil {
		b.api.SendTextMessage(c.facebookID, "Get users failed: "+err.Error())
	}
	for u, v := range *users {
		b.api.SendTextMessage(u.ID, "Someone is snooping on ya ;)")
		b.api.SendTextMessage(c.facebookID, u.Name+" has eaten "+strconv.Itoa(v)+" calories :o")
	}
	qr := []facebook.QuickReply{
		facebook.QuickReply{
			ContentType: "text",
			Title:       "Yes!",
			Payload:     "TEST_1",
		},
		facebook.QuickReply{
			ContentType: "text",
			Title:       "No!",
			Payload:     "TEST_2",
		},
	}
	return b.api.SendQuickReplyTemplate(c.facebookID, "Are you hungry?", qr)
}


func (b *Bot) update(c *Context) error {
	user, err := b.api.GetUserProfile(c.facebookID)
	if err != nil {
		return err
	}
	u := server.User{
		ID:       c.facebookID,
		Timezone: user.Timezone,
	}
	return b.server.UpdateUserTimezone(u)
}
