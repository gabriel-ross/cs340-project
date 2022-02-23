package pokemon

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gabriel-ross/cs340-project/server/service/database/model/pokemon"
	"github.com/gin-gonic/gin"
)

type Service struct {
	model pokemon.Model
}

func NewService(model pokemon.Model) *Service {
	return &Service{model: model}
}

func (s *Service) RegisterRoutes(g *gin.RouterGroup) {
	pk := g.Group("/pokemon")
	pk.GET("/all", s.handleGetAllPokemon)
	pk.GET("", s.handleGetPokemon)
	pk.POST("", s.handleCreatePokemon)
	pk.PATCH("/:id", s.handleUpdatePokemonByID)
	pk.DELETE("/:id", s.handleDeletePokemonByID)
}

func (s *Service) handleGetAllPokemon(c *gin.Context) {
	result, err := s.model.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// if id is set search by id (as it is a primary key) otherwise search
// by variadic params
func (s *Service) handleGetPokemon(c *gin.Context) {
	// return all Pokemon if no query parameters passed
	if len(c.Request.URL.Query()) == 0 {
		s.handleGetAllPokemon(c)
		return
	}
	if id := c.Query("id"); id != "" {
		result, err := s.model.FindByID(id)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, result)

	} else {
		query := map[string]string{
			"name":       c.Query("name"),
			"type":       c.Query("type"),
			"generation": c.Query("generation"),
		}
		result, err := s.model.Find(query)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func (s *Service) handleCreatePokemon(c *gin.Context) {
	defer c.Request.Body.Close()
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pk := &pokemon.Pokemon{}
	err = json.Unmarshal(data, pk)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := s.model.Insert(pk)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (s *Service) handleUpdatePokemonByID(c *gin.Context) {
	id := c.Param("id")
	pokemon, err := s.model.FindByID(id)
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
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// prevent updating id via patch - must delete and post
	pokemon.Id = id
	result, err := s.model.UpdateByID(pokemon)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// TODO: implement PATCH handler for updating by name so we can alter id if need be

func (s *Service) handleDeletePokemonByID(c *gin.Context) {
	id := c.Param("id")
	err := s.model.DeleteByID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
