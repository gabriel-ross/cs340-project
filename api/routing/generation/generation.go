package generation

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gabriel-ross/cs340-project/server/service/database/model/generation"
	"github.com/gin-gonic/gin"
)

type Service struct {
	model generation.Model
}

func NewService(model generation.Model) *Service {
	return &Service{model: model}
}

func (s *Service) RegisterRoutes(g *gin.RouterGroup) {
	gg := g.Group("/generations")
	gg.GET("/", s.handleGetAllGenerations)
	gg.POST("/:name", s.handleCreateGeneration)
	gg.PATCH("/:id", s.handleUpdateGeneration)
	gg.DELETE("/:name", s.handleDeleteGenerationByName)
}

func (s *Service) handleGetAllGenerations(c *gin.Context) {
	result, err := s.model.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// model/SQL should handle duplicates
func (s *Service) handleCreateGeneration(c *gin.Context) {
	defer c.Request.Body.Close()
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	generation := &generation.Generation{}
	err = json.Unmarshal(data, generation)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := s.model.Insert(generation)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (s *Service) handleUpdateGeneration(c *gin.Context) {
	id := c.Param("id")
	generation, err := s.model.FindByID(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	defer c.Request.Body.Close()
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(data, &generation)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	generation.Id = id
	result, err := s.model.UpdateByID(generation)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *Service) handleDeleteGenerationByName(c *gin.Context) {
	name := c.Param("name")
	err := s.model.DeleteByName(name)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
