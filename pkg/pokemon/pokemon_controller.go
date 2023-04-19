package pokemon

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type PokemonController struct {
	dao *PokemonDAO
}

func NewPokemonController(dao *PokemonDAO) *PokemonController {
	return &PokemonController{dao: dao}
}

func (ctrl *PokemonController) GetPokemonByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.Wrap(err, "invalid ID").Error()})
	}

	pokemon, err := ctrl.dao.GetPokemonByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": errors.Wrap(err, "unable to get the pokemon").Error()})
	}

	c.JSON(200, gin.H{"pokemon": pokemon})
}
