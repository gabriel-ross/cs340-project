package routing

import (
	"encoding/json"
	"github.com/gabriel-ross/cs340-project/server/service/database/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// type referred to as ElementalType to avoid conflict with Go "type" keyword
type ElementalTypeService struct {
	elementalType model.ElementalTypeModel
}

func (s *ElementalTypeService) RegisterRoutes(g *gin.RouterGroup) {
	g.GET("/types", s.handleGetAllTypes)
	tg := g.Group("type")
	tg.POST("/:name", s.handleCreateType)
	tg.PATCH("/:id", s.handleUpdateType)
	tg.DELETE("/:name", s.handleDeleteTypeByName)
}

func (s *ElementalTypeService) handleGetAllTypes(c *gin.Context) {
	result, err := s.elementalType.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// TODO: should preventing duplicate entries be done here or in the model/db
// and then the error propagated?
func (s *ElementalTypeService) handleCreateType(c *gin.Context) {
	//
}

func (s *ElementalTypeService) handleUpdateType(c *gin.Context) {
	id := c.Param("id")
	elementalType, err := s.elementalType.FindByID(id)
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
	result, err := s.elementalType.UpdateTypeByID(elementalType)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *ElementalTypeService) handleDeleteTypeByName(c *gin.Context) {
	name := c.Param("name")
	err := s.elementalType.DeleteTypeByName(name)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
