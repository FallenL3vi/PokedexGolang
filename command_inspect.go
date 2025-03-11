package main

import (
	"errors"
)

func commandInspect(conf *config, parameter *string) error {
	if parameter == nil || *parameter == "" {
		return errors.New("Inaalid pokemon name")
	}
	
	err := conf.pokeapiClient.InspectPokemon(*parameter)

	if err != nil {
		return err
	}
	
	return nil
}