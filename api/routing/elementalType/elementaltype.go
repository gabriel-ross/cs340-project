package elementalType

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gabriel-ross/cs340-project/server/storage/model/elementalType"
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
	tg.GET("/", s.handleGetTypes)
	tg.POST("/", s.handleCreateType)
	tg.PATCH("/:tid", s.handleUpdateType)
	tg.DELETE("/", s.handleDeleteTypeByName)
	tg.DELETE("/:tid", s.handleDeleteTypeByID)
}

func (s *Service) handleGetTypes(c *gin.Context) {
	if name := c.Query("name"); name != "" {
		result, err := s.model.FindByName(name)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, result)
		return
	}

	result, err := s.model.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

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

	result, err := s.model.Insert(elementType)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (s *Service) handleUpdateType(c *gin.Context) {
	id := c.Param("tid")
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
	result, err := s.model.Update(elementalType)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *Service) handleDeleteTypeByName(c *gin.Context) {
	if name := c.Query("name"); name != "" {
		err := s.model.DeleteByName(name)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.Status(http.StatusNoContent)
		return
	}
	c.Status(http.StatusBadRequest)
}

func (s *Service) handleDeleteTypeByID(c *gin.Context) {
	id := c.Param("tid")
	err := s.model.DeleteByID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
