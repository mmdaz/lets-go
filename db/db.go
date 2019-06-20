package db

import (
	"database/sql"
	"log"
)

var ConnectionPool *sql.DB

func InitDB() {
	connStr := "postgres://muhammad:1540487768@localhost/url_db"
	var err error
	ConnectionPool, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
