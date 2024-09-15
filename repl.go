package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func startREPL() {
	commandList := getCommandList()

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
			command.callback()
		}
	}
}
