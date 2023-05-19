/*
Package model defines the structure of an object.
*/
package model // import "go-bootcamp/model"

// Pokemon is the model which represents a Pokemon object
type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PokemonColor is the model which represents a Pokemon object and its color
type PokemonColor struct {
	Pokemon Pokemon `json:"data"`
	Color   string  `json:"color"`
}
