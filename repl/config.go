package main

import "github.com/ModstDev/Pokedex/internal/pokeapi"

type config struct {
	pokeapiClient *pokeapi.Client
	nextUrl       *string
	previousUrl   *string
}
