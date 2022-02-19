package elementalType

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gabriel-ross/cs340-project/server/service/database/model"
	"github.com/gin-gonic/gin"
)

// TODO: unexport the Service struct

// type referred to as ElementalType to avoid conflict with Go "type" keyword
type Service struct {
	model model.ElementalTypeModel
}

func NewService(model model.ElementalTypeModel) *Service {
	return &Service{model: model}
}

func (s *Service) RegisterRoutes(g *gin.RouterGroup) {
	g.GET("/types", s.handleGetAllTypes)
	tg := g.Group("type")
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

// model/SQL should handle duplicates
func (s *Service) handleCreateType(c *gin.Context) {
	defer c.Request.Body.Close()
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	elementType := &model.ElementalType{}
	err = json.Unmarshal(data, elementType)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := s.model.InsertType(elementType)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
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
	result, err := s.model.UpdateTypeByID(elementalType)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *Service) handleDeleteTypeByName(c *gin.Context) {
	name := c.Param("name")
	err := s.model.DeleteTypeByName(name)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
