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
