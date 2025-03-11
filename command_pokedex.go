package main

import (
	"errors"
	"fmt"
)

func commandPokedex(conf *config, parameter *string) error {
	data, err := conf.pokeapiClient.GetCaughtPokemons()

	if err != nil {
		return err
	}
	
	if len(data) == 0 {
		return errors.New("Empty list of caught pokemons\n")
	}
	fmt.Printf("Your Pokedex:\n")
	for _, value := range data {
		fmt.Printf(" - %s\n", value)
	}
	return nil
}