package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *config) error {
	res, err := http.Get(cfg.NextUrl)
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

	var areas LocationArea
	err = json.Unmarshal(body, &areas)
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
	res, err := http.Get(cfg.PreviousUrl)
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

	var areas LocationArea
	err = json.Unmarshal(body, &areas)
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
