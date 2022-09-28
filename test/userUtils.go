package main

import (
	"math/rand"
	"pokemon_cli"
	"time"
)

func GetActivePokemonAsPokemon(user pokemon_cli.PokemonUser) *pokemon_cli.Pokemon {
	return &user.PokemonInventory[user.ActivePokemon]
}

func hasRemainingPokemon(user *pokemon_cli.PokemonUser) bool {
	for _, p := range user.PokemonInventory {
		if p.Hp > 0 {
			return true
		}
	}
	return false
}

type dropRate struct {
	min int
	max int
}

func GeneratePotionInventory() []pokemon_cli.Potion {
	potion := []pokemon_cli.Potion{}
	for i := 0; i < 8; i++ {
		actualDrop := 0
		rand.Seed(time.Now().UnixNano())
		drop := rand.Intn(100)
		for _, p := range CastMapToSliceOfPotions() {
			actualDrop += p.DropRate
			if drop <= actualDrop {
				potion = append(potion, p)
				break
			}
		}
	}
	return potion
}
