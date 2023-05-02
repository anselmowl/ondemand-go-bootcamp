/*
Package router implements the HTTP router.
*/
package router // import "go-bootcamp/router"

import (
	"go-bootcamp/api/controller"
	"go-bootcamp/service"

	"github.com/gin-gonic/gin"
)

// Sets up the HTTP router for the controllers functions
func InitRouter(pokemonService service.PokemonService) *gin.Engine {
	router := gin.Default()

	pokemonGroup := router.Group("/pokemon")
	pokemonController := controller.NewPokemonController(pokemonService)
	pokemonGroup.GET("/:id", pokemonController.GetPokemonByID)

	return router
}
