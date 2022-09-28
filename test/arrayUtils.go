package main

import (
	"math/rand"
	"pokemon_cli"
	"time"
)

func SortTableByString(table []pokemon_cli.Pokemon, reference string) {
	for i := 0; i < len(table); i++ {
		for j := i; j > 0 && StrIndex(table[j-1].Name, reference) > StrIndex(table[j].Name, reference); j-- {
			table[j], table[j-1] = table[j-1], table[j]
		}
	}
}

func SortTableByName(table []pokemon_cli.Pokemon) {
	for i := 0; i < len(table); i++ {
		for j := i; j > 0 && table[j-1].Name > table[j].Name; j-- {
			table[j], table[j-1] = table[j-1], table[j]
		}
	}
}

func removeIndex(index int, list []pokemon_cli.Pokemon) []pokemon_cli.Pokemon {
	if index == 0 {
		list = list[index+1:]
	} else if index == len(list)-1 {
		list = list[:index]
	} else {
		list = append(list[:index], list[index+1:]...)
	}
	return list
}

func removeIndexPotions(index int, list []pokemon_cli.Potion) []pokemon_cli.Potion {
	if index == 0 {
		list = list[index+1:]
	} else if index == len(list)-1 {
		list = list[:index]
	} else {
		list = append(list[:index], list[index+1:]...)
	}
	return list
}

func Split(str, sep string) []string {
	for len(str) > 0 {
		i := StrIndex(str, sep)
		if i < 0 {
			return []string{str}
		}
		return append([]string{str[:i]}, Split(str[i+len([]rune(sep)):], sep)...)
	}
	return []string{str}
}

func getRandomPokemonList() []pokemon_cli.Pokemon {
	rand.Seed(time.Now().UnixNano())
	pokemons := CastMapToSliceOf()
	list := []pokemon_cli.Pokemon{}
	for i := 0; i < 6; i++ {
		index := rand.Intn(len(pokemons) - 1)
		list = append(list, pokemons[index])
		pokemons = removeIndex(index, pokemons)
	}
	return list
}
