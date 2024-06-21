package main

import (
	"fmt"
	"math/rand"

	"github.com/tangerinefrog/pokedexcli/internal/inventory"
	"github.com/tangerinefrog/pokedexcli/internal/pokeapi"
)

const catchChance float64 = 0.5

func catchCallback(_ *pagingParam, param string) error {
	if param == "" {
		return fmt.Errorf("please provide Pokémon name")
	}

	pokemonResp, err := pokeapi.GetPokemon(param)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a ball at %s...\n", pokemonResp.Name)

	if rand.Float64() < catchChance {
		var abilities []string
		for _, v := range pokemonResp.Abilities {
			abilities = append(abilities, v.Ability.Name)
		}
		var types []string
		for _, v := range pokemonResp.Types {
			types = append(types, v.Type.Name)
		}
		stats := make(map[string]int)
		for _, v := range pokemonResp.Stats {
			stats[v.Stat.Name] = v.BaseStat
		}

		inventory.AddPokemon(
			inventory.Pokemon{
				Name:      pokemonResp.Name,
				Height:    pokemonResp.Height,
				Weight:    pokemonResp.Weight,
				Abilities: abilities,
				Types:     types,
				Stats:     stats,
			})

		fmt.Printf("%s was caught!\nYou may now inspect it with the inspect command.\n", pokemonResp.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
	}

	return nil
}
