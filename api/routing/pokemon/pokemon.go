package pokemon

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gabriel-ross/cs340-project/server/service/database/model/pokemon"
	"github.com/gin-gonic/gin"
)

// todo: modify routes to take in name of type & generation
// todo: modify model to insert FK using subqueries
// maybe?

type Service struct {
	model pokemon.Model
}

func NewService(model pokemon.Model) *Service {
	return &Service{model: model}
}

func (s *Service) RegisterRoutes(g *gin.RouterGroup) {
	pg := g.Group("/pokemon")

	pg.GET("/", s.handleGetAllPokemon)
	pg.GET("/:pkid", s.handleGetPokemon)
	pg.POST("/:pkid", s.handleCreatePokemon)
	pg.PATCH("/:pkid", s.handleUpdatePokemonByID)
	pg.DELETE("/:pkid", s.handleDeletePokemonByID)

	pg.GET("/:pkid/moves", s.handleGetPokemonAllMoves)
	pg.POST("/:pkid/moves/:mvid", s.handlePokemonCreateMove)
	pg.DELETE("/:pkid/moves/:mvid", s.handlePokemonDeleteMove)
}

// todo: implement filtering by queries in this path
func (s *Service) handleGetAllPokemon(c *gin.Context) {
	result, err := s.model.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (s *Service) handleGetPokemon(c *gin.Context) {
	id := c.Param("pkid")
	result, err := s.model.FindByID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (s *Service) handleCreatePokemon(c *gin.Context) {
	defer c.Request.Body.Close()
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pk := &pokemon.Pokemon{}

	if err := json.Unmarshal(data, pk); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if pk.Id != c.Param("pkid") {
		c.AbortWithError(http.StatusBadRequest, errors.New(`url parameter "id" and body "id" do not match`))
		return
	}

	result, err := s.model.Insert(pk)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
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

	if err := json.Unmarshal(data, &pokemon); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if pokemon.Id != c.Param("pkid") {
		c.AbortWithError(http.StatusBadRequest, errors.New(`url parameter "id" and body "id" do not match`))
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

func (s *Service) handleGetPokemonAllMoves(c *gin.Context) {
	id := c.Param("pkid")
	result, err := s.model.FindAllMovesByPokemonID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (s *Service) handlePokemonCreateMove(c *gin.Context) {
	pkid := c.Param("pkid")
	mvid := c.Param("mvid")
	result, err := s.model.InsertPokemonMove(pkid, mvid)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (s *Service) handlePokemonDeleteMove(c *gin.Context) {
	pkid := c.Param("pkid")
	mvid := c.Param("mvid")
	result, err := s.model.DeletePokemonMove(pkid, mvid)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
