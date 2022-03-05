package generation

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gabriel-ross/cs340-project/server/storage/model/generation"
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
	gg.POST("/", s.handleCreateGeneration)
	gg.PATCH("/:gid", s.handleUpdateGeneration)
	gg.DELETE("/:gid", s.handleDeleteGenerationByID)
}

func (s *Service) handleGetAllGenerations(c *gin.Context) {
	result, err := s.model.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (s *Service) handleCreateGeneration(c *gin.Context) {
	defer c.Request.Body.Close()
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	generation := &generation.Generation{}
	if err := json.Unmarshal(data, generation); err != nil {
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
	id := c.Param("gid")
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
	result, err := s.model.Update(generation)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *Service) handleDeleteGenerationByID(c *gin.Context) {
	id := c.Param("gid")
	err := s.model.DeleteByID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
