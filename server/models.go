package server

/*
This file defines the models used by our server's database.
See descriptions in the wiki: https://github.com/bobheadxi/calories/wiki/Schemas
*/

// User : Describes an entry in our "Users" table
type User struct {
	ID       string `json:"user_id"`
	MaxCal   int    `json:"max_cal"`
	Timezone int    `json:"timezone"`
	Name     string `json:"name"`
}

// Entry : Describes an entry in our "Entries" table
type Entry struct {
	ID       string `json:"fuser_id"`
	Time     int64  `json:"time"`
	Item     string `json:"item"`
	Calories int    `json:"calories"`
}
