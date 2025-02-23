package main

import (
	"time"

	"github.com/VincNT21/pokedexcli/internal/pokeapi"
)

func main() {
	timeout := 5 * time.Second
	interval := 5 * time.Minute
	pokeClient := pokeapi.NewClient(timeout, interval)
	cfg := &config{
		pokeClient: pokeClient,
	}

	startRepl(cfg)
}
