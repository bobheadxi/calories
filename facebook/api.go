package facebook

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
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

// Handler : Listens for all HTTP requests and decides what to do with them
func (api *API) Handler(rw http.ResponseWriter, req *http.Request) {
	log.Println("Handler accepted Request")
	if req.Method == "GET" {
		log.Println("Checking authentication")
		query := req.URL.Query()
		verifyToken := query.Get("hub.verify_token")
		if verifyToken != api.Token {
			rw.WriteHeader(http.StatusUnauthorized)
			log.Println("Authentication failed")
			return
		}
		rw.WriteHeader(http.StatusOK)
		log.Println("Authentication success")
		rw.Write([]byte(query.Get("hub.challenge")))
	} else if req.Method == "POST" {
		api.handlePOST(rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// HandlePOST : works on all POST requests passed to the server
func (api *API) handlePOST(rw http.ResponseWriter, req *http.Request) {
	log.Println("POST request received: ")
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
				log.Print("type message")
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
