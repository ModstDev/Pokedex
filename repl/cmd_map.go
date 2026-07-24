package main

import (
	"fmt"
)

func commandMap(cfg *config, args []string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextUrl)
	if err != nil {
		return err
	}

	cfg.nextUrl = locationsResp.Next
	cfg.previousUrl = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
