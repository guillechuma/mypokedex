package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	commands map[string]cliCommand
	Next     *string
	Previous *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	registry := getCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Errorf("Scanner error: %v", err)
		}
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		command := words[0]

		value, ok := registry[command]
		if ok {
			err := value.callback(cfg, words[1:]...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next page of location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous page of location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List of all the Pokemon in a given Area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Get the stats of a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all Pokemons in your Pokedex",
			callback:    commandPokedex,
		},
	}
}

func newConfig() *config {
	cfg := config{
		commands: getCommands(),
	}
	return &cfg
}

func cleanInput(text string) []string {
	res := strings.Fields(strings.ToLower(text))
	return res
}
