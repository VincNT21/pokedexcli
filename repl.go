package main

import (
	"fmt"
	"strings"

	"github.com/VincNT21/pokedexcli/internal/pokeapi"
	"github.com/chzyer/readline"
)

// Initialize the config registry
type config struct {
	pokeClient  pokeapi.Client
	nextUrl     *string
	previousUrl *string
	pokedex     pokeapi.Pokedex
}

// Command registry
type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

// Get a sort command list
var commandNames = []string{"help", "map", "mapb", "explore", "catch", "difficulty", "show", "exit"}

func startRepl(cfg *config) {

	// Welcome message + Help display
	fmt.Println()
	command := getCommands()["help"]
	command.callback(cfg)
	fmt.Println()

	// Create new readline instance
	rl, err := readline.New("Pokedex > ")
	if err != nil {
		fmt.Println("Error initializing readline:", err)
		return
	}
	defer rl.Close()

	// forever loop
	for {
		// Use readline to check for input
		line, err := rl.Readline()
		if err != nil { // ctrl-c, ctrl-d, etc.
			break
		}

		// Get only the first word (= command)
		words := cleanInput(line)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		// Get the rest of the input (= parameter args)
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		// Check if command exists in command registry and call the callback function
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command - Type 'help' for available commands")
			continue
		}
	}
}

// Convert input to lowercase slice of strings
func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	cleaned := strings.Fields(lowerText)
	return cleaned
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display this help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
		},
		"difficulty": {
			name:        "difficulty <pokemon_name>",
			description: "Show difficulty to catch a pokemon",
			callback:    commandDifficulty,
		},
		"show": {
			name:        "show",
			description: "Display Pokemon in the pokedex",
			callback:    commandShow,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
