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

func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", pokemon).Methods("GET")
	r.HandleFunc("/types", types).Methods("GET")
	r.HandleFunc("/generations", generations).Methods("GET")
	r.HandleFunc("/moves", moves).Methods("GET")
	r.HandleFunc("/evolutions", evolutions).Methods("GET")
	r.HandleFunc("/db/test", dbDiagnostic).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))

	return r
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file, ", err)
	}

	r := newRouter()

	fmt.Printf("Server listening on port %d", PORT)
	http.ListenAndServe(":"+strconv.Itoa(PORT), r)
}

func pokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Pokemon")
}
func moves(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Moves")
}
func generations(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Generations")
}
func types(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Types")
}
func evolutions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Evolutions")
}
