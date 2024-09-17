package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandInspect(cfg *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("you must provide a Pokemon name")
    }

    pokemonName := strings.ToLower(args[0])
    pokemon, ok := cfg.caughtPokemon[pokemonName]
    if !ok {
        return fmt.Errorf("you have not caught that pokemon")
    }

    fmt.Printf("Name: %s\n", pokemon.Name)
    fmt.Printf("Height: %d\n", pokemon.Height)
    fmt.Printf("Weight: %d\n", pokemon.Weight)
    fmt.Println("Stats:")
    for _, stat := range pokemon.Stats {
        fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
    }
    fmt.Println("Types:")
    for _, typeInfo := range pokemon.Types {
        fmt.Printf("  - %s\n", typeInfo.Type.Name)
    }

    return nil
}