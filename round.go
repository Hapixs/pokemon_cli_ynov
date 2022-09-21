package pokemon_cli

import (
	"fmt"
	"math/rand"
)

type UserAction struct {
	Name string
}

type PokemonSelection struct {
	selectedPokemon *Pokemon
}
type PotionSelection struct {
	Potion Potion
}

type PotionAction struct {
	Potion Potion
}

func StartNewRound(roundId int, user1 PokemonUser, user2 PokemonUser) {
	prepareRound(user1, roundId)
	prepareRound(user2, roundId)
	if roundId == 0 {
		user1 = initUser(user1)
		println("User 2: ")
		user2 = initUser(user2)
		fmt.Println(user1.PokemonInventory)
		fmt.Println(user2.PokemonInventory)
	}
	println("\n User 1 :")
	user1 = performForUser(user1, user2, roundId)
	println("\n User 2 :")
	user2 = performForUser(user2, user1, roundId)

	if user1.ActivePokemon.Hp <= 0 {
		if !hasRemainingPokemon(user1) {
			println("User2 Win ")
			return
		}
		askForPokemonSelection(user1, user1.PokemonInventory, true)
	} else if user2.ActivePokemon.Hp <= 0 {
		if !hasRemainingPokemon(user2) {
			println("User 1 Win")
			return
		}
		askForPokemonSelection(user2, user2.PokemonInventory, true)
	}

	StartNewRound(roundId+1, user1, user2)
}

func actionAttack(user, user_target PokemonUser, roundId int) {
	attackPokemonContext := pokemonAttackPokemon(user.ActivePokemon, user_target.ActivePokemon, roundId)
	println(attackPokemonContext.Attacker.Name + " attacked !")
}

func actionChangePokemon(user PokemonUser) PokemonUser {
	println("[!] Selection un pokemon de ton inventaire")
	pokemonSelection := askForPokemonSelection(user, user.PokemonInventory, true)
	user.ActivePokemon = pokemonSelection.selectedPokemon
	return user
}

func actionUsePotion(user PokemonUser) {
	potionSelection := askForPotionSelection(user.ObjectInventory)
	pokemonSelection := askForPokemonSelection(user, user.PokemonInventory, true)
	potionEffectContext := usePotions(potionSelection.Potion, *pokemonSelection.selectedPokemon)
	println(potionEffectContext.infoMessage)
}

func prepareRound(user PokemonUser, roundId int) {
	for _, p := range user.PokemonInventory {
		for i, e := range p.ActiveEffects {
			if e.Name == "brûlure" {
				p.Hp -= 3
			} else if e.Name == "poison" {
				p.Hp -= 5
			}
			if e.RoundStarted+e.RoundDuration == roundId {
				p.ActiveEffects = append(p.ActiveEffects[:i], p.ActiveEffects[i+1:]...)
			}
		}
	}

}

func performForUser(active, enemy PokemonUser, roundId int) PokemonUser {
	if active.Ai {
		// TODO generate AI
	}
	if active.ActivePokemon == nil {
		active = actionChangePokemon(active)
	}
	action := askForAction()
	if action.Name == "Attack" {
		actionAttack(active, enemy, roundId)
	} else if action.Name == "usePotion" {
		actionUsePotion(active)
	} else if action.Name == "changePokemon" {
		actionChangePokemon(active)
	}
	return active
}

func initUser(user PokemonUser) PokemonUser {
	if user.Ai {
		println("[!] Selection des pokemon de l'ia...")
		for i := 0; i < 7; i++ {
			user.PokemonInventory = append(user.PokemonInventory, GetPokemonList()[rand.Intn(len(GetPokemonList())-1)])
		}
	} else {
		println("[!] Selectionne un total de 7 pokemon pour jouer avec !")
		PrintList(GetPokemonList())
		for i := 0; i < 7; i++ {
			user.PokemonInventory = append(user.PokemonInventory, *askForPokemonSelection(user, GetPokemonList(), false).selectedPokemon)
		}
	}
	user.ActivePokemon = &user.PokemonInventory[0]
	return user
}

func askForAction() UserAction {
	println("[!] Choisis une action: ")
	print(" [1] Attaquer\n[2] Utiliser Object\n[3] Changer de pokemon")
	response := ""
	fmt.Scanln(&response)
	if response == "1" {
		return UserAction{"Attack"}
	} else if response == "2" {
		return UserAction{"usePotion"}
	} else if response == "3" {
		return UserAction{"changePokemon"}
	}
	return UserAction{""}
}

func askForPokemonSelection(user PokemonUser, selection []Pokemon, print bool) PokemonSelection {
	println("[!] Choisis un pokemon: ")
	if print {
		PrintList(selection)
	}

	response := ""
	fmt.Scanln(&response)

	total := 0
	for _, r := range response {
		if r >= 48 && r <= 57 {
			total = total*10 + int(r-48)
			continue
		}
	}
	if selection[total-1].Hp <= 0 {
		println("[!] Ce pokemon n'as plus d'hp !")
		return askForPokemonSelection(user, selection, false)
	}
	return PokemonSelection{&selection[total-1]} // Todo: check for out of range
}

func askForPotionSelection(potions []Potion) PotionSelection {
	println("[!] Selection une potion: ")
	for i, p := range potions {
		println("[" + string(rune(i+48+1)) + "] " + p.Name)
	}
	response := ""
	fmt.Scanln(&response)
	total := 0
	for _, r := range response {
		if r >= 48 && r <= 57 {
			total = total*10 + int(r-48)
			continue
		}
	}
	return PotionSelection{potions[total+1]} // Todo: check for out of range
}
