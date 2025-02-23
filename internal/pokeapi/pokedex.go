package pokeapi

import (
	"sync"
)

type Pokedex struct {
	NameList   []string
	catchedMap map[string]PokemonDetails
	mu         *sync.RWMutex
}

// To create a new empty pokedex
func NewPokedex() Pokedex {
	p := Pokedex{
		catchedMap: make(map[string]PokemonDetails),
		mu:         &sync.RWMutex{},
	}
	return p
}

// To add a new pokemon in the pokedex
func (p *Pokedex) Add(pokemon PokemonDetails) bool {
	// Manage Mutex lock
	p.mu.Lock()
	defer p.mu.Unlock()

	// Check if Pokemon already in pokedex
	_, already := p.catchedMap[pokemon.Name]

	// Add the new Pokemon if not
	if !already {
		p.catchedMap[pokemon.Name] = pokemon
		p.NameList = append(p.NameList, pokemon.Name)
	}

	return already
}

// To get the details of a pokemon from the pokedex
func (p *Pokedex) Get(pokemon string) PokemonDetails {

}
