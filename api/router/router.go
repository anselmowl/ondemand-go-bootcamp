/*
Package router implements the HTTP router.
*/
package router // import "go-bootcamp/router"

import (
	"go-bootcamp/api/controller"

	"github.com/gin-gonic/gin"
)

// Sets up the HTTP router for the controllers functions
func InitRouter(pokemonController controller.PokemonController) *gin.Engine {
	router := gin.Default()

	pokemonGroup := router.Group("/pokemon")
	pokemonGroup.GET("/:id", pokemonController.GetPokemonByID)

	return router
}
