package server

/*
This file defines the models used by our server's database.
See descriptions in the wiki: https://github.com/bobheadxi/calories/wiki/Schemas
*/

// User : Describes an entry in our "Users" table
type User struct {
	ID     string `json:"user_id"`
	MaxCal int    `json:"max_cal"`
}

// Entry : Describes an entry in our "Entries table"
type Entry struct {
	ID       string `json:"f_user_id"`
	Item     string `json:"item"`
	Time     string `json:"time"`
	Calories int    `json:"calories"`
}
