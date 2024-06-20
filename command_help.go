package main

import "fmt"

func helpCallback() error {
	fmt.Printf("Welcome to PokedexCLI!\n")
	for _, v := range commands {
		fmt.Printf("\n\t%s\t%s", v.name, v.description)
	}
	fmt.Printf("\n\n")

	return nil
}
