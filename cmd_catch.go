package main

import (
	"fmt"

	"github.com/hatimhas/pokedexcli/internal/utils"
)

func commandCatch(cfg *config, pokemonName ...string) error {
	if len(pokemonName) == 0 {
		return fmt.Errorf("please provide a Pokemon name")
	}
	fmt.Println("Throwing a Pokeball at " + pokemonName[0] + "...")
	// Call function to catch the pokemon
	pokemonNameResp, err := cfg.pokeapiClient.FetchPokemonDetails(pokemonName[0])
	if err != nil {
		return err
	}
	pokemonBaseXP := pokemonNameResp.BaseExperience
	fmt.Println("Base XP: " + fmt.Sprint(pokemonBaseXP))
	if utils.CatchRateCalc(pokemonBaseXP) {
		fmt.Println("Caught " + pokemonNameResp.Name + "!")
		cfg.pokeDex[pokemonNameResp.Name] = pokemonNameResp
		return nil

	} else {
		fmt.Println("Failed to catch " + pokemonNameResp.Name + "!")
		return nil
	}
}
