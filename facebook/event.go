package facebook

import "encoding/json"

type EventStream struct {
	Object  string          `json:"object"`
	Entries []*MessageEvent `json:"entry"`
}

// MessageEvent : a message event from Facebook
type MessageEvent struct {
	Event
	Messaging []struct {
		Sender  *Sender          `json:"sender"`
		Message *ReceivedMessage `json:"message,omitempty"`
	} `json:"messaging"`
}

type Event struct {
	ID   json.Number `json:"id"`
	Time int64       `json:"time"`
}

type Sender struct {
	ID string `json:"id"`
}

// ReceivedMessage : the content of the message that was received
// by the bot
type ReceivedMessage struct {
	ID   string `json:"mid"`
	Text string `json:"text,omitempty"`
	Seq  int    `json:"seq"`
}
