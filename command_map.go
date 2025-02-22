package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	// Check if it's the last page (and not just at the beginning)
	if cfg.nextUrl == nil && cfg.previousUrl != nil {
		return errors.New("you're already on the last page")
	}

	// Get the results
	result, err := cfg.pokeClient.GetLocationAreas(cfg.nextUrl)
	if err != nil {
		return err
	}

	// Update next/previous url
	cfg.nextUrl = result.Next
	cfg.previousUrl = result.Previous

	// Print the results
	fmt.Println()
	for _, loc := range result.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	return nil
}

func commandMapb(cfg *config) error {
	// Check if it's the first page
	if cfg.previousUrl == nil {
		return errors.New("you're already on the first page")
	}

	// Get the results
	result, err := cfg.pokeClient.GetLocationAreas(cfg.previousUrl)
	if err != nil {
		return err
	}

	// Update next/previous url
	cfg.nextUrl = result.Next
	cfg.previousUrl = result.Previous

	// Print the results
	fmt.Println()
	for _, loc := range result.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	return nil
}
