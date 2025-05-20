package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, params ...string) error {
	fmt.Println(params)
	for _, pokemon := range cfg.pokeDex {
		fmt.Println("Pokemon in Pokedex:", pokemon.pokeInfo.Name)
	}
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
