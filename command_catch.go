package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/guillechuma/mypokedex/internal/pokeapi"
)

var pokedex = make(map[string]pokeapi.Pokemon)

func commandCatch(cfg *config, pokemon ...string) error {
	if len(pokemon) == 0 {
		return nil
	}
	name := pokemon[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	caughtPokemon, err := pokeapi.GetPokemon(url)
	if err != nil {
		return err
	}
	experience := float64(caughtPokemon.BaseExperience)
	chance := 100.0 / (100.0 + experience)

	if rand.Float64() < chance {
		fmt.Printf("%s was caught!\n", name)
		pokedex[name] = caughtPokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
