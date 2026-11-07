package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}
