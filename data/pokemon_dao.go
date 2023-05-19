/*
Package data implements the logic to retrive structured objects from data sources.
*/
package data // import "go-bootcamp/data"

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"go-bootcamp/model"

	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// PokemonDAO is the interface which wraps all the functions with the logic to retrive Data Access Objects related to pokemon resources
type PokemonDAO interface {
	GetPokemonByID(id int) (model.Pokemon, error)
	GetPokemonColor(id int) (model.PokemonColor, error)
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
func (dao *pokemonDAO) GetPokemonByID(id int) (model.Pokemon, error) {
	// Open CSV file
	f, err := os.Open(dao.filename)
	if err != nil {
		return model.Pokemon{}, errors.Wrap(err, "unable to open CSV")
	}
	defer f.Close()

	// read data from CSV
	reader := csv.NewReader(f)
	pokemons, err := reader.ReadAll()
	if err != nil {
		return model.Pokemon{}, errors.Wrap(err, "unable to read CSV")
	}

	// search the pokemon by id
	for _, pkm := range pokemons {
		if (pkm[0]) == strconv.Itoa(id) {
			pokemon := model.Pokemon{
				ID:   id,
				Name: pkm[1],
			}
			return pokemon, nil
		}
	}
	return model.Pokemon{}, errors.New("pokemon not found")
}

// GetPokemonColor returns a Pokemon object and its color with the given ID from an external API
func (dao *pokemonDAO) GetPokemonColor(id int) (model.PokemonColor, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-species/%s", strconv.Itoa(id))
	resp, err := http.Get(url)
	if err != nil {
		return model.PokemonColor{}, errors.Wrap(err, "unable to get pokemon color")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.PokemonColor{}, errors.New("failed to get pokemon color: status code " + strconv.Itoa(resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.PokemonColor{}, errors.Wrap(err, "unable to read response")
	}

	var data map[string]any
	json.Unmarshal([]byte(body), &data)

	caser := cases.Title(language.Und)

	color := data["color"].(map[string]any)

	pokemon := model.Pokemon{
		ID:   id,
		Name: caser.String(data["name"].(string)),
	}

	pokemonColor := model.PokemonColor{
		Pokemon: pokemon,
		Color:   caser.String(color["name"].(string)),
	}

	csvFile, _ := os.OpenFile(dao.filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	record := []string{
		strconv.Itoa(pokemon.ID),
		pokemon.Name,
		pokemonColor.Color,
	}

	writer.Write(record)

	return pokemonColor, nil
}

// Get the the file path (The SCV file is at root directory).
func getFilePath() (rootPath string) {
	workingDirectory, _ := os.Getwd()
	rootDirectory := filepath.Dir(filepath.Dir(workingDirectory))
	return rootDirectory + "/Pokemon.csv"
}
