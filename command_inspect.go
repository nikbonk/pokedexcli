package main

import "fmt"

func commandInspect(cfg *config, input []string) error {
	if len(input) != 1 {
		return fmt.Errorf("inspect requires a Pokémon name")
	}

	if _, ok := cfg.caughtPokemon[input[0]]; !ok {
		return fmt.Errorf("you have not caught this Pokémon")
	}

	fmt.Printf("Name: %s\n", input[0])
	fmt.Printf("Height: %d\n", cfg.caughtPokemon[input[0]].Height)
	fmt.Printf("Weight: %d\n", cfg.caughtPokemon[input[0]].Weight)
	fmt.Print("Stats:\n")
	for _, stat := range cfg.caughtPokemon[input[0]].Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Print("Types:\n")
	for _, pokemonType := range cfg.caughtPokemon[input[0]].Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}

	return nil
}
