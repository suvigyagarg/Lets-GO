package main

import (
	"math/rand"
	"errors"
	"fmt"
)

func commandCatch(cfg *config ,args ...string) error  {

	
	if len(args) != 1{
		return errors.New("no pokemon provided")
	}
	pokemonName := args[0]
 fmt.Printf("Throwing a Pokeball at %s... \n",pokemonName)

   resp ,err := cfg.pokeapiClient.GetPokemon(pokemonName)
   if err != nil {
      return err
   }
    const threshold = 50

   randNum := rand.Intn( resp.BaseExperience)

   if randNum >threshold{
    return  fmt.Errorf("%s escaped!",pokemonName)
   }
      cfg.caughtPokemon[pokemonName] =resp
	fmt.Printf("%s was caught! \n",pokemonName)
  
   
	return nil
}
