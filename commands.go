package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

type config struct {
	Next     string
	Previous string
}

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil // never reached. added to match the callback signature for commands
}

func commandHelp(conf *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}

func commandMap(conf *config) error {
	if conf.Next == "" {
		return errors.New("you're on the last page")
	}
	location_area, err := fetchPokemonMap(conf.Next, conf)
	if err != nil {
		return err
	}
	for _, location := range location_area.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapB(conf *config) error {
	if conf.Previous == "" {
		return errors.New("you're on the first page")
	}
	location_area, err := fetchPokemonMap(conf.Previous, conf)
	if err != nil {
		return err
	}
	for _, location := range location_area.Results {
		fmt.Println(location.Name)
	}
	return nil
}
