package internal

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(p *PokedexClient) error
}

func CommandExit(p *PokedexClient) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil // never reached. added to match the callback signature for commands
}

func CommandHelp(p *PokedexClient) error {
	commands := GetCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, command := range commands {
		fmt.Printf("\n%s: %s\n", command.name, command.description)
	}
	return nil
}

func CommandMap(p *PokedexClient) error {
	if p.Next == "" {
		return errors.New("you're on the last page")
	}
	location_area, err := p.FetchLocations(p.Next)
	if err != nil {
		return err
	}
	for _, location := range location_area.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func CommandMapB(p *PokedexClient) error {
	if p.Previous == "" {
		return errors.New("you're on the first page")
	}
	location_area, err := p.FetchLocations(p.Previous)
	if err != nil {
		return err
	}
	for _, location := range location_area.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			Callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the previous 20 location areas in the Pokemon world",
			Callback:    CommandMapB,
		},
	}
}
