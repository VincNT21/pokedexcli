package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func (c *Client) GetPokemonDetails(name string) (PokemonDetails, error) {
	// Url for the selected pokemon
	url := BaseURL + "/pokemon/" + name

	// Check if cached data for this url first
	if entry, ok := c.cache.Get(url); ok {
		// If yes, return the cached data decoded
		fmt.Println("*Cache data used!*")

		var result PokemonDetails
		err := json.Unmarshal(entry, &result)
		if err != nil {
			return PokemonDetails{}, err
		}
		return result, nil
	}

	// If not,
	// make GET request and handle result
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonDetails{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonDetails{}, nil
	}

	// Add the data to cache
	c.cache.Add(url, data)

	// Decode data to result
	var result PokemonDetails
	err = json.Unmarshal(data, &result)
	if err != nil {
		return PokemonDetails{}, err
	}

	return result, nil

}

func (c *Client) TryToCatchPokemon(pokemon PokemonDetails) bool {
	// Get Pokemon difficulty and set the limit to catch accordingly
	difficulty := c.GetPokemonDifficulty(pokemon)
	limitToCatch := pokemon.BaseExperience / difficulty

	// Try to catch it and return result (true/false)
	return rand.Intn(pokemon.BaseExperience) <= limitToCatch
}

func (c *Client) GetPokemonDifficulty(pokemon PokemonDetails) int {
	// Get Pokemon Base experience and set the difficulty of catch accordingly
	pokeBaseExp := pokemon.BaseExperience
	var difficulty int
	if pokeBaseExp < 50 {
		difficulty = 1
	} else if pokeBaseExp < 60 {
		difficulty = 2
	} else if pokeBaseExp < 70 {
		difficulty = 3
	} else if pokeBaseExp < 80 {
		difficulty = 4
	} else if pokeBaseExp < 90 {
		difficulty = 5
	} else if pokeBaseExp < 100 {
		difficulty = 6
	} else if pokeBaseExp < 150 {
		difficulty = 7
	} else if pokeBaseExp < 175 {
		difficulty = 8
	} else if pokeBaseExp < 200 {
		difficulty = 9
	} else {
		difficulty = 10
	}
	return difficulty
}
