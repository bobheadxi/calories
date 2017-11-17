package server

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"

	"github.com/bobheadxi/calories/config"
)

// Server : Contains the app's database and offers an
// interface to interact with it.
type Server struct {
	db *sql.DB
}

// New : Instantiate server
func New(cfg *config.EnvConfig) *Server {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		db: db,
	}
}

// AddUser : insert user into database
func (s *Server) AddUser(user User) error {
	sqlStatement := `
	INSERT INTO users (user_id, max_cal, timezone, name)
	VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(sqlStatement, user.ID, user.MaxCal, user.Timezone, user.Name)
	if err != nil {
		log.Print("AddUser failed: " + err.Error())
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
		log.Print("AddEntry failed: " + err.Error())
		return err
	}
	return nil
}

// GetUser : return a user from a Users table based on a given id
func (s *Server) GetUser(id string) (*User, error) {
	user := &User{}
	sqlStatement := `
	SELECT user_id, max_cal, timezone, name
	FROM users
	WHERE user_id = $1`
	row := s.db.QueryRow(sqlStatement, id)
	err := row.Scan(&user.ID, &user.MaxCal, &user.Timezone, user.Name)
	if err != nil {
		log.Print("GetUser failed: " + err.Error())
		return nil, err
	}
	return user, nil
}

// GetEntries : return a list of entries from a Users table
func (s *Server) GetEntries(id string) (*[]Entry, error) {
	l := []Entry{}
	sqlStatement := `
	SELECT fuser_id, item, time, calories
	FROM entries
	WHERE fuser_id = $1`
	rows, err := s.db.Query(sqlStatement, id)
	if err != nil {
		log.Print("GetEntries failed: " + err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		entry := Entry{}
		err := rows.Scan(&entry.ID, &entry.Item, &entry.Time, &entry.Calories)
		if err != nil {
			break
		}
		l = append(l, entry)
	}
	return &l, nil
}

// SumCalories : return sum of calories from entries for specific user
func (s *Server) SumCalories(id string) (int, error) {
	sqlStatement := `
	SELECT SUM(calories)
	FROM entries
	WHERE fuser_id = $1`
	rows, err := s.db.Query(sqlStatement, id)
	if err != nil {
		log.Print("SumCalories failed: " + err.Error())
		return 0, nil
	}
	var sum int
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&sum)
		if err != nil {
			log.Print("SumCalories failed: " + err.Error())
			return 0, nil
		}
	}
	return sum, nil
}

// CheckDB : return a boolean check on the database structure
func (s *Server) CheckDB() (bool) {
	sqlStatement := `
	SELECT column_name, data_type
	FROM information_schema.columns
	WHERE table_name = 'users'`
	rows, err := s.db.Query(sqlStatement)
	if err != nil {
		log.Print("CheckDB on Users failed: " + err.Error())
		return false
	}
	idx := 0
	var col string
	var typ string
	UsersSchema := [...]string{"user_id", "max_cal", "timezone", "name"}
	UsersSchemaType := [...]string{"bigint", "bigint", "integer", "text"}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&col, &typ)
		if err != nil {
			log.Print("CheckDB on Users failed: " + err.Error())
			return false
		}
		if (col != UsersSchema[idx] || typ != UsersSchemaType[idx]) {
			log.Print("CheckDB on Users failed: " + col + " " + typ)
			return false
		}
		idx++
	}


	// Entries
	sqlStatement = `
	SELECT column_name, data_type
	FROM information_schema.columns
	WHERE table_name = 'entries'`
	rows, err = s.db.Query(sqlStatement)
	if err != nil {
		log.Print("CheckDB on Entries failed: " + err.Error())
		return false
	}
	idx = 0
	// col & typ reused
	EntriesSchema := [...]string{"fuser_id", "time", "item", "calories"}
	EntriesSchemaType := [...]string{"bigint", "bigint", "text", "bigint"}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&col, &typ)
		if err != nil {
			log.Print("CheckDB on Entries failed: " + err.Error())
			return false
		}
		if (col != EntriesSchema[idx] || typ != EntriesSchemaType[idx]) {
			log.Print("CheckDB on Entries failed: " + col + " " + typ)
			return false
		}
		idx++
	}

	// No fatals
	log.Print("Successfully passed DB Schema check!")
	return true
}
