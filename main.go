package main

import (
	"time"

	"github.com/VincNT21/pokedexcli/internal/pokeapi"
)

func main() {
	timeout := 5 * time.Second
	interval := 5 * time.Minute
	pokeClient := pokeapi.NewClient(timeout, interval)
	pokedex := pokeapi.NewPokedex()
	cfg := &config{
		pokeClient: pokeClient,
		pokedex:    pokedex,
	}
	startRepl(cfg)
}
