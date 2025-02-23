package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	// Get the pokedex list
	catchedList := cfg.pokedex.NameList

	// Handle empty pokedex
	if len(catchedList) == 0 {
		return fmt.Errorf("pokedex is empty")
	}

	// Print results
	fmt.Println()
	fmt.Println("Pokemon in pokedex:")
	for _, poke := range catchedList {
		fmt.Printf("- %v", poke)
	}

	fmt.Println()

	return nil
}
