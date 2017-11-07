package facebook

import "encoding/json"

type EventStream struct {
	Object  string          `json:"object"`
	Entries []*MessageEvent `json:"entry"`
}

type Event struct {
	ID   json.Number `json:"id"`
	Time int64       `json:"time"`
}

type MessageOpts struct {
	Sender struct {
		ID string `json:"id"`
	} `json:"sender"`
	Recipient struct {
		ID string `json:"id"`
	} `json:"recipient"`
	Timestamp int64 `json:"timestamp"`
}

type MessageEvent struct {
	Event
	Messaging []struct {
		MessageOpts
		Message  *ReceivedMessage  `json:"message,omitempty"`
		Delivery *DeliveredMessage `json:"delivery,omitempty"`
	} `json:"messaging"`
}

type ReceivedMessage struct {
	ID   string `json:"mid"`
	Text string `json:"text,omitempty"`
	Seq  int    `json:"seq"`
}

type DeliveredMessage struct {
	MessageIDS []string `json:"mids"`
	Watermark  int64    `json:"watermark"`
	Seq        int      `json:"seq"`
}
