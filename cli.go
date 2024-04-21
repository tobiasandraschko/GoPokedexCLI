package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startCli(config *config) {
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
			command.callback(config)
			continue
		} else {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
			continue
		}
	}
}

