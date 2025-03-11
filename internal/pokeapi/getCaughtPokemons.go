package pokeapi

import (
	"errors"
)

func (c *Client) GetCaughtPokemons() ([]string, error) {

	if len(c.catchedPokemons) == 0 {
		return nil, errors.New("You have not caught any pokemon\n")
	}

	var pokemons []string = make([]string, len(c.catchedPokemons))

	var index int = 0
	for _, value := range c.catchedPokemons {
		pokemons[index] = *value.Name
		index += 1
	}


	return pokemons, nil
}