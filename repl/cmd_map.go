package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(cfg *config) error {
	url := cfg.nextUrl

	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var data locationAreaResponse

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return err
	}

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	cfg.nextUrl = data.Next
	cfg.previousUrl = data.Previous

	return nil
}
