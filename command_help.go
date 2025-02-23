package main

import (
	"fmt"
)

// commandHelp callback
func commandHelp(cfg *config, parameter ...string) error {
	fmt.Println()
	fmt.Println("------")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range commandNames {
		fmt.Printf("%s: %s\n", command, getCommands()[command].description)
	}
	fmt.Println("------")
	fmt.Println()
	return nil
}
