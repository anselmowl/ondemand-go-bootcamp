package service

import (
	"go-bootcamp/data"
	"go-bootcamp/model"
)

type PokemonService interface {
	GetPokemonByID(id int) (model.Pokemon, error)
	GetPokemonColor(id int) (model.PokemonColor, error)
}

type pokemonService struct {
	dao data.PokemonDAO
}

func NewPokemonService(dao data.PokemonDAO) PokemonService {
	return &pokemonService{dao: dao}
}

func (s *pokemonService) GetPokemonByID(id int) (model.Pokemon, error) {
	return s.dao.GetPokemonByID(id)
}

func (s *pokemonService) GetPokemonColor(id int) (model.PokemonColor, error) {
	return s.dao.GetPokemonColor(id)
}
