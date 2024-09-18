package database

import (
	"database/sql"
	"os/user"

	_ "github/lib/pq"
)

var DB *sql.DB 

func OpenDatabase() error {
	var err error
	DB, err = sql.Open("postgres", user="quang263")
}
