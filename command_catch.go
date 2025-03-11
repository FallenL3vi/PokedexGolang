package main

import (
	"fmt"
	"math/rand"
	"errors"
)

func commandCatch(conf *config, parameter *string) error {
	if parameter == nil || *parameter == "" {
		return errors.New("you must provide a pokemon name\n")
	}

	ok, err := conf.pokeapiClient.IsCaughtPokemon(parameter)

	if err != nil {
		return err
	}

	if ok {
		fmt.Printf("You have already caught that pokemon\n")
		return nil
	}

	dataPokemonInfo, err := conf.pokeapiClient.GetPokemonInfo(*parameter)

	if err != nil {
		fmt.Printf("error could not get pokemon info: %s\n", err)
		return err
	}

	if dataPokemonInfo.Name == nil {
		return errors.New("Could not get pokemon data\n")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", *dataPokemonInfo.Name)

	catchChance := rand.Intn(100)
	pokemonExp :=  dataPokemonInfo.BaseExperience
	if pokemonExp < 60 {
		catchChance += 60
	} else if pokemonExp < 100 {
		catchChance += 30
	} else if pokemonExp < 140 {
		catchChance += 20
	}

	if catchChance >= 70 {
		fmt.Printf("%s was caught!\n", *dataPokemonInfo.Name)

		/*if varr, ok := conf.pokeapiClient.catchedPokemons[*dataPokemonInfo.Name]; ok {
			return nil
		}*/

		err := conf.pokeapiClient.InsertPokemon(*dataPokemonInfo.Name, dataPokemonInfo)

		if err != nil {
			return err
		}

		return nil
	}

	fmt.Printf("%s escaped!\n", *dataPokemonInfo.Name)
	return nil
}