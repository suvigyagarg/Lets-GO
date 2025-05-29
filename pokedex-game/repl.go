package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var commands map[string]cliCommand


func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map":{
           name: "map",
		   description: "List name of 20 location area in Pokemon World",
		   callback: commandMap,
		},
		"mapb":{
           name: "mapb",
		   description: "List name of 20 location area in Pokemon World you previously viewed",
		   callback: commandMapBack,
		},
	}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := commands[commandName]

		if !exists {
			fmt.Println("Unknown command:", commandName)
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}






