package main

import (
	"time"

	"github.com/VincNT21/pokedexcli/internal/pokeapi"
	"github.com/VincNT21/pokedexcli/internal/pokecache"
)

func main() {
	interval := 5 * time.Second
	pokeClient := pokeapi.NewClient()
	pokeCache := pokecache.NewCache(interval)
	cfg := &config{
		pokeClient: pokeClient,
		pokeCache:  pokeCache,
	}
	startRepl(cfg)
}
