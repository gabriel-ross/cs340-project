package elementalType

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gabriel-ross/cs340-project/server/service/database/model/elementalType"
	"github.com/gin-gonic/gin"
)

// type referred to as ElementalType to avoid conflict with Go "type" keyword
type Service struct {
	model elementalType.Model
}

func NewService(model elementalType.Model) *Service {
	return &Service{model: model}
}

func (s *Service) RegisterRoutes(g *gin.RouterGroup) {
	tg := g.Group("/types")
	tg.GET("/", s.handleGetAllTypes)
	tg.GET("/:name", s.handleGetTypeIDByName)
	tg.POST("/:name", s.handleCreateType)
	tg.PATCH("/:id", s.handleUpdateType)
	tg.DELETE("/:name", s.handleDeleteTypeByName)
}

func (s *Service) handleGetAllTypes(c *gin.Context) {
	result, err := s.model.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (s *Service) handleGetTypeIDByName(c *gin.Context) {
	name := c.Param("name")
	result, err := s.model.FindByName(name)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// model/SQL should handle duplicates
func (s *Service) handleCreateType(c *gin.Context) {
	defer c.Request.Body.Close()
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	elementType := &elementalType.ElementalType{}
	err = json.Unmarshal(data, elementType)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if elementType.Id != c.Param("id") {
		c.AbortWithError(http.StatusBadRequest, errors.New(`url parameter "id" and body "id" do not match`))
		return
	}

	result, err := s.model.Insert(elementType)
	if err != nil {
		c.AbortWithError(http.StatusConflict, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (s *Service) handleUpdateType(c *gin.Context) {
	id := c.Param("id")
	elementalType, err := s.model.FindByID(id)
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
	err = json.Unmarshal(data, &elementalType)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	elementalType.Id = id
	result, err := s.model.UpdateByID(elementalType)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *Service) handleDeleteTypeByName(c *gin.Context) {
	name := c.Param("name")
	err := s.model.DeleteByName(name)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
