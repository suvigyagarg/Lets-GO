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
	callback    func(*config ,...string) error
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
		"explore":{
           name: "explore <location_area>",
		   description: "List all the Pokemons in a location area",
		   callback: commandExplore,
		},
		"catch":{
           name: "catch <pokemon_name>",
		   description: "Catch a pokemon",
		   callback: commandCatch,
		},
		"inspect":{
           name: "inspect <pokemon_name>",
		   description: "ispect about a caught pokemon",
		   callback: commandInspect,
		},
		"pokedex":{
           name: "pokedex",
		   description: "List all the pokemons caught",
		   callback: commandPokedex,
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
		args := []string{}
		if len(words)>1{
			args = words[1:]
		}
		command, exists := commands[commandName]

		if !exists {
			fmt.Println("Unknown command:", commandName)
			continue
		}

		err := command.callback(cfg ,args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}






