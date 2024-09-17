package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/macrespo42/pokedexcli/internal/pokeapi"
	"github.com/macrespo42/pokedexcli/internal/pokecache"
)

type config struct {
	NextUrl     string
	PreviousUrl string
	Client      pokeapi.LocationArea
	cache       pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func getCommandList() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas, each call displays 20 next maps.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 location areas, each call displays 20 previous maps.",
			callback:    commandMapb,
		},
	}
}

func startREPL() {
	commandList := getCommandList()

	cfg := &config{
		NextUrl:     "https://pokeapi.co/api/v2/location-area/",
		PreviousUrl: "https://pokeapi.co/api/v2/location-area/",
		cache:       pokecache.NewCache(5 * time.Second),
	}

	for {
		fmt.Print("pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		rawCommand := scanner.Text()
		command, exist := commandList[rawCommand]
		if !exist {
			fmt.Println(rawCommand, "command not found")
			fmt.Println("type 'help' to get a list of available commands")
		} else {
			command.callback(cfg)
		}
	}
}
