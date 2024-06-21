package main

import (
	"fmt"

	"github.com/tangerinefrog/pokedexcli/internal/inventory"
)

func inspectCallback(_ *pagingParam, param string) error {
	if param == "" {
		return fmt.Errorf("please provide Pokémon name")
	}

	p, ok := inventory.GetPokemon(param)
	if !ok {
		fmt.Printf("you have not caught %s", param)
	} else {
		printPokemon(p)
	}

	return nil
}

func printPokemon(pokemon inventory.Pokemon) {
	fmt.Printf("\nName: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for s, v := range pokemon.Stats {
		fmt.Printf("   - %s: %d\n", s, v)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("   - %s\n", t)
	}
	fmt.Printf("Abilities:\n")
	for _, a := range pokemon.Abilities {
		fmt.Printf("   - %s\n", a)
	}
}
