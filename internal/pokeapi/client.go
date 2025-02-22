package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	BaseURL    string
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		BaseURL: "https://pokeapi.co/api/v2",
	}
}

func (c *Client) GetLocationAreas(url string) (Location, error) {
	// If no URL provided, use the default endpoint to location-area
	if url == "" {
		url = c.BaseURL + "/location-area/"
	}

	// Make GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, fmt.Errorf("error creating new request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, fmt.Errorf("error with GET request: %v", err)
	}
	defer res.Body.Close()

	// Decode results
	var result Location
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return Location{}, fmt.Errorf("error decoding data: %v", err)
	}

	return result, nil
}
