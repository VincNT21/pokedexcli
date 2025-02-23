package main

import (
	"fmt"
	"slices"

	"github.com/VincNT21/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	// Handle no pokemon name given
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemonDetails := pokeapi.PokemonDetails{}

	// Check if pokemon is in the pokedex
	if slices.Contains(cfg.pokedex.NameList, pokemonName) {
		pokemonDetails = cfg.pokedex.Get(pokemonName)
	} else {
		return fmt.Errorf("you have not caught '%v'", pokemonName)
	}

	// Proper print of selected info
	fmt.Println()
	fmt.Printf("Name: %v\n", pokemonDetails.Name)
	fmt.Printf("Height: %v\n", pokemonDetails.Height)
	fmt.Printf("Weight: %v\n", pokemonDetails.Weight)
	fmt.Printf("Custom difficulty to catch: %v\n", cfg.pokeClient.GetPokemonDifficulty(pokemonDetails))
	fmt.Printf("Stats: \n")
	// NEED TO MANAGE STAT STRUCT IN A STRUCT...
	fmt.Printf("  -hp: %v\n", pokemonDetails.Stats[])
}
