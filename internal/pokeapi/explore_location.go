package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInLocation(location string) (LocationDetails, error) {
	// Url for the selected location
	url := BaseURL + "/location-area/" + location

	// Check if cached data for this url first
	if entry, ok := c.cache.Get(url); ok {
		// If yes, return the cached data decoded
		fmt.Println("*Cache data used!*")

		var result LocationDetails
		err := json.Unmarshal(entry, &result)
		if err != nil {
			return LocationDetails{}, err
		}
		return result, nil
	}

	// If not,
	// make GET request and handle result
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationDetails{}, nil
	}

	// Add the data to cache
	c.cache.Add(url, data)

	// Decode data to result
	var result LocationDetails
	err = json.Unmarshal(data, &result)
	if err != nil {
		return LocationDetails{}, err
	}

	return result, nil
}
