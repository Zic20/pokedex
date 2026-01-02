package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type apiResponse struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []locationsResult `json:"results"`
}

type locationsResult struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func FetchPokemonMap(url string, conf *Config) (apiResponse, error) {
	var result apiResponse
	res, err := http.Get(url)
	if err != nil {
		return result, fmt.Errorf("Error fetching location-area: %s", err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return result, fmt.Errorf("Error loading result: %s", err)
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("Error reading json data: %s", err)
	}

	conf.Next = result.Next
	conf.Previous = result.Previous
	return result, nil
}
