package main

import "fmt"

func commandDifficulty(cfg *config, args ...string) error {
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

	// Get the Pokemon difficulty
	pokemonDifficulty := cfg.pokeClient.GetPokemonDifficulty(pokemonDetails)

	// Print it
	fmt.Println()
	fmt.Printf("%v difficulty to catch is: %v\n", pokemonName, pokemonDifficulty)
	fmt.Println()

	return nil
}
