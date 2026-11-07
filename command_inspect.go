package main

import (
	"fmt"
)

func commandInspect(cfg *config, pokemon ...string) error {
	if len(pokemon) == 0 {
		return nil
	}
	name := pokemon[0]
	val, ok := pokedex[name]
	if ok {
		fmt.Printf("Name: %s\n", val.Name)
		fmt.Printf("Height: %d\n", val.Height)
		fmt.Printf("Weight: %d\n", val.Weight)

		fmt.Println("Stats:")
		for _, stat := range val.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")
		for _, pokemonType := range val.Types {
			fmt.Printf("  - %s\n", pokemonType.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
