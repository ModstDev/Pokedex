package main

import (
	"fmt"
	"os"
)

func commandExit(cfd *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
