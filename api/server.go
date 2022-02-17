package server

import (
	"database/sql"
	"log"
	"net/http"

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
	newServer.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello, world")
	})
	return &newServer
}

func (p *PokedexServer) Use(middlewares ...gin.HandlerFunc) {
	for _, mw := range middlewares {
		p.router.Use(mw)
	}
}

func (p *PokedexServer) Run(port string) {
	if p.db != nil {
		defer p.db.Close()
	}
	log.Fatal(p.router.Run(port))
}

func (p *PokedexServer) RegisterDB(db *sql.DB) {
	p.db = db
}

func (p *PokedexServer) RegisterMariaDB(config mariadb.Config) error {
	db, err := mariadb.ConnectWithConfig(config)
	if err != nil {
		return err
	}
	p.RegisterDB(db)
	return nil
}

func (p *PokedexServer) RegisterRoutes(services ...RouterService) {
	for _, service := range services {
		service.RegisterRoutes(p.router)
	}
}
