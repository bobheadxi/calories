package bot

// Tests for the bot package

import (
	"errors"
	"testing"

	"github.com/bobheadxi/calories/facebook"
	"github.com/bobheadxi/calories/server"
)

// TestNewBot : Test Bot instantiation
func TestNewBot(t *testing.T) {
	api := facebook.API{}
	ser := fakeServer{}
	b := New(&api, &ser)
	if b == nil {
		t.Errorf("Bot instantiation failed")
	}

	// TODO
}

// TestMessageHandlerTokenization : Test standard message handler's text tokenization
func TestMessageHandlerTokenization(t *testing.T) {
	b := Bot{}
	b.commands = map[string]Handler{
		"eat": tokenizerHandler,
	}

	// the "eat" command should be recognized from a variety of contexts
	msg := facebook.ReceivedMessage{Text: "I will eat an apple today"}
	err := b.MessageHandler(facebook.Event{}, facebook.Sender{}, msg)
	if err == nil {
		t.Errorf("Expected handler was not called")
	}
	if err.Error() != msg.Text {
		t.Errorf("Expected handler was not called correctly")
	}
}

func tokenizerHandler(c *Context) error {
	return errors.New(c.content)
}

// fakeServer is an alternative to the real server for testing, implements ServerLayer
type fakeServer struct{}

func (fs *fakeServer) AddUser(u server.User) error                  { return nil }
func (fs *fakeServer) AddEntry(e server.Entry) error                { return nil }
func (fs *fakeServer) GetUser(s string) (*server.User, error)       { return &server.User{}, nil }
func (fs *fakeServer) GetEntries(s string) (*[]server.Entry, error) { return nil, errors.New("") }
func (fs *fakeServer) SumCalories(s string) (int, error)            { return 0, nil }
func (fs *fakeServer) UpdateUserTimezone(u server.User) error       { return nil }
func (fs *fakeServer) GetUsersInTimezone(i int) (*map[*server.User]int, error) {
	return nil, errors.New("")
}
