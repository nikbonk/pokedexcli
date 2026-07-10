package main

import "fmt"

func commandPokedex(c *config, input []string) error {
	if len(c.caughtPokemon) == 0 {
		fmt.Println("No Pokémon caught yet.")
		return nil
	}
	fmt.Println("Caught Pokémon:")
	for name := range c.caughtPokemon {
		fmt.Printf("- %s\n", name)
	}
	return nil
}
