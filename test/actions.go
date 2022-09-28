package main

import "pokemon_cli"

func userAttackUser(attacker *pokemon_cli.PokemonUser, target *pokemon_cli.PokemonUser) {
	pattacker := GetActivePokemonAsPokemon(*attacker)
	ptarget := GetActivePokemonAsPokemon(*target)
	if pattacker.Hp > 0 {
		if pattacker.Dmg > ptarget.Hp {
			ptarget.Hp = 0
		} else {
			ptarget.Hp = ptarget.Hp - pattacker.Dmg
		}
	}
	letAiPlay()
}

func aiAttackUser(attacker *pokemon_cli.PokemonUser, target *pokemon_cli.PokemonUser) {
	pattacker := GetActivePokemonAsPokemon(*attacker)
	ptarget := GetActivePokemonAsPokemon(*target)
	if pattacker.Hp > 0 {
		if pattacker.Dmg > ptarget.Hp {
			ptarget.Hp = 0
		} else {
			ptarget.Hp = ptarget.Hp - pattacker.Dmg
		}
	}
	EndRound()
}

func userWantChangePokemon(user *pokemon_cli.PokemonUser) {
	SwitchToPokemonSelectionInCombat(user)
}

func userChangePokemon(user *pokemon_cli.PokemonUser, pokemonIndex int, aiPlay bool) {
	user.ActivePokemon = pokemonIndex
	if aiPlay {
		letAiPlay()
	} else {
		startRound()
	}
}

func userWantUsePotion(user *pokemon_cli.PokemonUser) {
	SwitchToPotionActionSelection(user)
}

func userSelectedPotion(user *pokemon_cli.PokemonUser, potionIndex int) {
	SwitchToPotionToPokemonActionSelection(user, potionIndex)
}

func userUsePotionOnPokemon(user *pokemon_cli.PokemonUser, potionIndex int, pokemonIndex int) {
	potion := user.ObjectInventory[potionIndex]
	pokemon := &user.PokemonInventory[pokemonIndex]
	if potion.Name == "Revive" {
		return
	}
	pokemon.Hp += potion.Heal
	if pokemon.Hp > pokemon.MaxHp {
		pokemon.Hp = pokemon.MaxHp
	}
	user.ObjectInventory = removeIndexPotions(potionIndex, user.ObjectInventory)
	letAiPlay()
}
