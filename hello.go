package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/getURL/{baseURL}/{newURL}", getURL)
	myRouter.HandleFunc("/redirect/{newURL}", redirectURL)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
func main() {

	handleRequests()
}

func homePage(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func connectTODb() *sql.DB {
	connStr := "postgres://muhammad:1540487768@localhost/url_db"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
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
	db := connectTODb()
	err := db.QueryRow(`INSERT INTO urls(base_url, new_url) VALUES($1, $2) RETURNING id`, baseURL, newURL).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
}

func redirectURL(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	db := connectTODb()
	newURL := vars["newURL"]
	baseURL, err := db.Query(`SELECT * FROM urls WHERE new_url = $1`, newURL)
	if err != nil {
		log.Fatal(err)
	}
	var base string
	print(baseURL.Scan(&base))
	// http.Redirect(w, request, baseURL, http.StatusSeeOther)

}
