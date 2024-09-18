package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/macrespo42/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, pokemonName string) error {
	fullUrl := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	res, err := http.Get(fullUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with StatusCode: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var pokemon pokeapi.Pokemon

	err = json.Unmarshal(body, &pokemon)
	fmt.Printf("Throwing a pokeball at %s...\n", pokemonName)

	chanceToCatch := rand.Intn(int(pokemon.BaseExperience/50) + 1)
	isCatched := chanceToCatch == 1
	if isCatched {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.pokedex[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}
