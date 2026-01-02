package internal

import (
	"errors"
	"fmt"
	"os"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(conf *Config) error
}

type Config struct {
	Next     string
	Previous string
}

func CommandExit(conf *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil // never reached. added to match the callback signature for commands
}

func CommandHelp(conf *Config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}

func CommandMap(conf *Config) error {
	if conf.Next == "" {
		return errors.New("you're on the last page")
	}
	location_area, err := FetchPokemonMap(conf.Next, conf)
	if err != nil {
		return err
	}
	for _, location := range location_area.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func CommandMapB(conf *Config) error {
	if conf.Previous == "" {
		return errors.New("you're on the first page")
	}
	location_area, err := FetchPokemonMap(conf.Previous, conf)
	if err != nil {
		return err
	}
	for _, location := range location_area.Results {
		fmt.Println(location.Name)
	}
	return nil
}
