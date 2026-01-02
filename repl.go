package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zic20/pokedex/internal"
)

func CleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	str := strings.Trim(strings.ToLower(text), " ")
	return strings.Split(str, " ")
}

func runrepl(p *internal.PokedexClient) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := internal.GetCommands()
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			inputs := CleanInput(userInput)
			if len(inputs) == 0 {
				fmt.Println("No input provided")
				continue
			}
			cmd, ok := commands[inputs[0]]

			if !ok {
				fmt.Println("Unknown command")
			} else {
				var err error
				if len(inputs) > 1 {
					err = cmd.Callback(p, inputs[1])
				} else {
					err = cmd.Callback(p, "")
				}
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		}
	}
}
