package main

import (
	"math/rand"
	"pokemon_cli"
	"time"
)

func ChangeIAPokemon(iauser *pokemon_cli.PokemonUser) {
	for {
		if !hasRemainingPokemon(iauser) {
			return
		}
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(len(iauser.PokemonInventory))
		if iauser.PokemonInventory[i].Hp <= 0 {
			continue
		}
		iauser.ActivePokemon = i
		return
	}
}

func makeAiChoice() {
	aiAttackUser(&puser2, &puser1)
}
