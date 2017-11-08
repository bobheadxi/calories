package facebook

// Package : Holds the content of a Message
type Package struct {
	Text string `json:"text,omitempty"`
	// TODO: Add support for attachments?
	// Attachment *Attachment `json:"attachment,omitempty"`
}

// Recipient : The user to deliver a Message to
type Recipient struct {
	ID string `json:"id,omitempty"`
}

// Message : A package and recepient to be delivered
type Message struct {
	Recipient Recipient `json:"recipient"`
	Message   Package   `json:"message"`
}
