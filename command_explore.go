package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	// Handle no location given
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	location := args[0]

	// Get the results
	locationDetails, err := cfg.pokeClient.GetPokemonInLocation(location)
	if err != nil {
		return fmt.Errorf("location '%v' seems not to exist", location)
	}

	// Print the results
	fmt.Println()
	fmt.Printf("Exploring %v...\n", location)
	fmt.Println()
	fmt.Println("Found Pokemon:")
	for _, enc := range locationDetails.PokemonEncounters {
		fmt.Printf("- %v\n", enc.Pokemon.Name)
	}
	fmt.Println()

	return nil
}
