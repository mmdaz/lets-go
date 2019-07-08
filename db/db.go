package db

import (
	"database/sql"
	"log"
)

var ConnectionPool *sql.DB

func InitDB() {
	connStr := "postgres://DBNAME:PASSWORD@localhost/url_db"
	var err error
	ConnectionPool, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
