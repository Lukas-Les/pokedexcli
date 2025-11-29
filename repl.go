package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	return strings.Fields(output)
}

func parseInput(text string) (string, []string) {
	clean := cleanInput(text)
	if len(clean) < 1 {
		return "", []string{}
	}
	cmdString := clean[0]
	if len(clean) < 2 {
		return cmdString, []string{}
	}
	args := clean[1:]
	return cmdString, args
}

func startRepl(cfg *config) {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	for {
		fmt.Print("Pokedex >")
		if scanner.Scan() {
			input := scanner.Text()
			cmdString, params := parseInput(input)
			cmd, err := getCommand(cmdString)
			if err != nil {
				fmt.Println(fmt.Errorf("Failed to run the command: %w\n", err))
				continue
			}
			cmd.callback(cfg, params)
		}
	}
}
