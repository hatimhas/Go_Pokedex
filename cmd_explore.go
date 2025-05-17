package main

import (
	"fmt"
)

func commandExplore(cfg *config, areaName ...string) error {
	if len(areaName) == 0 {
		return fmt.Errorf("please provide an area name")
	}

	pokemonAreaResp, err := cfg.pokeapiClient.FetchAreaDetails(areaName[0])
	if err != nil {
		return err
	}
	fmt.Println("Exploring: " + pokemonAreaResp.Name + "...")
	fmt.Println("Findable Pokemon at:", pokemonAreaResp.Name)
	for _, encounter := range pokemonAreaResp.PokemonEncounters {
		fmt.Println(" -", encounter.Pokemon.Name)
	}
	return nil
}
