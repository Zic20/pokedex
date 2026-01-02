package internal

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(p *PokedexClient, url string) error
}

func CommandExit(p *PokedexClient, url string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil // never reached. added to match the callback signature for commands
}

func CommandHelp(p *PokedexClient, _ string) error {
	commands := GetCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, command := range commands {
		fmt.Printf("\n%s: %s\n", command.name, command.description)
	}
	return nil
}

func CommandMap(p *PokedexClient, _ string) error {
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
	p.Next = location_area.Next
	p.Previous = location_area.Previous
	return nil
}

func CommandMapB(p *PokedexClient, _ string) error {
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
	p.Next = location_area.Next
	p.Previous = location_area.Previous
	return nil
}

func CommandExplore(p *PokedexClient, url string) error {
	if url == "" {
		return errors.New("Explore expects exactly 1 argument")
	}
	mapDetails, err := p.ExploreLocation(url)
	if err != nil {
		return err
	}

	for _, encounter := range mapDetails.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}

func CommandCatch(p *PokedexClient, url string) error {
	if url == "" {
		return errors.New("Catch expects exactly 1 argument")
	}

	fullUrl := "https://pokeapi.co/api/v2/pokemon/" + url
	fmt.Printf("Throwing a Pokeball at %s...\n", url)
	pokemonInfo, err := p.FetchPokemonInfo(fullUrl)
	if err != nil {
		return err
	}

	chances := rand.Intn(pokemonInfo.BaseExperience + 20)
	fmt.Println(chances, pokemonInfo.BaseExperience)
	if chances >= pokemonInfo.BaseExperience {
		fmt.Printf("%s was caught\n", url)
		p.Pokedex[url] = pokemonInfo
	} else {
		fmt.Printf("%s escaped!\n", url)
	}
	return nil
}

func CommandInspect(p *PokedexClient, pokemonName string) error {
	if pokemonName == "" {
		return errors.New("Inspect expects exactly 1 argument")
	}

	pokemonInfo, ok := p.Pokedex[pokemonName]
	if !ok {
		fmt.Printf("you have not caught %s\n", pokemonName)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemonInfo.Name)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, c := range pokemonInfo.Types {
		fmt.Printf("- %s\n", c.Type.Name)
	}

	return nil
}

func CommandPokedex(p *PokedexClient, pokemonName string) error {
	if len(p.Pokedex) == 0 {
		fmt.Println("No pokemon has been added to your pokedex")
		return nil
	}

	for key := range p.Pokedex {
		fmt.Printf("- %s\n", key)
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
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			Callback:    CommandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Lists all Pokemons located at a location. It takes a city argument",
			Callback:    CommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to the the Pokemon specified by the user",
			Callback:    CommandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Gets the info of a pokemon that the user has caught",
			Callback:    CommandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists a pokemon has the user has caught",
			Callback:    CommandPokedex,
		},
	}
}
