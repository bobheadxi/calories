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
		Sender  *Sender          `json:"sender"`
		Message *ReceivedMessage `json:"message,omitempty"`
	} `json:"messaging"`
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
	ID   string `json:"mid"`
	Text string `json:"text,omitempty"`
	Seq  int    `json:"seq"`
}

// DeliveryResponse : Response from Facebook after a message has been delivered
type DeliveryResponse struct {
	RecipientID string `json:"recipient_id"`
	MessageID   string `json:"message_id"`
}
