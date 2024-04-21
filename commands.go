package main

import (
	"errors"
	"fmt"
	"os"
)

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    listHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCli,
		},
		"map": {
			name:        "map",
			description: "Displays the name of 20 location areas in the Pokemon world",
			callback:    getNextMaps,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    getPreviousMaps,
		},
	}
}

func listHelp(config *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func exitCli(config *config) error {
	os.Exit(0)
	return nil
}


func getNextMaps(config *config) error {
	locationsResp, err := config.client.FetchLocations(config.nextUrl)
	if err != nil {
		return err
	}

	config.nextUrl = locationsResp.Next
	config.prevUrl = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func getPreviousMaps(config *config) error {
	if config.prevUrl == nil {
		fmt.Println("You're on the first page")
		return errors.New("you're on the first page")
	}

	locationResp, err := config.client.FetchLocations(config.prevUrl)
	if err != nil {
		return err
	}

	config.nextUrl = locationResp.Next
	config.prevUrl = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}