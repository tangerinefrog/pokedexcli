package main

import (
	"fmt"

	"github.com/tangerinefrog/pokedexcli/internal/pokeapi"
)

func exploreCallback(_ *pagingParam, param string) error {
	if param == "" {
		return fmt.Errorf("please, provide an area. To list all location areas, use 'map'")
	}

	locationArea, err := pokeapi.GetLocationArea(param)
	if err != nil {
		return err
	}

	fmt.Printf("All pokemon encounters in %s:\n\n", param)
	for _, v := range locationArea.PokemonEncounters {
		fmt.Printf("   - %s\n", v.Pokemon.Name)
	}

	return nil
}
