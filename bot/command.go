package bot

func (b *Bot) help(c *Context) error {
	return b.api.SendTextMessage(c.facebookID, "Hello! I'm pretty useless and can't really do anything right now, sorry :(")
}
