package main

import "github.com/VincNT21/pokedexcli/internal/pokeapi"

func main() {
	pokeClient := pokeapi.NewClient()
	cfg := &config{
		pokeClient: pokeClient,
	}
	startRepl(cfg)
}
