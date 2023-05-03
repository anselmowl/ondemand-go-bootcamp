/*
Package data implements the logic to retrive structured objects from data sources.
*/
package data // import "go-bootcamp/data"

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"

	"go-bootcamp/model"

	"github.com/pkg/errors"
)

// PokemonDAO is the interface which wraps all the functions with the logic to retrive Data Access Objects related to pokemon resources
type PokemonDAO interface {
	GetPokemonByID(id int) (*model.Pokemon, error)
}

// pokemonDAO is a struct that defines a Data Access Object for pokemons
type pokemonDAO struct {
	filename string // The CSV file path where pokemon data is stored
}

// NewPokemonDAO cretes a new instance of pokemonDAO
func NewPokemonDAO() PokemonDAO {
	return &pokemonDAO{filename: getFilePath()}
}

// GetPokemonByID returns a Pokemon object with the fiven ID from the CSV file
func (dao *pokemonDAO) GetPokemonByID(id int) (*model.Pokemon, error) {
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
			pokemon := &model.Pokemon{
				ID:   id,
				Name: pkm[1],
			}
			return pokemon, nil
		}
	}
	return nil, errors.New("pokemon not found")
}

// Get the the file path (The SCV file is at root directory).
func getFilePath() (rootPath string) {
	workingDirectory, _ := os.Getwd()
	rootDirectory := filepath.Dir(filepath.Dir(workingDirectory))
	return rootDirectory + "/Pokemon.csv"
}
