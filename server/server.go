package server

import (
	"database/sql"
	"log"

	"github.com/bobheadxi/calories/config"
)

// Server : Contains the app's database and offers an
// interface to interact with it.
type Server struct {
	db *sql.DB
}

// New : Instantiasafsldfalksdjf
func New(cfg *config.EnvConfig) *Server {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	return &Server{
		db: db,
	}
}

// InsertDataExample : Example use of database insertion
// DEPRECATE ASAP
func (s *Server) InsertDataExample(id string, content string) (int, error) {
	var userid int
	err := s.db.QueryRow(`INSERT INTO test(id, message) VALUES(` + id + `, '` + content + `') RETURNING id`).Scan(&userid)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return userid, nil
}

// AddUser : insert user into database
func (s *Server) AddUser(user User) error {
	sqlStatement := `  
	INSERT INTO users (user_id, max_cal)  
	VALUES ($1, $2)`
	_, err := s.db.Exec(sqlStatement, user.ID, user.MaxCal)
	if err != nil {
		log.Print("Error adding user: " + err.Error())
		return err
	}
	return nil
}

// AddEntry : add an entry to the database
func (s *Server) AddEntry(entry Entry) error {
	sqlStatement := `  
	INSERT INTO entries (fuser_id, time, item, calories)  
	VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(sqlStatement, entry.ID, entry.Time, entry.Item, entry.Calories)
	if err != nil {
		log.Print("Error adding entry: " + err.Error())
		return err
	}
	return nil
}

// GetUser : return a user from a Users table based on a given id
func (s *Server) GetUser(id string) (*User, error) {
	user := &User{}
	sqlStatement := `
	SELECT user_id, max_cal 
	FROM users
	WHERE user_id = $1`
	row := s.db.QueryRow(sqlStatement, id)
	err := row.Scan(&user.ID, &user.MaxCal)
	if err != nil {
		log.Print("Error finding a user: " + err.Error())
		return nil, err
	}
	return user, nil
}

// GetEntries: return a row of entries from a Users table
//   entries: fuser_id:bigserial
//  		  item: string (char[256])
//  		  time: bigint
//  		  calories: bigint
/*
func (s *Server) GetEntries(id string) (*Entry[], error) {

}
*/
