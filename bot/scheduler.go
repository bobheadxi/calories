package bot

import (
	"log"
	"strconv"
	"time"
)

// scheduler : start and check every hour to send scheduled summaries to users
func (b *Bot) scheduler() {
	t := time.NewTicker(time.Second * 1)
	go func() {
		for _ = range t.C {
			b.sendSummary()
		}
	}()
}

// sendSummary : appropriately send all users the scheduled summaries depending on their timezone
func (b *Bot) sendSummary() {
	summaryTime := 22 // 24-hour time for sending out summary

	utcTime := time.Now().UTC()
	timezone := summaryTime - utcTime.Hour() - 24
	log.Print("Ticking... " + strconv.Itoa(utcTime.Hour()))
	log.Print(timezone)
	users, err := b.server.GetUsersInTimezone(timezone)
	if err != nil {
		log.Print("Get users failed: " + err.Error())
	}
	for u, v := range *users {
		b.api.SendTextMessage(u.ID, "You consumed"+strconv.Itoa(v)+"calories today")
	}
}
