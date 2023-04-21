package main

import (
	"fmt"
	"go-bootcamp/api/router"
	"go-bootcamp/data"
	"go-bootcamp/service"
	"log"
	"os"
	"path/filepath"
)

func main() {
	workingDirectory, _ := os.Getwd()
	rootDirectory := filepath.Dir(filepath.Dir(workingDirectory))

	pokemonDAO := data.NewPokemonDAO(rootDirectory + "/Pokemon.csv")

	pokemonService := service.NewPokemonService(pokemonDAO)

	router := router.InitRouter(pokemonService)

	err := router.Run(":8080")

	if err != nil {
		log.Fatalf(fmt.Sprintf("Error starting server: %s", err.Error()))
	}
}
