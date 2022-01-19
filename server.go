package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const PORT = 8080

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world")
}

func dbDiagnostic(w http.ResponseWriter, r *http.Request) {

	var (
		uName    = os.Getenv("MARIADB_USERNAME")
		password = os.Getenv("MARIADB_PASSWORD")
		hostName = os.Getenv("DB_HOST_NAME")
		dbName   = os.Getenv("DB_NAME")
	)

	connInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s", uName, password, hostName, dbName)
	db, err := sql.Open("mysql", connInfo)
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

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file, ", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/", helloWorld).Methods("GET")
	router.HandleFunc("/db/test", dbDiagnostic).Methods("GET")
	router.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))

	fmt.Printf("Server listening on port %d", PORT)
	http.ListenAndServe(":"+strconv.Itoa(PORT), router)
}
