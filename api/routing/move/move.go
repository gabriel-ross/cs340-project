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
	mg.GET("/", s.handleGetManyMoves)
	mg.POST("/", s.handleCreateMove)
	mg.PATCH("/:mvid", s.handleUpdateMove)
	mg.DELETE("/:mvid", s.handleDeleteMoveByID)
}

func (s *Service) handleGetManyMoves(c *gin.Context) {
	if ok, filterBy := s.buildFilter([]string{"type"}, c); ok {
		result, err := s.model.Find(filterBy)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, result)
	} else {
		result, err := s.model.FindAll()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func (s *Service) buildFilter(params []string, c *gin.Context) (bool, map[string]string) {
	filterParamsPresent := false
	filters := map[string]string{}
	for _, param := range params {
		val := c.Query(param)
		if val != "" {
			filterParamsPresent = true
			filters[param] = val
		}
	}
	return filterParamsPresent, filters
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
