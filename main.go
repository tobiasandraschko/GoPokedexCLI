package main

import (
	"time"

	"github.com/tobiasandraschko/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)

	newConfig := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(newConfig)
}
