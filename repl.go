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

func startRepl(cfg *config) {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	for {
		fmt.Print("Pokedex >")
		if scanner.Scan() {
			input := scanner.Text()
			cmd, err := getCommand(input)
			if err != nil {
				fmt.Println(fmt.Errorf("Failed to run the command: %w\n", err))
				continue
			}
			cmd.callback(cfg)
		}
	}
}
