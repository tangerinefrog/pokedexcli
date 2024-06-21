package main

import (
	"fmt"

	"github.com/tangerinefrog/pokedexcli/internal/pokeapi"
)

func mapCallback() error {
	areasResp, err := pokeapi.ListLocationAreas(20, 20)
	if err != nil {
		return err
	}

	for _, v := range areasResp.Results {
		fmt.Printf("%s\n", v.Name)
	}

	return nil
}
