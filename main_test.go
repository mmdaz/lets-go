package main

import (
	"awesomeProject/db"
	"log"
	"os"
	"testing"
)

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS urls_test
(
    id INT  PRIMARY KEY,
    base_url VARCHAR(50) NOT NULL,
    new_url VARCHAR(50) NOT NULL
)`

func TestMain(m *testing.M) {
	// server.HandleRequests()
	db.InitDB()
	ensureTableExist()
	code := m.Run()
	clearTable()
	os.Exit(code)

}

func ensureTableExist() {
	if _, err := db.ConnectionPool.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	db.ConnectionPool.Exec("DELETE FROM urls_test")
	db.ConnectionPool.Exec("ALTER TABLE urls_test AUTO_INCREMENT = 1")
}
