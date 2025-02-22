package main

import (
	"fmt"
)

func commandMap(cfg *Config) error {
	// Check if it's the last page
	if cfg.isAtLastPage {
		fmt.Println("You've reached the last page! Use 'mapb' to go back")
		return nil
	}

	// Get the results
	result, err := cfg.pokeClient.GetLocationAreas(cfg.nextUrl)
	if err != nil {
		return err
	}

	// Update next url and is at last page
	if result.Next == nil {
		cfg.isAtLastPage = true
		cfg.nextUrl = ""
	} else {
		cfg.isAtLastPage = false
		cfg.nextUrl = *result.Next
	}

	// Update previous url and is at first page
	if result.Previous == nil {
		cfg.isAtFirstPage = true
		cfg.previousUrl = ""
	} else {
		cfg.isAtFirstPage = false
		cfg.previousUrl = *result.Previous
	}

	// Print the results
	fmt.Println()
	for _, loc := range result.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	return nil
}
