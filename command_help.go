package main

import (
	"fmt"
	"sort"
)

func helpCallback(*pagingParam, string) error {
	fmt.Printf("Welcome to PokedexCLI!\n\n")

	maxLength := 0
	keys := make([]string, 0)
	for k, c := range commands {
		keys = append(keys, k)
		if len(c.name) > maxLength {
			maxLength = len(c.name)
		}
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%-*s  %s\n", maxLength, commands[k].name, commands[k].description)
	}

	return nil
}
