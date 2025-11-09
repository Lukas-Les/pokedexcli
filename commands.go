package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	start_location_url string = "https://pokeapi.co/api/v2/location-area/"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommand(name string) (cliCommand, error) {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Help",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "show map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "show previous map",
			callback:    commandMapb,
		},
	}
	cmd, ok := commands[name]
	if !ok {
		return cliCommand{}, errors.New("Unknown command")
	}
	return cmd, nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	messasge := `
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`
	fmt.Println(messasge)
	return nil
}

func commandMap(cfg *config) error {
	var url string
	if cfg.Next != nil {
		url = *cfg.Next
	} else {
		url = start_location_url
	}
	locationResponse, err := cfg.pokeapiClient.GetLocationAreas(url)
	if err != nil {
		return err
	}
	cfg.Next = locationResponse.Next
	cfg.Previous = locationResponse.Previous

	for _, area := range locationResponse.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	var url string
	if cfg.Previous != nil {
		url = *cfg.Previous
	} else {
		return errors.New("You're on the first page")
	}
	locationResponse, err := cfg.pokeapiClient.GetLocationAreas(url)
	if err != nil {
		return err
	}
	cfg.Next = locationResponse.Next
	cfg.Previous = locationResponse.Previous

	for _, area := range locationResponse.Results {
		fmt.Println(area.Name)
	}
	return nil
}
