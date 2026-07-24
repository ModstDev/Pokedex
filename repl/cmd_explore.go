package main

import "fmt"

func commandExplore(cfg *config, args []string) error {

	if len(args) < 1 {
		return fmt.Errorf("usage: explore <location>")
	}

	location := args[0]
	fmt.Printf("Exploring %s...\n", location)

	//call API
	pokemonsResp, err := cfg.pokeapiClient.ExploreLocation(location)
	if err != nil {
		fmt.Println("Provided location is incorrect")
		return err
	}

	fmt.Println("Found:")

	for _, encounter := range pokemonsResp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
