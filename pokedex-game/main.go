package main

import (
	"time"

	pokeapi "github.com/suvigyagarg/pokedex-game/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),  
	}
	startRepl(&cfg)

}
