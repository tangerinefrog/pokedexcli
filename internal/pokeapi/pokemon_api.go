package pokeapi

import (
	"encoding/json"
	"fmt"
)

type PokemonResp struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`
	Height int `json:"height"`
	Name  string `json:"name"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func GetPokemon(name string) (PokemonResp, error) {
	resource := fmt.Sprintf("/pokemon/%s", name)
	data, err := fetchResource(resource)
	if err != nil {
		return PokemonResp{}, err
	}

	pokemon := PokemonResp{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return PokemonResp{}, err
	}

	return pokemon, nil
}
