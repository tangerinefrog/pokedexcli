package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func init() {
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
	}
}

func getCommand(name string) (cliCommand, bool) {
	c, ok := commands[name]

	return c, ok
}
