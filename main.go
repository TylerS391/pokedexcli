package main

import (
	"time"

	"github.com/TylerS391/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeApiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	caughtPokemon           map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeApiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
