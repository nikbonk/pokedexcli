package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, input []string) error {
	if len(input) < 1 {
		return errors.New("explore requires a location name")
	}

	encountersResp, err := cfg.pokeapiClient.GetLocation(input[0])
	if err != nil {
		return err
	}

	for _, encounter := range encountersResp.Encounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
