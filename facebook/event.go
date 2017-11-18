package facebook

import "encoding/json"

// EventStream : Contains list of events coming from Facebook
type EventStream struct {
	Object  string          `json:"object"`
	Entries []*MessageEvent `json:"entry"`
}

// MessageEvent : A message event from Facebook
type MessageEvent struct {
	Event
	Messaging []struct {
		Sender   *Sender          `json:"sender"`
		Message  *ReceivedMessage `json:"message,omitempty"`
		Postback *Postback        `json:"postback,omitempty"`
	} `json:"messaging"`
}

// Postback : A postback payload (typically from button presses, etc)
type Postback struct {
	Payload string `json:"payload"`
}

// Event : Data about a MessageEvent
type Event struct {
	ID   json.Number `json:"id"`
	Time int64       `json:"time"`
}

// Sender : The sender of a ReceivedMessage
type Sender struct {
	ID string `json:"id"`
}

// ReceivedMessage : The content of a message that was received
type ReceivedMessage struct {
	Text     string    `json:"text,omitempty"`
	Postback *Postback `json:"quick_reply,omitempty"`
}
