package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationPokemon(locationAreaName string) (LocationAreaPokemons, error) {

	endpoint := "/location-area/" + locationAreaName
	fullUrl := baseURL + endpoint 
	
	//check cache
	cachedData, ok := c.cache.Get(fullUrl)

	if ok {
		// fmt.Println("cache hit!!")
		locationAreaPokemons := LocationAreaPokemons{}
		err := json.Unmarshal(cachedData, &locationAreaPokemons)
		if err != nil {
			return LocationAreaPokemons{}, err
		}
		return locationAreaPokemons, nil

	}
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaPokemons{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaPokemons{}, err
	}

	//closes that response object once we are done with it
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaPokemons{} , fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaPokemons{}, err
	}
	locationAreaPokemons := LocationAreaPokemons{}
	err = json.Unmarshal(data, &locationAreaPokemons)

	if err != nil {
		return LocationAreaPokemons{}, err
	}
 
    c.cache.Add(fullUrl ,data)

	return locationAreaPokemons, nil

}

