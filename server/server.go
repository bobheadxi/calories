package server

import (
	"database/sql"
	"log"

	"github.com/bobheadxi/calories/config"
	_ "github.com/lib/pq" // Postgres
)

// ServerLayer : Interface to interact with database
type ServerLayer interface {
	AddUser(User) error
	AddEntry(Entry) error
	GetUser(string) (*User, error)
	GetUsersInTimezone(int) (*map[*User]int, error)
	GetEntries(string) (*[]Entry, error)
	SumCalories(string) (int, error)
	UpdateUserTimezone(User) error
}

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

	server := &Server{
		db: db,
	}

	// Check integrity of database schema
	if !server.checkDatabaseIntegrity() {
		log.Fatal("Database formatted incorrectly.")
	}

	return server
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
	err := row.Scan(&user.ID, &user.MaxCal, &user.Timezone, &user.Name)
	if err != nil {
		log.Print("GetUser failed: " + err.Error())
		return nil, err
	}
	return user, nil
}

// GetUsersInTimezone : return all users in given timezone
// Returns a map with Users as keys and their total event calories as values
// Example usage:
// 	users, _ := b.server.GetUsersInTimezone(-8)
//  for u, v := range *users {
//		// do things
// 	}
func (s *Server) GetUsersInTimezone(tz int) (*map[*User]int, error) {
	users := make(map[*User]int)
	sqlStatement := `
	SELECT users.user_id, users.max_cal, users.name, users.timezone, SUM(entries.calories)
	FROM users JOIN entries ON users.user_id=entries.fuser_id
	WHERE timezone = $1
	GROUP BY users.user_id`
	rows, err := s.db.Query(sqlStatement, tz)
	if err != nil {
		log.Print("GetUsersInTimezone failed: " + err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var cVal int
		user := User{}
		err := rows.Scan(&user.ID, &user.MaxCal, &user.Name, &user.Timezone, &cVal)
		if err != nil {
			log.Print("Row scan failed in GetUsersInTimezone: " + err.Error())
			break
		}
		users[&user] = cVal
	}
	return &users, nil
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
			log.Print("Row scan failed in GetEntries: " + err.Error())
			break
		}
		l = append(l, entry)
	}
	return &l, nil
}

// SumCalories : return sum of calories from entries for specific user
// Example usage:
// 	calories, _ := b.server.SumCalories(c.facebookID)
//	log.Print(strconv.Itoa(calories))
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

// UpdateUserTimezone updates the timezone value of the given user
func (s *Server) UpdateUserTimezone(user User) error {
	sqlStatement := `
	UPDATE users
   	SET timezone = $2
	WHERE user_id = $1`
	_, err := s.db.Exec(sqlStatement, user.ID, user.Timezone)
	if err != nil {
		log.Print("AddUser failed: " + err.Error())
		return err
	}
	return nil
}
