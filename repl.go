package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tobiasandraschko/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	commandsMap := getCommands()

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := reader.Text()
		wordArray := strings.Fields(words)
		if len(wordArray) == 0 {
			continue
		}

		commandName := wordArray[0]
		command, exists := commandsMap[commandName]

		if exists {
			command.callback(cfg)
			continue
		} else {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
			continue
		}
	}
}

