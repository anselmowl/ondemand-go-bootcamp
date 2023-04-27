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
	"sync"

	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type PokemonDAO interface {
	GetPokemonByID(id int) (model.Pokemon, error)
	GetPokemonColor(id int) (model.PokemonColor, error)
	GetPokemonsByIDRange(minId, maxId, workers int) ([]model.Pokemon, error)
}

type pokemonDAO struct {
	filename string
}

func NewPokemonDAO(filename string) PokemonDAO {
	return &pokemonDAO{filename: filename}
}

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

func (dao *pokemonDAO) GetPokemonsByIDRange(minId, maxId, workers int) ([]model.Pokemon, error) {
	// var pokemons []model.Pokemon

	// idChannel := make(chan int)
	// var wg sync.WaitGroup

	// for i := 0; i < workers; i++ {
	// 	wg.Add(1)
	// 	go dao.GetPokemonsByIDRangeWorker(idChannel, &pokemons, &wg)
	// }

	// for i := minId; i <= maxId; i++ {
	// 	idChannel <- i
	// }

	// close(idChannel)

	// wg.Wait()

	// return pokemons, nil

	idCahannel := make(chan int, maxId-minId+1)
	defer close(idCahannel)

	for i := minId; i <= maxId; i++ {
		idCahannel <- i
	}

	wg := sync.WaitGroup{}

	pool := make(chan struct{}, workers)

	resultChannel := make(chan model.Pokemon, maxId-minId+1)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go dao.GetPokemonsByIDRangeWorker(idCahannel, pool, resultChannel, &wg)
	}

	wg.Wait()

	close(resultChannel)

	pokemons := make([]model.Pokemon, 0, maxId-minId+1)
	for pokemon := range resultChannel {
		pokemons = append(pokemons, pokemon)
	}
	return pokemons, nil
}

func (dao *pokemonDAO) GetPokemonsByIDRangeWorker(idChannel chan int, pool chan struct{}, resultChannel chan model.Pokemon, wg *sync.WaitGroup) {
	// defer wg.Done()
	// for id := range idChannel {
	// 	pokemon, err := dao.GetPokemonByID(id)
	// 	if err == nil {
	// 		*pokemons = append(*pokemons, pokemon)
	// 	}
	// }
	defer wg.Done()

	pool <- struct{}{}
	defer func() {
		<-pool
	}()

	id, ok := <-idChannel
	if !ok {
		return
	}

	pokemon, err := dao.GetPokemonByID(id)
	if err != nil {
		errors.Wrap(err, "unable to get pokemon with ID: "+strconv.Itoa(id))
		return
	}

	resultChannel <- pokemon

}
