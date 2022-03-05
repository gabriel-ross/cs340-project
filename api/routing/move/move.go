package move

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gabriel-ross/cs340-project/server/storage/model/move"
	"github.com/gin-gonic/gin"
)

type Service struct {
	model move.Model
}

func NewService(model move.Model) *Service {
	return &Service{model: model}
}

func (s *Service) RegisterRoutes(g *gin.RouterGroup) {
	mg := g.Group("/moves")
	mg.GET("/", s.handleGetAllMoves)
	mg.POST("/", s.handleCreateMove)
	mg.PATCH("/:mvid", s.handleUpdateMove)
	mg.DELETE("/:mvid", s.handleDeleteMoveByID)
}

// todo: implement filtering

func (s *Service) handleGetAllMoves(c *gin.Context) {
	result, err := s.model.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (s *Service) handleCreateMove(c *gin.Context) {
	defer c.Request.Body.Close()
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	move := &move.Move{}

	if err := json.Unmarshal(data, move); err != nil {
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
	id := c.Param("mvid")
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

	if err := json.Unmarshal(data, &move); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	move.Id = id

	result, err := s.model.Update(move)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *Service) handleDeleteMoveByID(c *gin.Context) {
	id := c.Param("mvid")
	err := s.model.DeleteByID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
