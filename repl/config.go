package main

import "github.com/ModstDev/Pokedex/internal/pokeapi"

type config struct {
	pokeapiClient *pokeapi.Client
	pokedex       map[string]pokeapi.Pokemon
	nextUrl       *string
	previousUrl   *string
}
