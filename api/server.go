package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouterService interface {
	RegisterRoutes(*gin.RouterGroup)
}

type PokedexServer struct {
	router  *gin.Engine
	baseURL *gin.RouterGroup
	db      *sql.DB
}

func NewPokedexServer() *PokedexServer {
	newServer := PokedexServer{
		router: gin.New(),
	}
	newServer.baseURL = newServer.router.Group("")
	newServer.baseURL.GET("/", func(c *gin.Context) {
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

func (p *PokedexServer) RegisterRoutes(services ...RouterService) {
	for _, service := range services {
		service.RegisterRoutes(p.baseURL)
	}
}
