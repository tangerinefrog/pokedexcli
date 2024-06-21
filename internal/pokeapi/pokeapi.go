package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

const baseUrl string = "https://pokeapi.co/api/v2"

func fetchResource(resource string) ([]byte, error) {
	url := baseUrl + resource

	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode > 299 {
		err = fmt.Errorf("resource %s returned unsuccessful status code: %d", url, resp.StatusCode)
		return []byte{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
