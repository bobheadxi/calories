package server

// Tests for the facebook package

import (
	"os"
	"testing"

	"github.com/bobheadxi/calories/config"
)

// TestNewServer : test Server instantiation
func TestNewServer(t *testing.T) {
	_ = newTestDBConnection()
}

func newTestDBConnection() *Server {
	travis := os.Getenv("TRAVIS")
	var cfg config.EnvConfig
	if travis == "true" {
		cfg = config.EnvConfig{
			DatabaseURL: "postgresql://localhost/test_db",
		}
	} else {
		cfg = config.EnvConfig{
			DatabaseURL: "postgresql://localhost",
		}
	}
	return New(&cfg)
}

func TestUserFunctions(t *testing.T) {
	ser := newTestDBConnection()

	u := User{
		ID:       "Some random id context from facebook",
		MaxCal:   2000,
		Timezone: -8,
		Name:     "Random Name",
	}
	ser.AddUser(u)

	usr, err := ser.GetUser("Some random id context from facebook")
	if err != nil {
		t.Errorf("Errored: " + err.Error())
	}
	if usr.ID != u.ID {
		t.Errorf("Errored: usr.ID doesn't match u.ID")
	}
	if usr.MaxCal != u.MaxCal {
		t.Errorf("Errored: usr.MaxCal doesn't match u.MaxCal")
	}
	if usr.Timezone != u.Timezone {
		t.Errorf("Errored: usr.ID doesn't match u.ID")
	}
	if usr.Name != u.Name {
		t.Errorf("Errored: usr.Name doesn't match u.Name")
	}
}

func TestEntryFunctions(t *testing.T) {
	ser := newTestDBConnection()

	e := Entry{
		ID:       "Some random id context from facebook",
		Time:     123456789,
		Item:     "Some random item context from facebook",
		Calories: 5,
	}
	ser.AddEntry(e)

	entries, err := ser.GetEntries(e.ID)
	if err != nil {
		t.Errorf("Errored: " + err.Error())
	}

	// Check if entry is as expected
	sum := 0
	for i := range *entries {
		entry := (*entries)[i]

		if entry.ID != e.ID {
			t.Errorf("Errored: entry.ID doesn't match e.ID")
		}
		if entry.Time != e.Time {
			t.Errorf("Errored: entry.Time doesn't match e.Time")
		}
		if entry.Item != e.Item {
			t.Errorf("Errored: entry.Item doesn't match e.Item")
		}
		if entry.Calories != e.Calories {
			t.Errorf("Errored: entry.Calories doesn't match e.Calories")
		}
		sum += entry.Calories
	}

	// Check if sum matches SumCalories()
	sum2, err := ser.SumCalories(e.ID)
	if err != nil {
		t.Errorf("Errored: " + err.Error())
	}
	if sum2 != sum {
		t.Errorf("Errored: SumCalories doesn't match actual sum")
	}
}
