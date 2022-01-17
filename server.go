package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const PORT = 8080

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world")
}

func main() {
	s := mux.NewRouter()
	s.HandleFunc("/", helloWorld).Methods("GET")
	fmt.Printf("Server listening on port %d", PORT)
	http.ListenAndServe(":"+strconv.Itoa(PORT), s)
}
