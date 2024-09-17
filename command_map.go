package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/macrespo42/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	var body []byte
	body, ok := cfg.cache.Get(cfg.NextUrl)

	if ok {
		fmt.Println("picking inside cache")
	}

	if !ok {
		res, err := http.Get(cfg.NextUrl)
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

		cfg.cache.Add(cfg.NextUrl, body)
	}

	var areas pokeapi.LocationArea
	err := json.Unmarshal(body, &areas)
	if err != nil {
		return err
	}

	cfg.Client = areas
	cfg.NextUrl = areas.Next
	cfg.PreviousUrl = areas.Previous

	for _, area := range cfg.Client.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	var body []byte
	body, ok := cfg.cache.Get(cfg.PreviousUrl)

	if ok {
		fmt.Println("picking inside cache")
	}

	if !ok {
		res, err := http.Get(cfg.PreviousUrl)
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
		cfg.cache.Add(cfg.PreviousUrl, body)
	}

	var areas pokeapi.LocationArea
	err := json.Unmarshal(body, &areas)
	if err != nil {
		return err
	}

	cfg.Client = areas
	cfg.NextUrl = areas.Next
	cfg.PreviousUrl = areas.Previous

	for _, area := range cfg.Client.Results {
		fmt.Println(area.Name)
	}

	return nil
}
