package main

import "fmt"

func commandHelp(cfg *config) error {
	commands := getCommandList()

	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")

	for _, value := range commands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	fmt.Println()

	return nil
}
