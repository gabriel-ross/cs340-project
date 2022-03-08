package pokemon

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gabriel-ross/cs340-project/server/storage/model/pokemon"
	"github.com/gin-gonic/gin"
)

type Service struct {
	model pokemon.Model
}

func NewService(model pokemon.Model) *Service {
	return &Service{model: model}
}

func (s *Service) RegisterRoutes(g *gin.RouterGroup) {
	pg := g.Group("/pokemon")

	pg.GET("/", s.handleGetManyPokemon)
	pg.GET("/:pkid", s.handleGetPokemon)
	pg.POST("/:pkid", s.handleCreatePokemon)
	pg.PATCH("/:pkid", s.handleUpdatePokemon)
	pg.DELETE("/:pkid", s.handleDeletePokemonByID)

	pg.GET("/moves", s.handleGetAllPokemonMoves)
	pg.GET("/:pkid/moves", s.handleGetPokemonAllMoves)
	pg.POST("/:pkid/moves/:mvid", s.handlePokemonCreateMove)
	pg.DELETE("/:pkid/moves/:mvid", s.handlePokemonDeleteMove)
}

// filter params will just be string names for now since that's much easier, though
// it isn't consistent with the rest of the project.
func (s *Service) handleGetManyPokemon(c *gin.Context) {
	if ok, filterBy := s.buildFilter([]string{"name", "type", "generation"}, c); ok {
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
		c.JSON(http.StatusAccepted, err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (s *Service) handleUpdatePokemon(c *gin.Context) {
	id := c.Param("pkid")
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

	pokemon.Id = id

	result, err := s.model.Update(pokemon)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *Service) handleDeletePokemonByID(c *gin.Context) {
	id := c.Param("pkid")
	err := s.model.DeleteByID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (s *Service) handleGetAllPokemonMoves(c *gin.Context) {
	result, err := s.model.FindAllPokemonMoves()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
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
