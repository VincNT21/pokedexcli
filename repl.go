package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/VincNT21/pokedexcli/internal/pokeapi"
)

// Initialize the config registry
type config struct {
	pokeClient  pokeapi.Client
	nextUrl     *string
	previousUrl *string
}

// Initialize the command registry
type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(cfg *config) {

	// Welcome message + Help display
	fmt.Println()
	command := getCommands()["help"]
	command.callback(cfg)
	fmt.Println()

	// start a scanner to wait for user inputs
	scanner := bufio.NewScanner(os.Stdin)

	// forever loop
	for {
		fmt.Print("Pokedex > ")
		// when user press enter
		scanner.Scan()

		// Get only the first word (= command)
		words := cleanInput(scanner.Text())
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
