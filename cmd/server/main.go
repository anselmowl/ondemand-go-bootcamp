package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"go-bootcamp/api/router"
	"go-bootcamp/data"
	"go-bootcamp/service"
)

func main() {
	// Get the the root directory (The SCV file is in that directory).
	workingDirectory, _ := os.Getwd()
	rootDirectory := filepath.Dir(filepath.Dir(workingDirectory))

	//Initialiize the pokemon service with the data layer instance.
	pokemonDAO := data.NewPokemonDAO(rootDirectory + "/Pokemon.csv")
	pokemonService := service.NewPokemonService(pokemonDAO)

	// Initilize the router using Gin-Gonic framework
	// the pokemonService instance is used as parameter to create an
	// instance of the pokemonController and add it to the router.
	router := router.InitRouter(pokemonService)

	// Run the server on the specified port
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error starting server: %s", err.Error()))
	}
}
