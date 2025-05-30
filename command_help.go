package main

import (
	"fmt"
)

func helpCommand() error {
	fmt.Println("Available commands:")
	for _, command := range commands {
		fmt.Printf("%s - %s\n", command.name, command.description)
	}
	return nil
}
