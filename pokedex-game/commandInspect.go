package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config ,args ...string) error  {

	
	if len(args) != 1{
		return errors.New("no pokemon provided")
	}
	pokemonName := args[0]

  pokemon ,ok :=cfg.caughtPokemon[pokemonName]
  if !ok {
	return errors.New("you have not caught the pokemon")
  }

    fmt.Printf("Name: %v\n",pokemon.Name)
	fmt.Printf("Height: %v\n",pokemon.Height)
	fmt.Printf("Weight : %v\n",pokemon.Weight)
	fmt.Println("Stats:")
	for _,stat := range pokemon.Stats{
    fmt.Printf("- %s: %v\n",stat.Stat.Name ,stat.BaseStat)
	}
	fmt.Println("Types:")
	for _,typ := range pokemon.Types{
    fmt.Printf("- %s\n",typ.Type.Name)
	}
	

	return nil
}
