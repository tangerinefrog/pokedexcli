package main

import (
	"fmt"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pagingParam) error
}

type pagingParam struct {
	urlNext string
	urlPrev string
}

var commands map[string]cliCommand
var mapParams pagingParam

const PageSize int = 20

func init() {
	mapParams = pagingParam{}
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays this message",
			callback:    helpCallback,
		},
		"exit": {
			name:        "exit",
			description: "Exits PokedexCLI",
			callback:    exitCallback,
		},
		"map": {
			name:        "map",
			description: fmt.Sprintf("Prints next %d location areas", PageSize),
			callback:    mapCallback,
		},
		"mapb": {
			name:        "mapb",
			description: fmt.Sprintf("Prints previous %d location areas", PageSize),
			callback:    mapbCallback,
		},
	}

}

func executeCommand(name string) error {
	c, ok := commands[name]
	if !ok {
		return fmt.Errorf("unknown command")
	}

	err := c.callback(&mapParams)
	if err != nil {
		return err
	}

	fmt.Println()

	return nil
}
