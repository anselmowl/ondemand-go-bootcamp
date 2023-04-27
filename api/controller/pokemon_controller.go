package controller

import (
	"go-bootcamp/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type PokemonController interface {
	GetPokemonByID(c *gin.Context)
	GetPokemonColor(c *gin.Context)
	GetPokemonsByIDRange(c *gin.Context)
}

type pokemonController struct {
	pokemonService service.PokemonService
}

func NewPokemonController(pokemonService service.PokemonService) PokemonController {
	return &pokemonController{pokemonService: pokemonService}
}

func (ctrl *pokemonController) GetPokemonByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.Wrap(err, "invalid ID").Error()})
	}

	pokemon, err := ctrl.pokemonService.GetPokemonByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": errors.Wrap(err, "unable to get the pokemon").Error()})
	}

	c.JSON(200, gin.H{"pokemon": pokemon})
}

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

func (ctrl *pokemonController) GetPokemonsByIDRange(c *gin.Context) {
	minIdParam := c.Query("min_id")
	maxIdParam := c.Query("max_id")
	workersParam := c.Query("workers")

	minId, err := strconv.Atoi(minIdParam)

	if err != nil {
		c.JSON(400, gin.H{"error": errors.Wrap(err, "invalid min ID").Error()})
	}

	maxId, err := strconv.Atoi(maxIdParam)

	if err != nil {
		c.JSON(400, gin.H{"error": errors.Wrap(err, "invalid max ID").Error()})
	}

	if minId > maxId {
		c.JSON(400, gin.H{"error": errors.Wrap(err, "min ID can't be greater than max ID").Error()})
	}

	workers, err := strconv.Atoi(workersParam)

	if err != nil {
		c.JSON(400, gin.H{"error": errors.Wrap(err, "workers should be a positive integer").Error()})
	}

	pokemons, err := ctrl.pokemonService.GetPokemonsByIDRange(minId, maxId, workers)
	if err != nil {
		c.JSON(404, gin.H{"error": errors.Wrap(err, "unable to get the pokemons by ID Range").Error()})
	}

	c.JSON(200, gin.H{"pokemons": pokemons})

}
