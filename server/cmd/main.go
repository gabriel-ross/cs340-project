package main

import (
	"github.com/gabriel-ross/cs340-project/server"
)

func main() {
	// instantiate server
	server := server.NewPokedexServer()
	// register services/routes
	// run server
	server.Run()
}
