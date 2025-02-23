package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (RespShallowLocation, error) {
	// If no URL provided, use the default endpoint to location-area
	url := BaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check if cached data for this url first
	var data []byte
	entry, ok := c.cache.Get(url)
	if ok {
		// If yes, store them for use
		fmt.Println("*Cache data used!*")
		data = entry
	} else {
		// If not,
		// make GET request and handle result
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocation{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocation{}, err
		}
		defer resp.Body.Close()

		readData, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocation{}, nil
		}
		data = readData

		// Add the data to cache
		c.cache.Add(url, data)
	}

	// Decode data to result
	var result RespShallowLocation
	err := json.Unmarshal(data, &result)
	if err != nil {
		return RespShallowLocation{}, err
	}

	return result, nil
}
