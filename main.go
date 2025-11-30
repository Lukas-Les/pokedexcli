package main

import (
	"github.com/Lukas-Les/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	Pokedex       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		Pokedex:       make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
