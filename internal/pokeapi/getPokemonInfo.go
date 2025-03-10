package pokeapi
import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)


func (c *Client) GetPokemonInfo(pokemonName string) (Pokemon, error) {
	newURL := "https://pokeapi.co/api/v2/pokemon/"  + pokemonName

	if val, ok := c.cache.Get(newURL); ok {
		pokemonInfo := Pokemon{}
		err := json.Unmarshal(val, &pokemonInfo)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemonInfo, nil
	}

	req, err := http.NewRequest("GET", newURL, nil)

	if err != nil {
		fmt.Printf("error could not create request: %s \n", err)
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		fmt.Printf("error making http request %s\n", err)
		return Pokemon{}, err
	}
	
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("error could not read response body")
	}

	var dataPokemon Pokemon

	err = json.Unmarshal(resBody, &dataPokemon)

	if err != nil {
		fmt.Printf("error could not unmarshal json: %s\n", err)
		return Pokemon{}, err
	}

	//fmt.Printf("client: got response\n")
	//fmt.Printf("client: status code %d \n", res.StatusCode)

	c.cache.Add(newURL, resBody)
	return dataPokemon, nil
}