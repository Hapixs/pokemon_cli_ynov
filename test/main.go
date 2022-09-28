package main

import (
	"os"
	"pokemon_cli"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func PrepareCliApp() {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' && !isSearching {
			app.Stop()
		} else if event.Rune() == 'a' && !isSearching {
			switchToSelectionPokemons()
		} else if event.Rune() == 'e' && !isSearching {
			switchToSelectedPokemons()
		} else if event.Rune() == 'p' && !isSearching {
			switchToSearchPokemon(true)
		} else if (event.Key() == tcell.KeyEsc || event.Key() == tcell.KeyDown) && isSearching {
			switchToSearchPokemon(false)
		}
		return event
	})
}

var flex = tview.NewFlex()

func main() {
	PrepareCliApp()
	PreparePokemonsViewList()
	PreparePokemonsSelectedViewList()
	PrepareTrimedViewList()
	UpdatePokeList(pokemons, pokemonsViewList)
	switchToSelectionPokemons()
	if err := app.SetRoot(flex, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}

var puser1 = pokemon_cli.PokemonUser{}
var puser2 = pokemon_cli.PokemonUser{}

var puser1Played = false

func startCombat(user1 pokemon_cli.PokemonUser, user2 pokemon_cli.PokemonUser) {
	puser1 = user1
	puser2 = user2
	startRound()
}

func startRound() {
	if !puser1.Ai {
		SwitchToCombatActionSelection()
	}
}

func letAiPlay() {
	if puser2.Ai {
		makeAiChoice()
	}
}

func EndRound() {
	if GetActivePokemonAsPokemon(puser2).Hp <= 0 {
		if !hasRemainingPokemon(&puser2) {
			os.Exit(1)
		}
		ChangeIAPokemon(&puser2)
	}
	if GetActivePokemonAsPokemon(puser1).Hp <= 0 {
		if !hasRemainingPokemon(&puser1) {
			os.Exit(1)
		}
		SwitchToPokemonSelectionAfterCombat(&puser1)
		return
	}
	startRound()
}

func trimByString(search string, list []pokemon_cli.Pokemon) []pokemon_cli.Pokemon {
	pokemonsTrimed = []pokemon_cli.Pokemon{}
	for _, p := range list {
		if StrIndex(p.Name, search) >= 0 {
			pokemonsTrimed = append(pokemonsTrimed, p)
		}
	}
	SortTableByString(pokemonsTrimed, search)
	return pokemonsTrimed
}

func CastMapToSliceOf() []pokemon_cli.Pokemon {
	pokemons := map[string]map[string]interface{}{
		"Bulbasaur": {
			"Type":  []string{"Grass", "Poison"},
			"HP":    22,
			"MaxHP": 22,
			"Dmg":   4,
		},
		"Charmander": {
			"Type":  []string{"Fire"},
			"HP":    24,
			"MaxHP": 24,
			"Dmg":   6,
		},
		"Squirtle": {
			"Type":  []string{"Water"},
			"HP":    23,
			"MaxHP": 23,
			"Dmg":   5,
		},
		"Pikachu": {
			"Type":  []string{"Electric"},
			"HP":    25,
			"MaxHP": 25,
			"Dmg":   5,
		},
		"Weedle": {
			"Type":  []string{"Bug", "Poison"},
			"HP":    16,
			"MaxHP": 16,
			"Dmg":   2,
		},
		"Ratata": {
			"Type":  []string{"Normal"},
			"HP":    20,
			"MaxHP": 20,
			"Dmg":   4,
		},
		"Sandshrew": {
			"Type":  []string{"Ground"},
			"HP":    19,
			"MaxHP": 19,
			"Dmg":   5,
		},
		"Nidoran": {
			"Type":  []string{"Poison"},
			"HP":    20,
			"MaxHP": 20,
			"Dmg":   4,
		},
		"Vulpix": {
			"Type":  []string{"Fire"},
			"HP":    23,
			"MaxHP": 23,
			"Dmg":   6,
		},
		"Psyduck": {
			"Type":  []string{"Water"},
			"HP":    23,
			"MaxHP": 23,
			"Dmg":   6,
		},
		"Mankey": {
			"Type":  []string{"Fighting"},
			"HP":    19,
			"MaxHP": 19,
			"Dmg":   7,
		},
		"Abra": {
			"Type":  []string{"Psychic"},
			"HP":    20,
			"MaxHP": 20,
			"Dmg":   6,
		},
		"Geodude": {
			"Type":  []string{"Rock", "Ground"},
			"HP":    23,
			"MaxHP": 23,
			"Dmg":   4,
		},
		"Magnemite": {
			"Type":  []string{"Electric", "Steel"},
			"HP":    24,
			"MaxHP": 24,
			"Dmg":   5,
		},
		"Mr. Mime": {
			"Type":  []string{"Psychic", "Fairy"},
			"HP":    24,
			"MaxHP": 24,
			"Dmg":   5,
		},
		"Eevee": {
			"Type":  []string{"Normal"},
			"HP":    20,
			"MaxHP": 20,
			"Dmg":   4,
		},
		"Vaporeon": {
			"Type":  []string{"Water"},
			"HP":    21,
			"MaxHP": 21,
			"Dmg":   4,
		},
		"Jolteon": {
			"Type":  []string{"Electric"},
			"HP":    23,
			"MaxHP": 23,
			"Dmg":   5,
		},
		"Flareon": {
			"Type":  []string{"Fire"},
			"HP":    22,
			"MaxHP": 22,
			"Dmg":   6,
		},
		"Sylveon": {
			"Type":  []string{"Fairy"},
			"HP":    21,
			"MaxHP": 21,
			"Dmg":   5,
		},
		"Espeon": {
			"Type":  []string{"Psychic"},
			"HP":    23,
			"MaxHP": 23,
			"Dmg":   4,
		},
		"Leafeon": {
			"Type":  []string{"Grass"},
			"HP":    22,
			"MaxHP": 22,
			"Dmg":   4,
		},
		"Glaceon": {
			"Type":  []string{"Ice"},
			"HP":    23,
			"MaxHP": 23,
			"Dmg":   5,
		},
		"Umbreon": {
			"Type":  []string{"Dark"},
			"HP":    25,
			"MaxHP": 25,
			"Dmg":   7,
		},
		"Mew": {
			"Type":  []string{"Psychic"},
			"HP":    27,
			"MaxHP": 27,
			"Dmg":   8,
		},
		"Jynx": {
			"Type":  []string{"Ice", "Psychic"},
			"HP":    20,
			"MaxHP": 20,
			"Dmg":   6,
		},
		"Glalie": {
			"Type":  []string{"Ice"},
			"HP":    19,
			"MaxHP": 19,
			"Dmg":   4,
		},
		"Ponyta": {
			"Type":  []string{"Fire"},
			"HP":    22,
			"MaxHP": 22,
			"Dmg":   5,
		},
		"Horsea": {
			"Type":  []string{"Water"},
			"HP":    21,
			"MaxHP": 21,
			"Dmg":   4,
		},
		"Vanillite": {
			"Type":  []string{"Ice"},
			"HP":    18,
			"MaxHP": 18,
			"Dmg":   4,
		},
		"Mareep": {
			"Type":  []string{"Electric"},
			"HP":    20,
			"MaxHP": 20,
			"Dmg":   5,
		},
		"Grimer": {
			"Type":  []string{"Poison"},
			"HP":    27,
			"MaxHP": 27,
			"Dmg":   6,
		},
		"Koffing": {
			"Type":  []string{"Poison"},
			"HP":    23,
			"MaxHP": 23,
			"Dmg":   4,
		},
		"Voltrob": {
			"Type":  []string{"Electric"},
			"HP":    21,
			"MaxHP": 21,
			"Dmg":   4,
		},
		"Magikarp": {
			"Type":  []string{"Water"},
			"HP":    14,
			"MaxHP": 14,
			"Dmg":   3,
		},
		"Cubchoo": {
			"Type":  []string{"Ice"},
			"HP":    18,
			"MaxHP": 18,
			"Dmg":   4,
		},
		"Torchic": {
			"Type":  []string{"Fire"},
			"HP":    19,
			"MaxHP": 19,
			"Dmg":   5,
		},
	}
	CleanPokemons := []pokemon_cli.Pokemon{}
	for k, p := range pokemons {
		CleanPokemons = append(CleanPokemons, pokemon_cli.Pokemon{k, p["Type"].([]string), p["HP"].(int), p["MaxHP"].(int), p["Dmg"].(int), []pokemon_cli.Effect{}, []pokemon_cli.Effect{}})
	}
	SortTableByName(CleanPokemons)
	return CleanPokemons
}

func CastMapToSliceOfPotions() []pokemon_cli.Potion {
	potions := map[string]map[string]interface{}{
		"Potion": {
			"Description": "Heal 6 HP",
			"Heal":        6,
			"DropRate":    60,
		},
		"Super Potion": {
			"Description": "Heal 10 HP",
			"Heal":        10,
			"DropRate":    20,
		},
		"Hyper Potion": {
			"Description": "Heal 20 HP",
			"Heal":        20,
			"DropRate":    10,
		},
		"Full Restore": {
			"Description": "Restore all HP and remove any status",
			"Heal":        50,
			"DropRate":    5,
		},
		"Revive": {
			"Description": "Revive a Pokemon with half his life",
			"Heal":        0,
			"DropRate":    5,
		},
	}
	CleanPotions := []pokemon_cli.Potion{}
	for k, p := range potions {
		CleanPotions = append(CleanPotions, pokemon_cli.Potion{k, p["Description"].(string), p["Heal"].(int), p["DropRate"].(int)})
	}
	return CleanPotions
}
