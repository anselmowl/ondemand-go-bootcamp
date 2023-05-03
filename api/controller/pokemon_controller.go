/*
Package controller implements the procesing of an user request to build an
appropriate model as response.
*/
package controller // import "go-bootcamp/controller"

import (
	"net/http"
	"strconv"

	"go-bootcamp/service"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// PokemonController is the interface which wraps all the functions which handles the HTTP requests related to pokemon resources
type PokemonController interface {
	GetPokemonByID(c *gin.Context)
	GetPokemonColor(c *gin.Context)
}

// pokemonController handles HTTP requests related to pokemon resources
type pokemonController struct {
	pokemonService service.PokemonService
}

// NewPokemonController creates a new instance of pokemonController
func NewPokemonController(pokemonService service.PokemonService) PokemonController {
	return &pokemonController{pokemonService: pokemonService}
}

// GetPokemonByID gets a Pokemon object that match its pokedex number with the given ID
// the given Id should be a positive integer in the range of 1 to 151.
func (ctrl *pokemonController) GetPokemonByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.Wrap(err, "invalid ID").Error()})
	}

	pokemon, err := ctrl.pokemonService.GetPokemonByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.Wrap(err, "unable to get the pokemon").Error()})
	}

	c.JSON(http.StatusOK, gin.H{"pokemon": pokemon})
}

// GetPokemonColor gets a Pokemon object and its color that match its pokedex number with the given ID
// the given Id should be a positive integer in the range of 1 to 151.
func (ctrl *pokemonController) GetPokemonColor(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.Wrap(err, "invalid ID").Error()})
	}

	pokemon, err := ctrl.pokemonService.GetPokemonColor(id)
	if err != nil {
		c.JSON(404, gin.H{"error": errors.Wrap(err, "unable to get the pokemon").Error()})
	}

	c.JSON(200, gin.H{"pokemon": pokemon})
}
