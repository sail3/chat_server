package utils

import (
	"database/sql"
	"fmt"
)

// Database pointer db coneccion
type Database struct {
	conn *sql.DB
}

// GetConnection retrieve a sql.DB pointer
func (db *Database) GetConnection() *sql.DB {
	if db.conn == nil {
		var err error
		db.conn, err = sql.Open("mysql", fmt.Sprintf(""))
		if err != nil {
			panic(err)
		}
	}
	return db.conn
}
