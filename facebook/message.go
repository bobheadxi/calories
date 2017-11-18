package facebook

// Package : Holds the content of a Message
// Could be many things
type Package struct {
	Text       string       `json:"text,omitempty"`
	QuickReply []QuickReply `json:"quick_replies,omitempty"`
}

// QuickReply : Defines a quick reply
// https://developers.facebook.com/docs/messenger-platform/send-messages/quick-replies
type QuickReply struct {
	// ContentType should be "text", though can also do "location"
	ContentType string `json:"content_type"`
	// Title is the label on the quickreply
	Title string `json:"title,omitempty"`
	// Payload is the postback string
	Payload string `json:"payload"`
	// ImageURL is an optional url to an icon for the quickreply
	ImageURL string `json:"image_url,omitempty"`
}

// Recipient : The user to deliver a Message to
type Recipient struct {
	ID string `json:"id,omitempty"`
}

// Message : A package and recepient to be delivered
type Message struct {
	MessageType string    `json:"messaging_type"`
	Recipient   Recipient `json:"recipient"`
	Package     Package   `json:"message"`
}
