package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
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
}

func GetPokemon(url string) (Pokemon, error) {
	// Use cache
	v, ok := cache.Get(url)
	if ok {
		var pokemon Pokemon
		if err := json.Unmarshal(v, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return Pokemon{}, err
	}

	cache.Add(url, body)
	return pokemon, nil
}
