package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {

	endpoint := "/pokemon/" +pokemonName
	fullUrl := baseURL + endpoint
	
	//check cache
	cachedData, ok := c.cache.Get(fullUrl)

	if ok {
		// fmt.Println("cache hit!!")
		pokemon := Pokemon{}
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil

	}
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	//closes that response object once we are done with it
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
 
    c.cache.Add(fullUrl ,data)

	return pokemon, nil

}

