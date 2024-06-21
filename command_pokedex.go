package main

import (
	"fmt"

	"github.com/tangerinefrog/pokedexcli/internal/inventory"
)

func pokedexCallback(*pagingParam, string) error {
	caughtPokemon := inventory.ListPokemon()
	if len(caughtPokemon) == 0 {
		fmt.Println("You have no Pokémon in your collection!")
	} else {
		fmt.Println("Your Pokedex:")
		for _, p := range caughtPokemon {
			fmt.Printf("   - %s\n", p)
		}
	}

	return nil
}
