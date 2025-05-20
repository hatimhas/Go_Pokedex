package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchPokemonDetails(pokemonName string) (PokemonDetailsResp, error) {
	pokemonURL := baseURL + "/pokemon/" + pokemonName
	if val, ok := c.cache.Get(pokemonURL); ok {
		fmt.Println("Cache Exist")
		detailedPokemonResp := PokemonDetailsResp{}

		err := json.Unmarshal(val, &detailedPokemonResp)
		if err != nil {
			return PokemonDetailsResp{}, err
		}

		return detailedPokemonResp, nil
	}

	req, err := http.NewRequest("GET", pokemonURL, nil)
	if err != nil {
		return PokemonDetailsResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonDetailsResp{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonDetailsResp{}, err
	}

	detailedPokemonResp := PokemonDetailsResp{}
	err = json.Unmarshal(dat, &detailedPokemonResp)
	if err != nil {
		return PokemonDetailsResp{}, err
	}

	c.cache.Add(pokemonURL, dat)
	return detailedPokemonResp, nil
}
