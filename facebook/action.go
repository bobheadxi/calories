package facebook

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// SendMessage : Delivers a Message (see struct Message from event)
func (api *API) SendMessage(m Message) error {
	byt, err := json.Marshal(m)
	if err != nil {
		return err
	}

	resp, err := api.makeRequest("POST", GraphAPI+"/v2.6/me/messages", bytes.NewReader(byt))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Bad response attempting to send a message to " + m.Recipient.ID)
	}

	return nil
}

// SendTextMessage : Sends a simple text message to specified recipient
func (api *API) SendTextMessage(recipientID string, message string) error {
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

// GetUserProfile : Get the profile of specified user
func (api *API) GetUserProfile(userID string) (*UserProfile, error) {
	resp, err := api.makeRequest("GET", fmt.Sprintf(GraphAPI+"/v2.6/%s?fields=first_name,timezone", userID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Error occured while requesting user profle: " + userID)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	profile := &UserProfile{}
	err = json.Unmarshal(respBody, profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

// SetWelcomeScreen : Sets up a welcome screen that greets first time users
func (api *API) SetWelcomeScreen() error {
	byt, err := json.Marshal(appProfile{
		GetStarted: getStarted{
			Payload: "INIT_NEW_USER",
		},
	})
	if err != nil {
		return err
	}

	resp, err := api.makeRequest("POST", GraphAPI+"/v2.6/me/messenger_profile", bytes.NewReader(byt))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return errors.New("Bad response attempting to set Welcome Screen: " + resp.Status + ", " + string(respBody))
	}
	log.Print("Welcome Screen successfully set.")
	return nil
}
