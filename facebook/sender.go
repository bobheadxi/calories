package facebook

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// MessageResponse : Response from Facebook after a message has been delivered
type MessageResponse struct {
	RecipientID string `json:"recipient_id"`
	MessageID   string `json:"message_id"`
}

// SendMessage : Deliver a Message
func (api *API) SendMessage(m Message) (*MessageResponse, error) {
	byt, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	resp, err := api.makeRequest("POST", GraphAPI+"/v2.6/me/messages", bytes.NewReader(byt))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Bad response")
	}
	response := &MessageResponse{}
	err = json.Unmarshal(read, response)
	log.Print("Message delivery complete")
	return response, err
}

//SendTextMessage : Send a simple text message to specified recipient
func (api *API) SendTextMessage(recipient string, message string) (*MessageResponse, error) {
	return api.SendMessage(
		Message{
			Recipient: Recipient{
				ID: recipient,
			},
			Message: Package{
				Text: message,
			},
		})
}
