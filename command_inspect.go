package main

import (
	"fmt"
)

func commandInspect(cfg *config, pokemonName string) error {
	pokemon, ok := cfg.pokedex[pokemonName]

	if !ok {
		fmt.Println("You have not caught that pokemon")
	} else {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")

		for _, stat := range pokemon.Stats {
			fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")

		for _, p := range pokemon.Types {
			fmt.Println("-", p.Type.Name)
		}

	}
	return nil
}
