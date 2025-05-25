package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands = make(map[string]cliCommand)

func init() {
	commands["help"] = cliCommand{
		name:        "help",
		description: "Display available commands",
		callback:    helpCommand,
		config:      configStruct{},
	}
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the REPL",
		callback:    exitCommand,
		config:      configStruct{},
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "List the map locations",
		callback:    mapCommand,
		config:      configStruct{},
	}
}
func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Welcome to the Pokedex!\n")
	for {
		fmt.Print("Pokedex> ")
		scanner.Scan()
		input := CleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		command := input[0]

		if cmd, ok := commands[command]; ok {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", command)
		}
	}
}

func CleanInput(text string) []string {
	var words []string
	output := strings.ToLower(text)
	words = strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      configStruct
}

type configStruct struct {
	next string
	back string
}
