package main

import (
	"fmt"
	"os"
)

func commandExit(conf *config, parameter *string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}