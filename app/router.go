package app

import (
	"go-bootcamp/api"

	"github.com/gin-gonic/gin"
)

type Router struct {
	pokemonHandler *api.PokemonHandler
}

func NewRouter(pokemonHandler *api.PokemonHandler) *Router {
	return &Router{pokemonHandler: pokemonHandler}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/pokemon", r.pokemonHandler.GetPokemonByIDHandler)

	return router
}
