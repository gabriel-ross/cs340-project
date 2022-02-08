package routing

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gabriel-ross/cs340-project/server/service/model"
	"github.com/gin-gonic/gin"
)

type Service struct {
	pokemon model.PokemonModel
}

func NewService(pokemon model.PokemonModel) *Service {
	return &Service{
		pokemon: pokemon,
	}
}

func (s *Service) RegisterRoutes(g *gin.RouterGroup) {
	pk := g.Group("/pokemon")
	pk.GET("/all", s.handleGetAllPokemon)
	pk.GET("/", s.handleGetPokemon)
	pk.POST("/", s.handleCreatePokemon)
	pk.DELETE("/:id", s.handleDeletePokemonByID)
}

func (s *Service) handleGetAllPokemon(c *gin.Context) {
	result, err := s.pokemon.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// if id is set search by id (as it is a primary key) otherwise search
// by variadic params]
func (s *Service) handleGetPokemon(c *gin.Context) {
	// return all Pokemon if no query parameters passed
	if len(c.Request.URL.Query()) == 0 {
		s.handleGetAllPokemon(c)
		return
	}
	if id := c.Query("id"); id != "" {
		result, err := s.pokemon.FindByID(id)
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
		result, err := s.pokemon.Find(query)
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

	pk := &model.Pokemon{}
	err = json.Unmarshal(data, pk)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := s.pokemon.InsertPokemon(pk)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (s *Service) handleDeletePokemonByID(c *gin.Context) {
	id := c.Param("id")
	err := s.pokemon.DeletePokemonByID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.Status(http.StatusNoContent)
}
