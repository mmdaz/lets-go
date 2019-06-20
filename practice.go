package main

import (
	"awesomeProject/db"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type URL struct {
	ID      int
	BaseUrl string
	NewUrl  string
}

func main() {
	db.InitDB()
	db := db.ConnectionPool
	sqlStatement := `SELECT * FROM urls WHERE new_url=$1;`
	var url URL
	row := db.QueryRow(sqlStatement, "qwer")
	err := row.Scan(&url.ID, &url.BaseUrl, &url.NewUrl)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(url)
	default:
		panic(err)
	}
}
