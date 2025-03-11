package pokeapi

import (
	"errors"
)

func (c *Client) IsCaughtPokemon(pokemonName *string) (bool, error) {
	if *pokemonName == "" || pokemonName == nil {
		return false, errors.New("Can not check the pokemon, Invalid name")
	}

	if _, ok := c.catchedPokemons[*pokemonName]; ok{
		return true, nil;
	}
	return false, nil
}