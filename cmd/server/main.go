package main

import (
	"go-bootcamp/api"
	"go-bootcamp/app"
	"go-bootcamp/pkg/pokemon"
	"os"
	"path/filepath"
)

func main() {
	workingDirectory, _ := os.Getwd()
	rootDirectory := filepath.Dir(filepath.Dir(workingDirectory))

	pokemonDAO := pokemon.NewPokemonDAO(rootDirectory + "/Pokemon.csv")

	pokemonService := pokemon.NewPokemonService(pokemonDAO)

	pokemonHandler := api.NewPokemonHandler(pokemonService)

	router := app.NewRouter(pokemonHandler)

	r := router.SetupRoutes()

	r.Run()
}
