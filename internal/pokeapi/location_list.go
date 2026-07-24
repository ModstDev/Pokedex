package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	//Check cache before making request
	data, ok := c.cache.Get(url)
	if ok {
		log.Println("CACHE HIT:", url)
		var locations RespLocations

		err := json.Unmarshal(data, &locations)
		if err != nil {
			return RespLocations{}, err
		}

		return locations, nil
	}
	//
	log.Println("CACHE MISS:", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return RespLocations{}, err
	}

	locationsResp := RespLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespLocations{}, err
	}

	//adding to cache new data from the url
	c.cache.Add(url, data)

	return locationsResp, nil
}
