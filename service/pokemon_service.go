/*
Package service implements the mapping between the controller (the requests)
and the data (the data sources).
*/
package service // import "go-bootcamp/service"

import (
	"go-bootcamp/data"
	"go-bootcamp/model"
)

// PokemonService is the interface which wraps all the functions which map the requested data with the source of that data.
type PokemonService interface {
	GetPokemonByID(id int) (model.Pokemon, error)
	GetPokemonColor(id int) (model.PokemonColor, error)
}

// pokemonService represents the service for pokemon objects.
type pokemonService struct {
	dao data.PokemonDAO
}

// NewPokemonService creates a new instance of pokemonService with the given data instance.
func NewPokemonService(dao data.PokemonDAO) PokemonService {
	return &pokemonService{dao: dao}
}

// GetPokemonById returns a pokemon object with the given ID from the data layer by searching the data from a CSV file.
func (s *pokemonService) GetPokemonByID(id int) (model.Pokemon, error) {
	return s.dao.GetPokemonByID(id)
}

// GetPokemonById returns a pokemon object and its color with the given ID from the data layer by consuming an external API.
func (s *pokemonService) GetPokemonColor(id int) (model.PokemonColor, error) {
	return s.dao.GetPokemonColor(id)
}
