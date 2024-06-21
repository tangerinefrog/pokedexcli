package main

import (
	"fmt"
	"time"

	"github.com/tangerinefrog/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pagingParam, string) error
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
		"explore": {
			name:        "explore <location_area_name>",
			description: "Prints a list of all Pokémon in a selected area",
			callback:    exploreCallback,
		},
	}
	pokecache.NewCache(10 * time.Second)
}

func executeCommand(name string, param string) error {

	c, ok := commands[name]
	if !ok {
		return fmt.Errorf("unknown command")
	}

	err := c.callback(&mapParams, param)
	if err != nil {
		return err
	}

	fmt.Println()

	return nil
}
