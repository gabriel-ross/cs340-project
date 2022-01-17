package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const PORT = 8080

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world")
}

func dbDiagnostic(w http.ResponseWriter, r *http.Request) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL_MODE"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, "Successfully connected to db")

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", helloWorld).Methods("GET")
	router.HandleFunc("/db/test", dbDiagnostic).Methods("GET")
	router.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))

	fmt.Printf("Server listening on port %d", PORT)
	http.ListenAndServe(":"+strconv.Itoa(PORT), router)
}
