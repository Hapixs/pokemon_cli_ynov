package main

import (
	"io/ioutil"
	"log"
	"pokemon_cli"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

var pokemonsViewList = tview.NewList().ShowSecondaryText(true)
var pokemons = castMapToSliceOf()

var pokemonsSelectedViewList = tview.NewList()
var pokemonsSelected = []pokemon_cli.Pokemon{}

var pokemonsTrimedViewList = tview.NewList()
var pokemonsTrimed = []pokemon_cli.Pokemon{}

var pokeText = tview.NewTextView()
var helpText = tview.NewTextView().SetText("(q) pour quitter \n(a) pour passer sur la liste de selection\n(e) pour passer sur la liste séléctionée")
var pokemonAsciiArt = tview.NewTextView()

var isSearching = false
var trimSearch = ""

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

func PreparePokemonsViewList() {
	pokemonsViewList.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		pokemonsSelected = append(pokemonsSelected, pokemons[i])
		updatePokeList(pokemonsSelected, pokemonsSelectedViewList)
		pokemons = removeIndex(i, pokemons)
		updatePokeList(pokemons, pokemonsViewList)
		UpdateViewListIndex(i, *pokemonsViewList)
	})
	pokemonsViewList.SetChangedFunc(func(index int, name string, second_name string, shortcut rune) {
		updatePokeText(pokemons[index])
	})
}

func PreparePokemonsSelectedViewList() {
	pokemonsSelectedViewList.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		pokemons = append(pokemons, pokemonsSelected[i])
		updatePokeList(pokemons, pokemonsViewList)
		pokemonsSelected = removeIndex(i, pokemonsSelected)
		updatePokeList(pokemonsSelected, pokemonsSelectedViewList)
		if len(pokemonsSelected) <= 0 {
			switchToSelectionPokemons()
			return
		}
		UpdateViewListIndex(i, *pokemonsSelectedViewList)
	})
	pokemonsSelectedViewList.SetChangedFunc(func(index int, name string, second_name string, shortcut rune) {
		updatePokeText(pokemonsSelected[index])
	})

}
func PrepareTrimedViewList() {
	pokemonsTrimedViewList.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		pokemonsSelected = append(pokemonsSelected, pokemonsTrimed[i])
		updatePokeList(pokemonsSelected, pokemonsSelectedViewList)
		for index, p := range pokemons {
			if p.Name == pokemonsTrimed[i].Name {
				pokemons = removeIndex(index, pokemons)
				updatePokeList(pokemons, pokemonsViewList)
				break
			}
		}
		pokemonsTrimed = removeIndex(i, pokemonsTrimed)
		updatePokeList(pokemonsTrimed, pokemonsTrimedViewList)
		UpdateViewListIndex(i, *pokemonsTrimedViewList)
	})
	pokemonsTrimedViewList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		updatePokeText(pokemonsTrimed[index])
	})

}

func UpdateViewListIndex(index int, viewList tview.List) {
	if index >= len(pokemons)-1 {
		pokemonsViewList.SetCurrentItem(len(pokemons) - 1)
	} else if index > 0 {
		pokemonsViewList.SetCurrentItem(index)
	} else {
		pokemonsViewList.SetCurrentItem(0)
	}
}

var flex = tview.NewFlex()

func main() {
	PrepareCliApp()
	PreparePokemonsViewList()
	PreparePokemonsSelectedViewList()
	PrepareTrimedViewList()

	updatePokeList(pokemons, pokemonsViewList)

	switchToSelectionPokemons()

	if err := app.SetRoot(flex, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}

func updatePokeText(pok pokemon_cli.Pokemon) {
	pokeText.Clear()
	pokeText = pokeText.SetText("\nName: " + pok.Name +
		"\nHp: " + strconv.Itoa(pok.Hp) + "/" + strconv.Itoa(pok.MaxHp) +
		"\nDégats: " + strconv.Itoa(pok.Dmg) +
		"\nTypes: " + parseTabe(pok.Type))
	pokemonAsciiArt.Clear()
	pokemonAsciiArt.SetText(getPokemonImage(pok))
}

func updatePokeList(list []pokemon_cli.Pokemon, tList *tview.List) {
	tList.Clear()
	for _, pok := range list {
		tList.AddItem(pok.Name, "", '*', nil)
	}
}

func GetPokeCliLogo() string {
	content, err := ioutil.ReadFile("assets/pokecouille.txt")

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func getPokemonImage(pokemon pokemon_cli.Pokemon) string {
	content, err := ioutil.ReadFile("assets/" + pokemon.Name + ".txt")

	if err != nil {
		return ""
	}

	return string(content)
}
func switchToSelectedPokemons() {
	isSearching = false
	if len(pokemonsSelected) <= 0 {
		switchToSelectionPokemons()
		return
	}
	flex.Clear()
	app.SetRoot(flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewTextArea().SetText(GetPokeCliLogo()+"\n", false), 19, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(tview.NewTextView().SetText("Selection De pokemons"), 0, 0, false).
				AddItem(pokemonsViewList.SetSelectedBackgroundColor(tcell.ColorAntiqueWhite), 0, 1, false), 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(pokemonAsciiArt, 10, 1, false).
				AddItem(pokeText, 0, 1, false), 0, 1, false).
			AddItem(pokemonsSelectedViewList.SetSelectedBackgroundColor(tcell.ColorOrange), 0, 1, true).
			AddItem(helpText, 0, 1, false), 0, 1, true), true)
	updatePokeText(pokemonsSelected[0])
}

func switchToSelectionPokemons() {
	isSearching = false
	flex.Clear()
	app.SetRoot(flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewTextArea().SetText(GetPokeCliLogo()+"\n", false), 19, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(tview.NewTextView().SetText("Selection de Pokemons"), 0, 0, false).
				AddItem(pokemonsViewList.SetSelectedBackgroundColor(tcell.ColorOrange), 0, 1, true), 0, 1, true).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(pokemonAsciiArt, 0, 1, false).
				AddItem(pokeText, 0, 1, false), 0, 1, false).
			AddItem(pokemonsSelectedViewList.SetSelectedBackgroundColor(tcell.ColorAntiqueWhite), 0, 1, false).
			AddItem(helpText, 0, 1, true), 0, 6, true), true)
	updatePokeText(pokemons[0])
}

func switchToSearchPokemon(search bool) {
	isSearching = search
	flex.Clear()
	updatePokeList(pokemonsTrimed, pokemonsTrimedViewList)
	app.SetRoot(flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewTextArea().SetText(GetPokeCliLogo()+"\n", false), 19, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(tview.NewTextView().SetText("Selection De pokemons"), 0, 0, false).
				AddItem(pokemonsViewList.SetSelectedBackgroundColor(tcell.ColorAntiqueWhite), 0, 1, false), 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(pokemonAsciiArt, 0, 1, false).
				AddItem(pokeText, 0, 1, false), 0, 1, false).
			AddItem(pokemonsSelectedViewList.SetSelectedBackgroundColor(tcell.ColorAntiqueWhite), 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(helpText, 0, 1, false).
				AddItem(tview.NewInputField().SetText(trimSearch).SetChangedFunc(func(currentText string) {
					trimSearch = currentText
					pokemonsTrimed = trimByString(trimSearch, pokemons)
					switchToSearchPokemon(search)
				}), 0, 1, search).
				AddItem(pokemonsTrimedViewList.SetSelectedBackgroundColor(tcell.ColorOrange), 0, 1, !search), 0, 1, true), 0, 1, true), true)
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

func castMapToSliceOf() []pokemon_cli.Pokemon {
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
