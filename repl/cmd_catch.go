package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: catch <pokemon name>")
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)

	if err != nil {
		return fmt.Errorf("Failed to get data about pokemon!")
	}

	roll := rand.IntN(100)
	chance := 100 - pokemon.BaseExperience/4

	if roll < chance {
		fmt.Println(name, "was caught!")

		cfg.pokedex[name] = pokemon
	} else {
		fmt.Println(name, "escaped!")
	}

	return nil
}
