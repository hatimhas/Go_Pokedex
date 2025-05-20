package main

import "fmt"

func commandPokedex(cfg *config, params ...string) error {
	if len(cfg.pokeDex) == 0 {
		return fmt.Errorf("your pokedex is empty. please catch some pokemon first")
	}
	fmt.Println("Pokedex:")
	for _, pokemon := range cfg.pokeDex {
		fmt.Println(" -", pokemon.Name)
	}
	return nil
}
