package pokeapi
import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)


func (c *Client) GetPokemons(areaName string) (pokemonInLocation, error) {
	newURL := baseURL + areaName

	if val, ok := c.cache.Get(newURL); ok {
		pokemons := pokemonInLocation{}
		err := json.Unmarshal(val, &pokemons)
		if err != nil {
			return pokemonInLocation{}, err
		}

		return pokemons, nil
	}

	req, err := http.NewRequest("GET", newURL, nil)

	if err != nil {
		fmt.Printf("error could not create request: %s \n", err)
		return pokemonInLocation{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		fmt.Printf("error making http request %s\n", err)
		return pokemonInLocation{}, err
	}
	
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("error could not read response body")
	}

	var dataPokemons pokemonInLocation

	err = json.Unmarshal(resBody, &dataPokemons)

	if err != nil {
		fmt.Printf("error could not unmarshal json: %s\n", err)
		return pokemonInLocation{}, err
	}

	//fmt.Printf("client: got response\n")
	//fmt.Printf("client: status code %d \n", res.StatusCode)

	c.cache.Add(newURL, resBody)
	return dataPokemons, nil
}