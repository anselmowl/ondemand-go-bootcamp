package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-bootcamp/model"
	"go-bootcamp/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPokemonDAO struct which creates a Mock of a Pokemon Data Access Object
type MockPokemonDAO struct {
	mock.Mock
}

// GetPokemonByID represents the data object that will be mocked for the data.PokemonDAO.GetPokemonByID function test
func (m *MockPokemonDAO) GetPokemonByID(id int) (model.Pokemon, error) {
	args := m.Called(id)
	return args.Get(0).(model.Pokemon), args.Error(1)
}

// GetPokemonColor represents the data object that will be mocked for the data.PokemonDAO.GetPokemonColor function test
func (m *MockPokemonDAO) GetPokemonColor(id int) (model.PokemonColor, error) {
	args := m.Called(id)
	return args.Get(0).(model.PokemonColor), args.Error(1)
}

// TestPokemonController_GetPokemonColor tests a succesful request to the Controller GetPokemonColor function
func TestPokemonController_GetPokemonColor(t *testing.T) {
	mockDAO := &MockPokemonDAO{}

	pokemon := model.Pokemon{
		ID:   1,
		Name: "Bulbasaur",
	}
	pokemonColor := model.PokemonColor{
		Pokemon: pokemon,
		Color:   "Green",
	}

	pokeMock, _ := json.Marshal(pokemonColor)

	mockDAO.On("GetPokemonColor", 1).Return(pokemonColor, nil)

	srv := service.NewPokemonService(mockDAO)
	ctrl := NewPokemonController(srv)

	router := gin.Default()
	router.GET("/pokemon/color/:id", ctrl.GetPokemonColor)

	req, err := http.NewRequest("GET", "/pokemon/color/1", nil)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.NoError(t, err)
	mockDAO.AssertCalled(t, "GetPokemonColor", 1)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), string(pokeMock))
}

// TestPokemonController_GetPokemonByID tests a succesful request to the Controller GetPokemonById function
func TestPokemonController_GetPokemonByID(t *testing.T) {
	mockDAO := &MockPokemonDAO{}

	pokemon := model.Pokemon{
		ID:   1,
		Name: "Bulbasaur",
	}

	pokeMock, _ := json.Marshal(pokemon)

	mockDAO.On("GetPokemonByID", 1).Return(pokemon, nil)

	srv := service.NewPokemonService(mockDAO)
	ctrl := NewPokemonController(srv)

	router := gin.Default()
	router.GET("/pokemon/:id", ctrl.GetPokemonByID)

	req, err := http.NewRequest("GET", "/pokemon/1", nil)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.NoError(t, err)
	mockDAO.AssertCalled(t, "GetPokemonByID", 1)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), string(pokeMock))
}
