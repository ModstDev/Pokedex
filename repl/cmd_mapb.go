package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMapb(cfg *config) error {
	if cfg.previousUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := http.Get(cfg.previousUrl)
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
