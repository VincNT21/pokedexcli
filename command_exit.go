package main

import (
	"fmt"
	"os"
)

// commandExit callback
func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
