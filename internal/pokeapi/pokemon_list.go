package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type PokemonList struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel int `json:"min_level"`
				MaxLevel int `json:"max_level"`
				Chance   int `json:"chance"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				ConditionValues []any `json:"condition_values"`
				PokemonDetails  any   `json:"pokemon_details"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetPokemonList(url string) (PokemonList, error) {
	// Use cache
	v, ok := cache.Get(url)
	if ok {
		var pokemonList PokemonList
		if err := json.Unmarshal(v, &pokemonList); err != nil {
			return PokemonList{}, err
		}
		return pokemonList, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return PokemonList{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonList{}, err
	}

	var pokemonList PokemonList
	if err := json.Unmarshal(body, &pokemonList); err != nil {
		return PokemonList{}, err
	}

	cache.Add(url, body)
	return pokemonList, nil
}
