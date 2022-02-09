package routing

import (
	"encoding/json"
	"github.com/gabriel-ross/cs340-project/server/service/database/model"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PokemonService struct {
	pokemon model.PokemonModel
}

func (s *PokemonService) RegisterRoutes(g *gin.RouterGroup) {
	pk := g.Group("/pokemon")
	pk.GET("/all", s.handleGetAllPokemon)
	pk.GET("/", s.handleGetPokemon)
	pk.POST("/", s.handleCreatePokemon)
	pk.PATCH("/:id", s.handleUpdatePokemonByID)
	pk.DELETE("/:id", s.handleDeletePokemonByID)
}

func (s *PokemonService) handleGetAllPokemon(c *gin.Context) {
	result, err := s.pokemon.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// if id is set search by id (as it is a primary key) otherwise search
// by variadic params]
func (s *PokemonService) handleGetPokemon(c *gin.Context) {
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

func (s *PokemonService) handleCreatePokemon(c *gin.Context) {
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

func (s *PokemonService) handleUpdatePokemonByID(c *gin.Context) {
	id := c.Param("id")
	pokemon, err := s.pokemon.FindByID(id)
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
	pokemon.NDexId = id
	result, err := s.pokemon.UpdatePokemonByID(pokemon)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// TODO: implement PATCH handler for updating by name so we can alter id if need be

func (s *PokemonService) handleDeletePokemonByID(c *gin.Context) {
	id := c.Param("id")
	err := s.pokemon.DeletePokemonByID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
