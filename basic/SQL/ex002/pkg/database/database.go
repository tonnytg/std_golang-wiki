package database

import (
	"database/sql"
)

type Database struct {
	conn *db.SQL
}

func NewDatabase() (*Database, error) {

	var db Database

	db.conn, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}

	db.conn.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, email TEXT, password TEXT)")

	// save new field
	// db.conn.Exec("ALTER TABLE users ADD COLUMN password TEXT")

	return &db, nil
}
