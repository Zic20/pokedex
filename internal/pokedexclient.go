package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/zic20/pokedex/internal/pokecache"
)

type PokedexClient struct {
	cache      pokecache.Cache
	httpClient http.Client
	Next       string
	Previous   string
}

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

func NewPokedex(timeout, cacheInterval time.Duration) PokedexClient {
	return PokedexClient{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "",
	}
}

func (p *PokedexClient) FetchLocations(url string) (apiResponse, error) {
	var result apiResponse

	if val, ok := p.cache.Get(url); ok {
		err := json.Unmarshal(val, &result)
		if err != nil {
			return result, err
		}
		p.Next = result.Next
		p.Previous = result.Previous
		return result, nil
	}

	res, err := p.httpClient.Get(url)
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

	p.cache.Add(url, data)
	p.Next = result.Next
	p.Previous = result.Previous
	return result, nil
}
