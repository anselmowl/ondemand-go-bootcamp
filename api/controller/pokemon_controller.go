package controller

import (
	"strconv"

	"go-bootcamp/service"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type PokemonController interface {
	GetPokemonByID(c *gin.Context)
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
