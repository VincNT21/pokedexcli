package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (RespShallowLocation, error) {
	// If no URL provided, use the default endpoint to location-area
	url := BaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Make GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocation{}, err
	}
	defer res.Body.Close()

	// Decode results
	var result RespShallowLocation
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return RespShallowLocation{}, err
	}

	return result, nil
}
