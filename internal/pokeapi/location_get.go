package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetLocation(location string) (PokemonEncounters, error) {
	url := baseURL + "/location-area/" + location

	if cached, ok := c.cache.Get(url); ok {
		var encountersResp PokemonEncounters
		err := json.Unmarshal(cached, &encountersResp)
		if err != nil {
			return PokemonEncounters{}, err
		}
		return encountersResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonEncounters{}, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return PokemonEncounters{}, errors.New("location not found")
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonEncounters{}, err
	}

	var encountersResp PokemonEncounters
	err = json.Unmarshal(data, &encountersResp)
	if err != nil {
		return PokemonEncounters{}, err
	}

	c.cache.Add(url, data)

	return encountersResp, nil
}
