package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) ListPokemons(location string) (Location, error) {
	url := baseURL + "/location-area/"
	if location != "" {
		url += location + "/"
	}

	//Check cache before making request
	data, ok := c.cache.Get(url)
	if ok {
		log.Println("CACHE HIT:", url)
		var pokemons Location

		err := json.Unmarshal(data, &pokemons)
		if err != nil {
			return Location{}, err
		}

		return pokemons, nil
	}
	//
	log.Println("CACHE MISS:", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	pokemonsResp := Location{}
	err = json.Unmarshal(data, &pokemonsResp)
	if err != nil {
		return Location{}, err
	}

	//adding to cache new data from the url
	c.cache.Add(url, data)

	return pokemonsResp, nil
}
