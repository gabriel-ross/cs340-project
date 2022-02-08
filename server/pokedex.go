package server

import (
	"database/sql"
	"log"

	"github.com/gabriel-ross/cs340-project/server/service/database/mariadb"
	"github.com/gin-gonic/gin"
)

type RouterService interface {
	RegisterRoutes(*gin.Engine)
}

type PokedexServer struct {
	router *gin.Engine
	db     *sql.DB
}

func NewPokedexServer() *PokedexServer {
	newServer := PokedexServer{
		router: gin.New(),
	}
	return &newServer
}

func (p *PokedexServer) Use(middlewares ...gin.HandlerFunc) {
	for _, mw := range middlewares {
		p.router.Use(mw)
	}
}

func (p *PokedexServer) Run() {
	if p.db != nil {
		defer p.db.Close()
	}
	log.Fatal(p.router.Run())
}

func (p *PokedexServer) RegisterRoutes(services ...RouterService) {
	for _, service := range services {
		service.RegisterRoutes(p.router)
	}
}

// TODO: save credentials in order to restart db?
func (p *PokedexServer) RegisterMariaDBService(username, password, host, port, dbName string) {
	db, err := mariadb.Connect(username, password, host, port, dbName)
	if err != nil {
		log.Fatal(err)
	}
	p.db = db
}
