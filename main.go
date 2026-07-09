package main

import (
	"time"

	"github.com/nikbonk/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 10*time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.RespPokemonInformation{},
	}

	startRepl(cfg)
}
