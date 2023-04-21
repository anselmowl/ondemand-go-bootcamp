package model

type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PokemonColor struct {
	Pokemon Pokemon `json:"data"`
	Color   string  `json:"color"`
}
