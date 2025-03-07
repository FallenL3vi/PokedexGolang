package main

import (
	"fmt"
)

func commandHelp(conf *config, parameter *string) error {
	commands := getCommands()
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, value := range commands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}