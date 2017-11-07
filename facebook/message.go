package facebook

// Package : message content
type Package struct {
	Text string `json:"text,omitempty"`
	// TODO: Add support for attachments?
	// Attachment *Attachment `json:"attachment,omitempty"`
}

// Recipient : the user to deliver a message to
type Recipient struct {
	ID string `json:"id,omitempty"`
}

// Message : a package and recepient to be delivered
type Message struct {
	Recipient Recipient `json:"recipient"`
	Message   Package   `json:"message"`
}
