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

type locationResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

const (
	BaseUrl = "https://pokeapi.co/api/v2/location-area"
)

type nameUrl struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type exploreRespone struct {
	EncounterMethodRates []struct {
		EncounterMethod nameUrl `json:"encounter_method"`
		VersionDetails  []struct {
			Rate    int     `json:"rate"`
			Version nameUrl `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`

	GameIndex int     `json:"game_index"`
	Id        int     `json:"id"`
	Location  nameUrl `json:"location"`
	Name      string  `json:"name"`
	Names     []struct {
		Language nameUrl `json:"language"`
		Name     string  `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon        nameUrl `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int     `json:"chance"`
				ConditionValues []any   `json:"condition_values"`
				MaxLevel        int     `json:"max_level"`
				Method          nameUrl `json:"method"`
				MinLevel        int     `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int     `json:"max_chance"`
			Version   nameUrl `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func NewPokedex(timeout, cacheInterval time.Duration) PokedexClient {
	return PokedexClient{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
		Next:     BaseUrl,
		Previous: "",
	}
}

func (p *PokedexClient) FetchLocations(url string) (locationResponse, error) {
	var result locationResponse
	if val, ok := p.cache.Get(url); ok {
		err := json.Unmarshal(val, &result)
		if err != nil {
			return result, err
		}
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
	return result, nil
}

func (p *PokedexClient) ExploreLocation(url string) (exploreRespone, error) {
	var result exploreRespone
	fullUrl := BaseUrl + "/" + url
	if val, ok := p.cache.Get(fullUrl); ok {
		err := json.Unmarshal(val, &result)
		if err != nil {
			return result, err
		}
		return result, nil
	}

	res, err := p.httpClient.Get(fullUrl)
	if err != nil {
		return result, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return result, fmt.Errorf("Error reading result: %s", err)
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("Error converting data: %s", err)
	}

	p.cache.Add(fullUrl, data)
	return result, nil
}
