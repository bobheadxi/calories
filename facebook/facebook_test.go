package facebook

// Tests for the facebook package

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bobheadxi/calories/config"
)

func apiSetUp() *API {
	cfg := config.EnvConfig{
		PageID: "123",
		Token:  "321",
	}

	api := New(&cfg)
	return api
}

// TestNewFacebook : test API instantiation
func TestNewFacebook(t *testing.T) {

	api := apiSetUp()

	if api.MessageHandler != nil && api.PostbackHandler != nil {
		t.Logf("Handler setup incorrect")
	}
	if api.PageID != "123" {
		t.Logf("PageID was incorrect, got: %s, want: %s.", api.PageID, "123")
	}
	if api.Token != "321" {
		t.Logf("Token was incorrect, got: %s, want: %s.", api.Token, "321")
	}
}

// TestSendTextMessage : Test sending a text message under different API
// conditions (ie when Facebook is online and offline)
func TestSendTextMessageAvailable(t *testing.T) {
	/* Test when server ok: Should make request with correct parameters */

	api := apiSetUp()

	// Set up fake Facebook server
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Return StatusOK response when a request is made
		rw.WriteHeader(http.StatusOK)

		// Check request method
		if req.Method != "POST" {
			t.Errorf("Request not POST")
		}

		// Check request body is as expected
		read, err := ioutil.ReadAll(req.Body)
		if err != nil {
			t.Errorf("Bad request body: " + err.Error())
		}
		msg := &Message{}
		err = json.Unmarshal(read, msg)
		if err != nil {
			t.Errorf("Bad message" + err.Error())
		}
		if msg.MessageType != "RESPONSE" {
			t.Errorf("Wrong message type, expecting %s, got %s", "RESPONSE", msg.MessageType)
		}
		if msg.Package.Text != "Hello" {
			t.Errorf("Wrong message text, expecting %s, got %s", "Hello", msg.Package.Text)
		}
		if msg.Recipient.ID != "789" {
			t.Errorf("Wrong recipient, expecting %s, got %s", "789", msg.Recipient.ID)
		}
	}))
	defer ts.Close()

	// Set API to use fake server's URL and call SendTextMesssage
	GraphAPI = ts.URL
	err := api.SendTextMessage("789", "Hello")
	if err != nil {
		t.Errorf("Errored: " + err.Error())
	}
}

func TestSendMessageUnavailable(t *testing.T) {
	/* Test when server down: Should return an error */

	api := apiSetUp()

	// Set up a new fake server
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// Return StatusServiceUnavailable when a request is made
		rw.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer ts.Close()

	// Set API to use fake server's URL and call SendTextMesssage
	GraphAPI = ts.URL
	err := api.SendTextMessage("789", "Hello")
	// Should return error
	if err == nil {
		t.Errorf("No error returned")
	}
}
