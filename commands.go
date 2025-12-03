package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

const (
	start_location_url string = "https://pokeapi.co/api/v2/location-area/"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
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
		"explore": {
			name:        "explore",
			description: "explore a specific location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch a specific pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect pokemon you have catched",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "look at your pokedex",
			callback:    commandPokedex,
		},
	}
	cmd, ok := commands[name]
	if !ok {
		return cliCommand{}, errors.New("Unknown command")
	}
	return cmd, nil
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args []string) error {
	messasge := `
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`
	fmt.Println(messasge)
	return nil
}

func commandMap(cfg *config, args []string) error {
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

func commandMapb(cfg *config, args []string) error {
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

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("Single parameter required: explore <location>")
	}
	location := args[0]
	fmt.Printf("Exploring %s\n", location)
	locationUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)
	locationResponse, err := cfg.pokeapiClient.GetLocationArea(locationUrl)
	if err != nil {
		return err
	}
	var names = []string{}
	for _, item := range locationResponse.PokemonEncounters {
		names = append(names, item.Pokemon.Name)
	}
	if len(names) < 1 {
		fmt.Println("No pokemons found")
		return nil
	}
	fmt.Printf("Found Pokemon:\n")
	for _, name := range names {
		fmt.Printf("- %s\n", name)
	}
	return nil
}

func catch(baseExp int) bool {
	const K = 200.0
	chance := K / (K + float64(baseExp))
	roll := rand.Float64()
	return roll >= chance
}

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("Single parameter required: catch <pokemon>")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s... \n", name)
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	if !catch(pokemon.BaseExperience) {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	cfg.Pokedex[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	return nil
}

func commandInspect(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("Single parameter required: inspect <pokemon>")
	}
	name := args[0]
	pokemon, ok := cfg.Pokedex[name]
	if !ok {
		fmt.Printf("you have not caught that pokemon")
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, tp := range pokemon.Types {
		fmt.Printf("\t- %s\n", tp.Type.Name)
	}
	return nil
}

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.Pokedex) < 1 {
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf("\t- %s\n", pokemon.Name)
	}
	return nil
}
