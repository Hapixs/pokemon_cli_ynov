package pokemon_cli

import (
	"math/rand"
	"strconv"
	"strings"
)

func GetPokemonList() []Pokemon {
	pokemonsList := []Pokemon{}
	pokemonsList = append(pokemonsList, Pokemon{"Bulbasaur", []string{"Grass", "Poison"}, 22, 22, 4, []Effect{}, []Effect{GetPoisonEffect(0)}})
	pokemonsList = append(pokemonsList, Pokemon{"Charmander", []string{"Fire"}, 24, 24, 6, []Effect{}, []Effect{GetFireEffect(0)}})
	pokemonsList = append(pokemonsList, Pokemon{"Squirtle", []string{"Water"}, 23, 23, 5, []Effect{}, []Effect{}})
	pokemonsList = append(pokemonsList, Pokemon{"Pikachu", []string{"Electric"}, 25, 25, 5, []Effect{}, []Effect{GetElectrikEffect(0)}})
	pokemonsList = append(pokemonsList, Pokemon{"Weedle", []string{"Bug", "Poison"}, 16, 16, 2, []Effect{}, []Effect{GetPoisonEffect(0)}})
	pokemonsList = append(pokemonsList, Pokemon{"Ratata", []string{"Normal"}, 20, 20, 4, []Effect{}, []Effect{}})
	pokemonsList = append(pokemonsList, Pokemon{"Sandshrew", []string{"Ground"}, 19, 19, 5, []Effect{}, []Effect{}})
	pokemonsList = append(pokemonsList, Pokemon{"Nidoran", []string{"Poison"}, 20, 20, 4, []Effect{}, []Effect{GetPoisonEffect(0)}})
	pokemonsList = append(pokemonsList, Pokemon{"Vulpix", []string{"Fire"}, 23, 23, 6, []Effect{}, []Effect{GetFireEffect(0)}})
	pokemonsList = append(pokemonsList, Pokemon{"Psyduck", []string{"Water"}, 17, 23, 6, []Effect{}, []Effect{}})
	pokemonsList = append(pokemonsList, Pokemon{"Mankey", []string{"Fighting"}, 19, 19, 7, []Effect{}, []Effect{}})
	pokemonsList = append(pokemonsList, Pokemon{"Abra", []string{"Psychic"}, 20, 20, 6, []Effect{}, []Effect{GetPsyEffect(0)}})
	pokemonsList = append(pokemonsList, Pokemon{"Geodude", []string{"Rock", "Ground"}, 23, 23, 4, []Effect{}, []Effect{}})
	return pokemonsList
}

func pokemonAttackPokemon(attacker, reciever *Pokemon, actualRound int) AttackPokemonContext {
	context := AttackPokemonContext{attacker, reciever, attacker.Dmg, []Effect{}}
	reciever.Hp -= attacker.Dmg
	for _, e := range attacker.AttackEffects {
		if rand.Intn(100) <= e.PerformChance {
			newEffect := e
			newEffect.RoundStarted = actualRound
			reciever.ActiveEffects = append(reciever.ActiveEffects, e)
			context.BadEffect = append(context.BadEffect, e)
		}
	}
	if reciever.Hp <= 0 {
		reciever = nil
	}
	return context
}

func usePotions(pot Potion, reciever Pokemon) PotionEffectContext {
	context := PotionEffectContext{reciever, pot, false, ""}
	if reciever.Hp <= 0 {
		if pot.Name == "Revive" {
			reciever.Hp = reciever.MaxHp / 2
			context.worked = true
		}
		context.infoMessage = "Tu ne peux pas utiliser cet object sur un pokemone mort !"
		return context
	}
	reciever.Hp += pot.Heal
	if reciever.Hp > reciever.MaxHp {
		reciever.Hp = reciever.MaxHp
	}
	context.infoMessage = reciever.Name + " viens de regagner " + string(rune(pot.Heal+48)) + " Hp!"
	if pot.Name == "Full Restore" {
		reciever.ActiveEffects = []Effect{}
		context.infoMessage = reciever.Name + " a été soigné et tous les effets ce sont dissipés !"
	}
	context.worked = true
	return context
}

func hasRemainingPokemon(user PokemonUser) bool {
	for _, p := range user.PokemonInventory {
		if p.Hp > 0 {
			return true
		}
	}
	return false
}

func ParseString(s string) int {
	total := 0
	for _, c := range s {
		if c >= 48 && c <= 57 {
			total = total*10 + int(c-48)
			continue
		}
	}
	return total
}

func PrintList(list []Pokemon) {
	for i, v := range list {
		println("[" + strconv.Itoa(i+1) + "] " + v.Name + "")
		if v.Hp <= 0 {
			println("   [!] DEAD")
		} else {
			println("   HP: " + strconv.Itoa(v.Hp) + "/" + strconv.Itoa(v.MaxHp))
			println("   TYPE: " + strings.Join(v.Type, " "))
			println("   BAD EFFECT: ")
			for j, e := range v.ActiveEffects {
				print(e.Name)
				if j >= len(v.ActiveEffects)-1 {
					print(", ")
				}
			}
		}
		print("\n")
	}
}
