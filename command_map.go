package main

import (
	"fmt"
	"errors"
)

func commandMap(conf *config, parameter *string) error {
	dataLocations, err := conf.pokeapiClient.GetLocations(conf.nextURL)

	if err != nil {
		fmt.Printf("error could not get locations: %s\n", err)
		return err
	}

	conf.nextURL = dataLocations.Next
	conf.previousURL = dataLocations.Previous

	if len(dataLocations.Results) == 0 {
		fmt.Printf("error 0 locations where found")
		return errors.New("empty locations")
	}

	for _, slice := range dataLocations.Results {
		fmt.Printf("%s\n", slice.Name)
	}
	//fmt.Printf("Got locations")
	return nil
}


func commandMapPrev(conf *config, parameter *string) error {
	if conf.previousURL == nil {
		return errors.New("This is the first page")
	}

	dataLocations, err := conf.pokeapiClient.GetLocations(conf.previousURL)

	if err != nil {
		fmt.Printf("error could not get locations: %s\n", err)
		return err
	}

	conf.nextURL = dataLocations.Next
	conf.previousURL = dataLocations.Previous

	if len(dataLocations.Results) == 0 {
		fmt.Printf("error 0 locations where found")
		return errors.New("empty locations")
	}

	for _, slice := range dataLocations.Results {
		fmt.Printf("%s\n", slice.Name)
	}
	//fmt.Printf("Got locations")
	return nil
}