package main

import (
	"time"

	"github.com/Lukas-Les/pokedexcli/internal/pokeapi"
	"github.com/Lukas-Les/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient pokeapi.Client
	cache         pokecache.Cache
	Next          *string
	Previous      *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		cache:         pokecache.NewCache(time.Second * 5),
	}
	startRepl(&cfg)
}
