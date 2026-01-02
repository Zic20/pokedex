package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	supportedCommands := map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapB,
		},
	}

	sysConfig := config{
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "",
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			cmd, ok := supportedCommands[userInput]
			if !ok {
				fmt.Println("Unknown command")
			} else {
				err := cmd.callback(&sysConfig)
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
			// cleanedInput := cleanInput(userInput)
		}
	}
}
