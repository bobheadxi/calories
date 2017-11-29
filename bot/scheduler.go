package bot

import (
	"log"
	"strconv"
	"time"
)

// scheduler : start and check every hour to send scheduled summaries to users
func (b *Bot) scheduler() {
	t := time.NewTicker(time.Hour * 1)
	for {
		b.sendSummary()
		<-t.C
	}
}

// sendSummary : appropriately send all users the scheduled summaries depending on their timezone
func (b *Bot) sendSummary() {
	utcTime := time.Now().UTC()
	timezone := 22 - utcTime.Hour()
	users, err := b.server.GetUsersInTimezone(timezone)
	if err != nil {
		log.Print("Get users failed: " + err.Error())
	}
	for u, v := range *users {
		b.api.SendTextMessage(u.ID, "You consumed"+strconv.Itoa(v)+"calories today")
	}
}
