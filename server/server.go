package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"awesomeProject/db"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	db.InitDB()
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/getURL/{baseURL}/{newURL}", getURL)
	myRouter.HandleFunc("/re/{newURL}", redirectURL)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func homePage(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getURL(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	baseURL := vars["baseURL"]
	newURL := vars["newURL"]
	insertToDb(baseURL, newURL)
	fmt.Fprintf(w, "Welcome to the GetUrlPage!")
}

func insertToDb(baseURL string, newURL string) {
	var id int
	err := db.ConnectionPool.QueryRow(`INSERT INTO urls(base_url, new_url) VALUES($1, $2) RETURNING id;`, baseURL, newURL).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
}

func redirectURL(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	newURL := vars["newURL"]
	sqlStatement := `SELECT base FROM urls WHERE new_urls=$1;`
	var baseURL string
	// rows, err := db.Query(`SELECT * FROM urls WHERE new_url = $1`, newURL)
	row := db.ConnectionPool.QueryRow(sqlStatement, newURL)
	err := row.Scan(&baseURL)
	if err != nil {
		log.Fatal(err)
	}
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(baseURL)
	default:
		panic(err)
	}
	http.Redirect(w, request, baseURL, 301)

}
