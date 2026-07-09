package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (RespPokemonInformation, error) {
	url := baseURL + "/pokemon/" + pokemon

	if cached, ok := c.cache.Get(url); ok {
		var pokemonResp RespPokemonInformation
		err := json.Unmarshal(cached, &pokemonResp)
		if err != nil {
			return RespPokemonInformation{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonInformation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonInformation{}, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return RespPokemonInformation{}, errors.New("pokemon not found")
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonInformation{}, err
	}

	var pokemonResp RespPokemonInformation
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return RespPokemonInformation{}, err
	}

	c.cache.Add(url, data)

	return pokemonResp, nil
}
