package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/VincNT21/pokedexcli/internal/pokecache"
)

func (c *Client) GetLocationAreas(pageURL *string, pkCache *pokecache.Cache) (RespShallowLocation, error) {
	// If no URL provided, use the default endpoint to location-area
	url := BaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check if cached data for this url first
	entry, ok := pkCache.Get(url)
	var data []byte
	if ok {
		// If yes, store them for use
		data = entry
	} else {
		// If not,
		// make GET request
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
		pkCache.Add(url, data)
	}

	// Decode data to result
	var result RespShallowLocation
	err := json.Unmarshal(data, &result)
	if err != nil {
		return RespShallowLocation{}, err
	}

	return result, nil
}
