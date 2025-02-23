package main

import (
	"fmt"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	// Handle no pokemon name given
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	pokemonName := args[0]

	// Get the pokemon info
	pokemonDetails, err := cfg.pokeClient.GetPokemonDetails(pokemonName)
	if err != nil {
		return fmt.Errorf("pokemon '%v' seems not to exist", pokemonName)
	}

	// Display text
	fmt.Println()
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)

	// Trying to catch pokemon
	time.Sleep(1 * time.Second)
	result := cfg.pokeClient.TryToCatchPokemon(pokemonDetails)

	// Handle failed/successful catch
	if !result {
		fmt.Printf("Failed... %v escaped!\n", pokemonName)
		fmt.Println()
		return nil
	} else {
		already := cfg.pokedex.Add(pokemonDetails)
		if !already {
			fmt.Printf("Success! %v was caught and added to Pokedex!\n", pokemonName)
			fmt.Println()
		} else {
			fmt.Printf("Success ! But '%v' already in pokedex\n", pokemonDetails.Name)
			fmt.Println()
		}

		return nil
	}
}
