package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL

	if name != "" {
		url += fmt.Sprintf("/pokemon/%s/", name)
	} else {
		return Pokemon{}, fmt.Errorf("Usage: catch <pokemon>")
	}

	//Check cache before making request
	data, ok := c.cache.Get(url)
	if ok {
		log.Println("CACHE HIT:", url)
		var pokemon Pokemon

		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, fmt.Errorf("failed fetching pokemon: %w", err)
		}

		return pokemon, nil
	}
	//
	log.Println("CACHE MISS:", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed fetching pokemon: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed fetching pokemon: %w", err)
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed fetching pokemon: %w", err)
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed fetching pokemon: %w", err)
	}

	//adding to cache new data from the url
	c.cache.Add(url, data)

	return pokemonResp, nil
}
