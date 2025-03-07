package main

import (
	"strings"
	"os"
	"fmt"
	"bufio"
	"github.com/FallenL3vi/PokedexGolang/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextURL *string
	previousURL *string
}

type cliCommand struct {
	name string
	description string
	callback func(conf *config, parameter *string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		}, "help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		}, "map": {
			name: "map",
			description: "Show next 20 maps",
			callback: commandMap,
		}, "mapb": {
			name: "mapb",
			description: "Show last 20 maps",
			callback: commandMapPrev,
		}, "explore": {
			name: "explore",
			description: "Explore <area> to list pokemons in this are",
			callback: commandExplore,
		},
	}
}

func cleanInput(text string) []string {
	lowerStr := strings.ToLower(text)
	separatedSlice := strings.Fields(lowerStr)
	return separatedSlice
}


func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			userInput := scanner.Text()

			words := cleanInput(userInput)

			if len(words) == 0 {
				continue
			}

			var argument string = ""
			if len(words) > 1 {
				argument = words[1]
			}

			command, exists := getCommands()[words[0]]

			if !exists {
				fmt.Printf("Unknown command\n")
				continue
			}

			err := command.callback(conf, &argument)

			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}