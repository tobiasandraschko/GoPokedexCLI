package main

import "github.com/tobiasandraschko/pokedexcli/internal/pokeapi"

type config struct {
	client  pokeapi.Client
	nextUrl *string
	prevUrl *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}