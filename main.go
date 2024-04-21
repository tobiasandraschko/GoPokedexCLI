package main

import (
	"time"

	"github.com/tobiasandraschko/pokedexcli/internal/pokeapi"
)

func main() {
	c := pokeapi.SpawnClient(5 * time.Second)
	config := &config{
		client: c,
	}
	startCli(config)
}
