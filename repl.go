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

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words

}
func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Lists next page of  location areas ",
			callback:    callBackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous page of location areas",
			callback:    callBackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Lists the pokemon in a location area",
			callback:    callBackExplore,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all caught pokemons",
			callback:    callbackPokedex,
		},

		"catch": {
			name:        "catch {pokemon_name}",
			description: "Tries to catch a pokemon with given name",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect{pokemon_name}",
			description: "Inspects a pokemon if caught",
			callback:    callBackInspect,
		},

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}
	return commands
}
