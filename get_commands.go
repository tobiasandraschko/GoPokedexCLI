package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    help,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exit,
		},
		"map": {
			name:        "map",
			description: "Displays the name of 20 location areas in the Pokemon world",
			callback:    mapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    mapB,
		},
	}
}