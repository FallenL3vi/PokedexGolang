package pokeapi

import (
	"errors"
)

func (c *Client) InsertPokemon(pokemonName string, pokemonData Pokemon) (error) {
	if pokemonName == "" {
		return errors.New("Can not insert empty pokemon name")
	}
	c.catchedPokemons[pokemonName] = pokemonData
	return nil
}