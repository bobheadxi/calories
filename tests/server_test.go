package tests

// Tests for the facebook package

import (
	"testing"
	"psql"

	"github.com/bobheadxi/calories/config"
	"github.com/bobheadxi/calories/server"
)

var ser server.Server

// TestNewServer : test Server instantiation
func TestNewServer(t *testing.T) {

	cfg := config.EnvConfig{
		DatabaseURL: "postgresql://localhost",
	}

	ser := server.New(&cfg);
	_ = ser
}

func TestUserFunctions(t *testing.T) {
	u := server.User{
		ID:			"Some random id context from facebook",
		MaxCal:		2000,
		Timezone:	-8,
		Name:		"Random Name",
	}
	ser.AddUser(u)

	usr, err := ser.GetUser("Some random id context from facebook")
	if err != nil {
		t.Errorf("Errored: " + err.Error())
	}
	if (usr.ID != u.ID) {
		t.Errorf("Errored: usr.ID doesn't match u.ID")
	}
	if (usr.MaxCal != u.MaxCal) {
		t.Errorf("Errored: usr.MaxCal doesn't match u.MaxCal")
	}
	if (usr.Timezone != u.Timezone) {
		t.Errorf("Errored: usr.ID doesn't match u.ID")
	}
	if (usr.Name != u.Name) {
		t.Errorf("Errored: usr.Name doesn't match u.Name")
	}
}

func TestEntryFunctions(t *testing.T) {
	e := server.Entry{
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
	// Check entry if they match or not
	sum := 0
	for i := range *entries {
		entry := (*entries)[i]

		if (entry.ID != e.ID) {
			t.Errorf("Errored: entry.ID doesn't match e.ID")
		}
		if (entry.Time != e.Time) {
			t.Errorf("Errored: entry.Time doesn't match e.Time")
		}
		if (entry.Item != e.Item) {
			t.Errorf("Errored: entry.Item doesn't match e.Item")
		}
		if (entry.Calories != e.Calories) {
			t.Errorf("Errored: entry.Calories doesn't match e.Calories")
		}
		sum += entry.Calories
	}

	sum2, err := ser.SumCalories(e.ID)
	if (err != nil) {
		t.Errorf("Errored: " + err.Error())
	}
	if (sum2 != sum) {
		t.Errorf("Errored: SumCalories doesn't match actual sum")
	}
}
