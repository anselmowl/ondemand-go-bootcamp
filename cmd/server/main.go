package main

import (
	"context"

	"go-bootcamp/api/controller"
	"go-bootcamp/api/router"
	"go-bootcamp/data"
	"go-bootcamp/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	// Provide all the dependencies
	app := fx.New(
		fx.Provide(
			data.NewPokemonDAO,
			service.NewPokemonService,
			controller.NewPokemonController,
			router.InitRouter,
		),
		fx.Invoke(NewHTTPServer),
	)

	app.Run()
}

func NewHTTPServer(lc fx.Lifecycle, ge *gin.Engine) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Run the server on the specified port
				go ge.Run(":8080")
				return nil
			},
		},
	)
}
