package facebook

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// SendMessage : Delivers a Message (see struct Message from event)
func (api *API) SendMessage(m Message) (*DeliveryResponse, error) {
	byt, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	resp, err := api.makeRequest("POST", GraphAPI+"/v2.6/me/messages", bytes.NewReader(byt))
	if err != nil {
		return nil, err
	}

	read, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Bad response")
	}

	response := &DeliveryResponse{}
	err = json.Unmarshal(read, response)
	return response, err
}

//SendTextMessage : Sends a simple text message to specified recipient
func (api *API) SendTextMessage(recipientID string, message string) (*DeliveryResponse, error) {
	mes := Message{
		Recipient: Recipient{
			ID: recipientID,
		},
		Message: Package{
			Text: message,
		},
	}
	return api.SendMessage(mes)
}
