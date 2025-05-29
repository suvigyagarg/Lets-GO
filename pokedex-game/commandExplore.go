package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config ,args ...string) error  {
	if len(args) != 1{
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]

   resp ,err := cfg.pokeapiClient.GetLocationPokemon(locationAreaName)
   if err != nil {
      return err
   }
   fmt.Printf("Pokemons in the %s:\n" ,resp.Name)
   for _,pokemon := range resp.PokemonEncounters{
	fmt.Printf("- %s \n" ,pokemon.Pokemon.Name)
   }
 
	return nil
}
