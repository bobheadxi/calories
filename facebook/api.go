package facebook

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/bobheadxi/calories/config"
)

var (
	//GraphAPI specifies host used for API requests
	GraphAPI = "https://graph.facebook.com"
)

// MessageHandler is called when a new message is received
type MessageHandler func(Event, MessageOpts, ReceivedMessage)

// API is the main service which handles all callbacks from facebook,
// events are given to appropriate handlers
type API struct {
	PageID         string
	Token          string
	MessageHandler MessageHandler
}

// New : Build a new instance of our Facebook API Handler
func New(config *config.EnvConfig) *API {
	return &API{
		PageID: config.PageID,
		Token:  config.Token,
	}
}

// SetHandlers : Sets API's handlers
// Currently just Messages, could add more
func (api *API) SetHandlers(messageHandler MessageHandler) {
	api.MessageHandler = messageHandler
}

// HandlePOST : works on all POST requests passed to the server
func (api *API) HandlePOST(rw http.ResponseWriter, req *http.Request) {
	read, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	event := &EventStream{}
	err = json.Unmarshal(read, event)
	if err != nil {
		return
	}

	// Iterate through event entries
	for _, entry := range event.Entries {
		for _, message := range entry.Messaging {
			if message.Message != nil {
				// start goroutine to handle received message
				if api.MessageHandler != nil {
					go api.MessageHandler(entry.Event, message.MessageOpts, *message.Message)
				}
			}
		}
	}
	rw.WriteHeader(http.StatusOK)
}

// makeRequest : Makes a request to API
func (api *API) makeRequest(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	query := req.URL.Query()
	query.Set("access_token", api.Token)
	req.URL.RawQuery = query.Encode()
	return http.DefaultClient.Do(req)
}
