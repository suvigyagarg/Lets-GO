package main

import (
	"fmt"
)

func commandPokedex(cfg *config ,args ...string) error  {
   fmt.Println("Pokemon in Pokedex : ")
   for _, pokemon := range cfg.caughtPokemon{
	fmt.Printf("Name: %v\n",pokemon.Name)
	// fmt.Printf("Height: %v\n",pokemon.Height)
	// fmt.Printf("Weight : %v\n",pokemon.Weight)
	// fmt.Println("Stats:")
	// for _,stat := range pokemon.Stats{
    // fmt.Printf("- %s: %v\n",stat.Stat.Name ,stat.BaseStat)
	// }
	// fmt.Println("Types:")
	// for _,typ := range pokemon.Types{
    // fmt.Printf("- %s\n",typ.Type.Name)
	// }
   }
	
	return nil
}
