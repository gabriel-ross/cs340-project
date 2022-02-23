package move

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gabriel-ross/cs340-project/server/service/database/model/move"
	"github.com/gin-gonic/gin"
)

type Service struct {
	model move.Model
}

func NewService(model move.Model) *Service {
	return &Service{model: model}
}

func (s *Service) RegisterRoutes(g *gin.RouterGroup) {
	g.GET("/moves", s.handleGetAllMoves)
	tg := g.Group("move")
	tg.POST("/:name", s.handleCreateMove)
	tg.PATCH("/:id", s.handleUpdateMove)
	tg.DELETE("/:name", s.handleDeleteMoveByName)
}

func (s *Service) handleGetAllMoves(c *gin.Context) {
	result, err := s.model.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// model/SQL should handle duplicates
func (s *Service) handleCreateMove(c *gin.Context) {
	defer c.Request.Body.Close()
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	move := &move.Move{}
	err = json.Unmarshal(data, move)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := s.model.Insert(move)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (s *Service) handleUpdateMove(c *gin.Context) {
	id := c.Param("id")
	move, err := s.model.FindByID(id)
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
	err = json.Unmarshal(data, &move)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	move.Id = id
	result, err := s.model.UpdateByID(move)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *Service) handleDeleteMoveByName(c *gin.Context) {
	name := c.Param("name")
	err := s.model.DeleteByName(name)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
