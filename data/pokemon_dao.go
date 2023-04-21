package data

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"go-bootcamp/model"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

type PokemonDAO interface {
	GetPokemonByID(id int) (*model.Pokemon, error)
	GetPokemonColor(id int) (*model.PokemonColor, error)
}

type pokemonDAO struct {
	filename string
}

func NewPokemonDAO(filename string) PokemonDAO {
	return &pokemonDAO{filename: filename}
}

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

func (dao *pokemonDAO) GetPokemonColor(id int) (*model.PokemonColor, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-species/%s", strconv.Itoa(id))
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get pokemon evolution")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get pokemon evolution: status code " + strconv.Itoa(resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response")
	}

	var data map[string]any
	json.Unmarshal([]byte(body), &data)

	color := data["color"].(map[string]any)

	pokemon := &model.Pokemon{
		ID:   id,
		Name: data["name"].(string),
	}

	pokemonColor := &model.PokemonColor{
		Pokemon: *pokemon,
		Color:   color["name"].(string),
	}

	return pokemonColor, nil
}
