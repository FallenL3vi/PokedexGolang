package main

import (
	"fmt"
	"errors"
)

func commandExplore(conf *config, parameter *string) error {
	dataPokemons, err := conf.pokeapiClient.GetPokemons(*parameter)

	if err != nil {
		fmt.Printf("error could not get pokemons in location: %s\n", err)
		return err
	}

	if len(dataPokemons.PokemonEncounters) == 0 {
		fmt.Printf("error 0 pokemons were found")
		return errors.New("empty pokemons")
	}

	fmt.Printf("Exploring %s...\n", *parameter)
	fmt.Printf("Found Pokemon:\n")

	for _, val := range dataPokemons.PokemonEncounters {
		fmt.Printf(" - %s\n", *val.Pokemon.Name)
	}

	return nil
}