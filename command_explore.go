package main

import (
	"fmt"

	"github.com/guillechuma/mypokedex/internal/pokeapi"
)

func commandExplore(cfg *config, location ...string) error {
	if len(location) == 0 {
		return nil
	}
	fmt.Printf("Exploring %s...\n", location[0])
	url := "https://pokeapi.co/api/v2/location-area/" + location[0]
	pokemonList, err := pokeapi.GetPokemonList(url)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonList.PokemonEncounters {
		fmt.Println(" -", pokemon.Pokemon.Name)
	}
	return nil
}
