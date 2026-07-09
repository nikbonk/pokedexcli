package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, input []string) error {
	if len(input) < 1 {
		return errors.New("no Pokémon name or ID provided")
	}

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(input[0])
	if err != nil {
		return err
	}

	pokemonBaseExperience := pokemonResp.BaseExperience

	fmt.Printf("Throwing a Pokeball at %v...", pokemonResp.Name)

	success := false
	catchValue := rand.IntN(pokemonBaseExperience)

	fmt.Println()

	if catchValue > 45 {
		success = true
	}

	if success {
		fmt.Println("You caught it!")
	} else {
		fmt.Println("It escaped!")
	}

	cfg.caughtPokemon[pokemonResp.Name] = pokemonResp

	return nil
}
