package facebook

import (
	"encoding/json"
	"fmt"
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
	VerifyToken    string
	AccessToken    string
	AppSecret      string
	MessageHandler MessageHandler
}

// New : Build a new instance of our Facebook API Handler
func New(config *config.EnvConfig) *API {
	return &API{
		PageID:         config.PageID,
		VerifyToken:    config.VerifyToken,
		AccessToken:    config.AccessToken,
		AppSecret:      config.AppSecret,
		MessageHandler: TestMessageReceivedAndReply,
	}
}

// TestMessageReceivedAndReply : Test that bot receives messages and replies.
// DEPRECATE ASAP - replace with Bot handlers or something
func TestMessageReceivedAndReply(event Event, opts MessageOpts, msg ReceivedMessage) {
	resp, err := SendTextMessage(opts.Sender.ID, fmt.Sprintf("Hello   , %s", msg.Text))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", opts)
}

// Handler : the main handler for the API service, receives all HTTP
// requests and decides what to do with them
func (api *API) Handler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		query := req.URL.Query()
		verifyToken := query.Get("hub.verify_token")
		if verifyToken == api.VerifyToken {
			rw.WriteHeader(http.StatusOK)
			log.Println("RET:", query.Get("hub.challenge"))
			rw.Write([]byte(query.Get("hub.challenge")))
		} else {
			log.Println("GET Request Unauthorized")
		}
	} else if req.Method == "POST" {
		api.handlePOST(rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// handlePOSTt : works on all POST requests passed to the server
func (api *API) handlePOST(rw http.ResponseWriter, req *http.Request) {
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
	query.Set("access_token", api.AccessToken)
	req.URL.RawQuery = query.Encode()
	return http.DefaultClient.Do(req)
}
