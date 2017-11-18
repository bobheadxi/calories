package server

import (
	"log"
)

// CheckDB : return a boolean check on the database structure
func (s *Server) checkDatabaseIntegrity() bool {
	sqlStatement := `
	SELECT column_name, data_type
	FROM information_schema.columns
	WHERE table_name = 'users'`
	rows, err := s.db.Query(sqlStatement)
	if err != nil {
		log.Print("CheckDB on Users failed: " + err.Error())
		return false
	}

	var col string
	var typ string

	UsersSchema := []string{"user_id", "max_cal", "timezone", "name"}
	UsersSchemaType := []string{"bigint", "bigint", "integer", "text"}

	idx := 0
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

	EntriesSchema := []string{"fuser_id", "time", "item", "calories"}
	EntriesSchemaType := []string{"bigint", "bigint", "text", "bigint"}

	idx = 0
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
