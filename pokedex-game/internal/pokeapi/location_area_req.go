package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea(pageURL *string) (LocationAreasResp, error) {

	endpoint := "/location-area"
	fullUrl := baseURL + endpoint
	if pageURL != nil {
		fullUrl = *pageURL
	}
	//check cache
	cachedData, ok := c.cache.Get(fullUrl)

	if ok {
		// fmt.Println("cache hit!!")
		locationAreaResp := LocationAreasResp{}
		err := json.Unmarshal(cachedData, &locationAreaResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreaResp, nil

	}
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}

	//closes that response object once we are done with it
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}
	locationAreaResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreasResp{}, err
	}
 
    c.cache.Add(fullUrl ,data)

	return locationAreaResp, nil

}




