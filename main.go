package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/zic20/pokedex/internal"
)

func main() {
	supportedCommands := map[string]internal.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    internal.CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    internal.CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas in the Pokemon world",
			Callback:    internal.CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Display the names of the previous 20 location areas in the Pokemon world",
			Callback:    internal.CommandMapB,
		},
	}

	sysConfig := internal.Config{
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
				err := cmd.Callback(&sysConfig)
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		}
	}
}
