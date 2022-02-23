package status

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) RegisterRoutes(g *gin.RouterGroup) {
	sg := g.Group("/status")
	sg.GET("/server", s.handleGetServerStatus)
	sg.GET("/database", s.handleGetDatabaseStatus)
}

func (s *Service) handleGetServerStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (s *Service) handleGetDatabaseStatus(c *gin.Context) {
	err := s.db.Ping()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
