package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
    if len(cfg.caughtPokemon) == 0 {
        fmt.Println("Your Pokedex is empty.")
        return nil
    }

    fmt.Println("Your Pokedex:")
    for name := range cfg.caughtPokemon {
        fmt.Printf(" - %s\n", name)
    }

    return nil
}