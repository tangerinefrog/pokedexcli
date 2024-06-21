package pokeapi

import (
	"encoding/json"
	"fmt"
)

type LocationAreasResp struct {
	Count   int     `json:"count"`
	UrlNext *string `json:"next"`
	UrlPrev *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationAreaEntity struct {
	ID        int `json:"id"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func ListLocationAreas(offset int, limit int) (LocationAreasResp, error) {
	resource := fmt.Sprintf("/location-area?offset=%d&limit=%d", offset, limit)

	data, err := fetchResource(resource)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreas := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreas, nil
}

func GetLocationArea(name string) (LocationAreaEntity, error) {
	resource := fmt.Sprintf("/location-area/%s", name)
	data, err := fetchResource(resource)
	if err != nil {
		return LocationAreaEntity{}, err
	}

	locationArea := LocationAreaEntity{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationAreaEntity{}, err
	}

	return locationArea, nil
}
