package main

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *config) error {
	if cfg.previousUrl == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousUrl)
	if err != nil {
		return err
	}

	cfg.nextUrl = locationResp.Next
	cfg.previousUrl = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
