package bot

import "github.com/bobheadxi/calories/server"

func (b *Bot) initUser(c *Context) error {
	profile, err := b.api.GetUserProfile(c.facebookID)
	if err != nil {
		return err
	}
	user := server.User{
		ID:       c.facebookID,
		MaxCal:   100,
		Timezone: profile.Timezone,
		Name:     profile.FirstName,
	}
	b.api.SendTextMessage(c.facebookID, "Nice to meet you, "+user.Name+" :)")
	return b.server.AddUser(user)
}
