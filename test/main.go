package main

import (
	"pokemon_cli"
	"time"

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
	UpdatePokeList(pokemons, pokemonsViewList)
	switchToSelectionPokemons()
	if err := app.SetRoot(flex, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}

func switchToSelectedPokemons() {
	isSearching = false
	if len(pokemonsSelected) <= 0 {
		switchToSelectionPokemons()
		return
	}
	flex.Clear()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(BuildCliLogoItem(), 19, 1, false)
	bodyFlex := tview.NewFlex().SetDirection(tview.FlexColumn)
	bodyFlex.AddItem(BuildSelectionViewListItem(false).SetDirection(tview.FlexRow), 0, 1, false)
	bodyFlex.AddItem(BuildPokemonInfoFlexItem().SetDirection(tview.FlexRow), 0, 1, false)
	bodyFlex.AddItem(BuildSelectedViewListItem(true), 0, 1, true)
	bodyFlex.AddItem(BuildHelpText(), 0, 1, false)
	flex.AddItem(bodyFlex, 0, 1, true)
	app.SetRoot(flex, true)
	UpdatePokeText(pokemonsSelected[0])
}

func switchToSelectionPokemons() {
	isSearching = false
	flex.Clear()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(BuildCliLogoItem(), 19, 1, false)
	bodyFlex := tview.NewFlex().SetDirection(tview.FlexColumn)
	bodyFlex.AddItem(BuildSelectionViewListItem(true).SetDirection(tview.FlexRow), 0, 1, true)
	bodyFlex.AddItem(BuildPokemonInfoFlexItem().SetDirection(tview.FlexRow), 0, 1, false)
	bodyFlex.AddItem(BuildSelectedViewListItem(false), 0, 1, false)
	bodyFlex.AddItem(BuildHelpText(), 0, 1, false)
	flex.AddItem(bodyFlex, 0, 1, true)
	app.SetRoot(flex, true)
	UpdatePokeText(pokemons[0])
}

func switchToSearchPokemon(search bool) {
	isSearching = search
	flex.Clear()
	UpdatePokeList(pokemonsTrimed, pokemonsTrimedViewList)
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

func SwitchToCombatActionSelection() {

	flex.Clear()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(BuildCliLogoItem(), 19, 1, false)
	bodyFlex := tview.NewFlex().SetDirection(tview.FlexColumn)
	bodyFlex.AddItem(tview.NewTextView().SetText("Le combat va commencer..."), 0, 1, true)
	flex.AddItem(bodyFlex, 0, 1, true)
	app.SetRoot(flex, true)

	callTimer()

	actionList := tview.NewList().AddItem("Attaquer", "", '*', func() {}).AddItem("Changer de pokemon", "", '*', func() {}).AddItem("Utiliser un objets", "", '*', func() {})
	flex.Clear()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(BuildCliLogoItem(), 19, 1, false)
	bodyFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	bodyFlex.AddItem(actionList, 0, 1, true)
	flex.AddItem(bodyFlex, 0, 1, true)
	app.SetRoot(flex, true)

}

func callTimer() {
	timer := time.NewTimer(time.Second)
	<-timer.C
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
