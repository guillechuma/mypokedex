package main

import (
	"fmt"

	"github.com/guillechuma/mypokedex/internal/pokeapi"
)

func commandMap(cfg *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	areas, err := pokeapi.GetLocationArea(url)
	if err != nil {
		return err
	}
	cfg.Next = areas.Next
	cfg.Previous = areas.Previous

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	areas, err := pokeapi.GetLocationArea(*cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = areas.Next
	cfg.Previous = areas.Previous

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}
