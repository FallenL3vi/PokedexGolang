package pokeapi
import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)


func (c *Client) GetLocations(targetURL *string) (pokemonLocations, error) {
	newURL := baseURL

	if targetURL != nil {
		newURL = *targetURL
	}

	if val, ok := c.cache.Get(newURL); ok {
		locations := pokemonLocations{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return pokemonLocations{}, err
		}

		return locations, nil
	}

	req, err := http.NewRequest("GET", newURL, nil)

	if err != nil {
		fmt.Printf("error could not create request: %s \n", err)
		return pokemonLocations{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		fmt.Printf("error making http request %s\n", err)
		return pokemonLocations{}, err
	}
	
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("error could not read response body")
	}

	var dataLocations pokemonLocations

	err = json.Unmarshal(resBody, &dataLocations)

	if err != nil {
		fmt.Printf("error could not unmarshal json: %s\n", err)
		return pokemonLocations{}, err
	}

	//fmt.Printf("client: got response\n")
	//fmt.Printf("client: status code %d \n", res.StatusCode)

	c.cache.Add(newURL, resBody)
	return dataLocations, nil
}