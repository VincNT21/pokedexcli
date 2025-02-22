package main

import "fmt"

func commandMapb(cfg *Config) error {
	// Check if it's the first page
	if cfg.isAtFirstPage {
		fmt.Println("You're already on the first page!")
		return nil
	}

	// Get the results
	result, err := cfg.pokeClient.GetLocationAreas(cfg.previousUrl)
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
