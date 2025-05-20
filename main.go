package main

import (
	"time"

	"github.com/hatimhas/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokeDex:       make(map[string]pokeapi.PokemonDetailsResp),
	}

	startRepl(cfg)
}
