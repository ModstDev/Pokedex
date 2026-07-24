package main

import (
	"fmt"
	"os"
)

func commandExit(cfd *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
