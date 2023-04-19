package api

import (
	"go-bootcamp/pkg/pokemon"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PokemonHandler struct {
	pokemonService pokemon.PokemonService
}

func NewPokemonHandler(pokemonService pokemon.PokemonService) *PokemonHandler {
	return &PokemonHandler{pokemonService: pokemonService}
}

func (h *PokemonHandler) GetPokemonByIDHandler(c *gin.Context) {
	pokemonIDParam := c.Query("id")
	pokemonID, err := strconv.Atoi(pokemonIDParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pokemon, err := h.pokemonService.GetPokemonByID(pokemonID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   pokemon.ID,
		"name": pokemon.Name,
	})
}
