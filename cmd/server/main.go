package main

import (
	"go-bootcamp/api"
	"go-bootcamp/app"
	"go-bootcamp/pkg/pokemon"
	"log"
	"os"
	"path/filepath"
)

func main() {
	path, err := os.Getwd()
	parent := filepath.Dir(path)
	grandParent := filepath.Dir(parent)
	log.Println(err)
	log.Println(path)
	log.Println(grandParent)
	pokemonDAO := pokemon.NewPokemonDAO(grandParent + "/Pokemon.csv")

	pokemonService := pokemon.NewPokemonService(pokemonDAO)

	pokemonHandler := api.NewPokemonHandler(pokemonService)

	router := app.NewRouter(pokemonHandler)

	r := router.SetupRoutes()

	r.Run()
}
