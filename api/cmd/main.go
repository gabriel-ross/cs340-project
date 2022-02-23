package main

import (
	"os"

	"github.com/gabriel-ross/cs340-project/server"
	"github.com/gabriel-ross/cs340-project/server/routing/elementalType"
	"github.com/gabriel-ross/cs340-project/server/routing/pokemon"
	"github.com/gabriel-ross/cs340-project/server/routing/status"
	"github.com/gabriel-ross/cs340-project/server/service/database/mariadb"
	eTypeModel "github.com/gabriel-ross/cs340-project/server/service/database/model/pokemon"
	generationModel "github.com/gabriel-ross/cs340-project/server/service/database/model/pokemon"
	moveModel "github.com/gabriel-ross/cs340-project/server/service/database/model/pokemon"
	pokemonModel "github.com/gabriel-ross/cs340-project/server/service/database/model/pokemon"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// instantiate server
	server := server.NewPokedexServer()
	db, err := mariadb.Connect(
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
	)
	if err != nil {
		panic(err)
	}
	server.RegisterDB(db)
	// instantiate routing services
	statusRoutes := status.NewService(db)
	pokemonRoutes := pokemon.NewService(*pokemonModel.NewModel(db))
	typeRoutes := elementalType.NewService(*eTypeModel.NewModel(db))
	generationRoutes := elementalType.NewService(*generationModel.NewModel(db))
	moveRoutes := elementalType.NewService(*moveModel.NewModel(db))
	// register services/routes
	server.RegisterRoutes(statusRoutes, pokemonRoutes, typeRoutes, generationRoutes, moveRoutes)
	// run server
	server.Run(":80")
}
