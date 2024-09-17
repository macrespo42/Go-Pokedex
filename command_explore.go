package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/macrespo42/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, area string) error {
	fullUrl := "https://pokeapi.co/api/v2/location-area/" + area

	var body []byte
	body, ok := cfg.cache.Get(fullUrl)

	if !ok {
		res, err := http.Get(fullUrl)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return fmt.Errorf("Response failed with StatusCode: %d", res.StatusCode)
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.cache.Add(fullUrl, body)
	}

	var areaDetail pokeapi.LocationAreaDetail
	err := json.Unmarshal(body, &areaDetail)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")

	for _, pokemon := range areaDetail.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
