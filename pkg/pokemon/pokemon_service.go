package pokemon

import "github.com/pkg/errors"

type PokemonService struct {
	dao *PokemonDAO
}

func NewPokemonService(dao *PokemonDAO) *PokemonService {
	return &PokemonService{dao: dao}
}

func (s *PokemonService) GetPokemonByID(id int) (*Pokemon, error) {
	pokemon, err := s.dao.GetPokemonByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get pokemon by ID")
	}

	return pokemon, nil
}
