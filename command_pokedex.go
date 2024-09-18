package main

import "fmt"

func commandPokedex(cfg *config, _ string) error {
	for _, pokemon := range cfg.pokedex {
		fmt.Println("-", pokemon.Name)
	}
	return nil
}
