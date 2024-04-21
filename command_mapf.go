package main

import (
	"fmt"
)

func mapF(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.FetchLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
