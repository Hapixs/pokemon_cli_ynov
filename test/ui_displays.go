package main

import (
	"pokemon_cli"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

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
	bodyFlex.AddItem(BuildActionList(), 0, 1, true).
		AddItem(BuildActivePokemon(puser1), 0, 1, false).
		AddItem(BuildActivePokemon(puser2), 0, 1, false)
	flex.AddItem(bodyFlex, 0, 1, true)
	app.SetRoot(flex, true)
}

func SwitchToPokemonSelectionInCombat(selector *pokemon_cli.PokemonUser) {
	flex.Clear()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(BuildCliLogoItem(), 19, 1, false)
	bodyFlex := tview.NewFlex().SetDirection(tview.FlexColumn)
	bodyFlex.AddItem(BuildPokemonSelectionInCombatListItem(selector), 0, 1, true).
		AddItem(BuildActivePokemon(puser1), 0, 1, false).
		AddItem(BuildActivePokemon(puser2), 0, 1, false)
	flex.AddItem(bodyFlex, 0, 1, true)
	app.SetRoot(flex, true)
}

func SwitchToPokemonSelectionAfterCombat(selector *pokemon_cli.PokemonUser) {
	flex.Clear()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(BuildCliLogoItem(), 19, 1, false)
	bodyFlex := tview.NewFlex().SetDirection(tview.FlexColumn)
	bodyFlex.AddItem(BuildPokemonSelectionAfterCombatListItem(selector), 0, 1, true).
		AddItem(BuildActivePokemon(puser1), 0, 1, false).
		AddItem(BuildActivePokemon(puser2), 0, 1, false)
	flex.AddItem(bodyFlex, 0, 1, true)
	app.SetRoot(flex, true)
}

func SwitchToPotionActionSelection(selector *pokemon_cli.PokemonUser) {
	flex.Clear()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(BuildCliLogoItem(), 19, 1, false)
	bodyFlex := tview.NewFlex().SetDirection(tview.FlexColumn)
	bodyFlex.AddItem(BuildPotionSelectionInCombatListItem(selector), 0, 1, true).
		AddItem(BuildActivePokemon(puser1), 0, 1, false).
		AddItem(BuildActivePokemon(puser2), 0, 1, false)
	flex.AddItem(bodyFlex, 0, 1, true)
	app.SetRoot(flex, true)
}

func SwitchToPotionToPokemonActionSelection(selector *pokemon_cli.PokemonUser, potionIndex int) {
	flex.Clear()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(BuildCliLogoItem(), 19, 1, false)
	bodyFlex := tview.NewFlex().SetDirection(tview.FlexColumn)
	bodyFlex.AddItem(BuildPokemonSelectionForPotionInCombatListItem(selector, potionIndex), 0, 1, true).
		AddItem(BuildActivePokemon(puser1), 0, 1, false).
		AddItem(BuildActivePokemon(puser2), 0, 1, false)
	flex.AddItem(bodyFlex, 0, 1, true)
	app.SetRoot(flex, true)
}
