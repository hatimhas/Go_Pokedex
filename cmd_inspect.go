package main

import "fmt"

func commandInspect(cfg *config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("please provide a Pokemon name")
	}

	pokeName := params[0]
	pokedexVal, exist := cfg.pokeDex[pokeName]
	if exist {
		fmt.Println("Name:", pokedexVal.Name)
		fmt.Println("Height:", pokedexVal.Height)
		fmt.Println("Weight:", pokedexVal.Weight)
		fmt.Println("Stats:")
		fmt.Println(" -hp:", pokedexVal.Stats[0].BaseStat)
		fmt.Println(" -attack:", pokedexVal.Stats[1].BaseStat)
		fmt.Println(" -defense:", pokedexVal.Stats[2].BaseStat)
		fmt.Println(" -special-attack:", pokedexVal.Stats[3].BaseStat)
		fmt.Println(" -special-defense:", pokedexVal.Stats[4].BaseStat)
		fmt.Println(" -speed:", pokedexVal.Stats[5].BaseStat)
		fmt.Println("Types:")
		for _, t := range pokedexVal.Types {
			fmt.Println(" -", t.Type.Name)
		}
	} else {
		fmt.Println("You have not caught " + pokeName + " yet.")
	}
	return nil
}
