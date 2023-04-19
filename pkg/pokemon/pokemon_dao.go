package pokemon

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

type PokemonDAO struct {
	filename string
}

func NewPokemonDAO(filename string) *PokemonDAO {
	return &PokemonDAO{filename: filename}
}

func (dao *PokemonDAO) GetPokemonByID(id int) (*Pokemon, error) {
	// Open CSV file
	f, err := os.Open(dao.filename)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open CSV")
	}
	defer f.Close()

	// read data from CSV
	reader := csv.NewReader(f)
	pokemons, err := reader.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "unable to read CSV")
	}

	// search the pokemon by id
	for _, pkm := range pokemons {
		if (pkm[0]) == strconv.Itoa(id) {
			pokemon := &Pokemon{
				ID:   id,
				Name: pkm[1],
			}
			return pokemon, nil
		}
	}
	return nil, errors.New("pokemon not found")
}
