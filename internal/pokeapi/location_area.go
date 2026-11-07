package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/guillechuma/mypokedex/internal/pokecache"
)

type Area struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []Area
}

var cache = pokecache.NewCache(5 * time.Minute)

func GetLocationArea(url string) (LocationArea, error) {
	// Use cache
	v, ok := cache.Get(url)
	if ok {

		var locationArea LocationArea
		if err := json.Unmarshal(v, &locationArea); err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	var locationArea LocationArea
	if err := json.Unmarshal(body, &locationArea); err != nil {
		return LocationArea{}, err
	}

	cache.Add(url, body)
	return locationArea, nil
}
