package pokeapi

import (
	"errors"
	"fmt"
)

func (c *Client) InspectPokemon(pokemonName string) (error) {
	if pokemonName == "" {
		return errors.New("Can not inspect the pokemon, Invalid name")
	}

	if varr, ok := c.catchedPokemons[pokemonName]; ok {
		fmt.Printf("Name: %s\n", *varr.Name)
		fmt.Printf("Height: %v\n", varr.Height)
		fmt.Printf("Weight: %v\n", varr.Weight)
		fmt.Printf("Stats:\n")
		for _,data := range varr.Stats {
			fmt.Printf(" -%s: %v\n", *data.Stat.Name, data.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _,data := range varr.Types {
			fmt.Printf(" -%s\n", *data.Type.Name)
		}
		return nil
	}
	return errors.New("You have not caught that pokemon")
}