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
var pokemons = CastMapToSliceOf()

var pokemonsSelectedViewList = tview.NewList()
var pokemonsSelected = []pokemon_cli.Pokemon{}

var pokemonsTrimedViewList = tview.NewList()
var pokemonsTrimed = []pokemon_cli.Pokemon{}

var pokeText = tview.NewTextView()
var helpText = tview.NewTextView().SetText("(q) pour quitter \n(a) pour passer sur la liste de selection\n(e) pour passer sur la liste séléctionée")
var pokemonAsciiArt = tview.NewTextView()

var isSearching = false
var trimSearch = ""

func PreparePokemonsViewList() {
	pokemonsViewList.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		pokemonsSelected = append(pokemonsSelected, pokemons[i])
		UpdatePokeList(pokemonsSelected, pokemonsSelectedViewList)
		pokemons = removeIndex(i, pokemons)
		UpdatePokeList(pokemons, pokemonsViewList)
		UpdateViewListIndex(i, *pokemonsViewList)
		if len(pokemonsSelected) == 5 {
			rdm := getRandomPokemonList()
			startCombat(pokemon_cli.PokemonUser{0, pokemonsSelected, GeneratePotionInventory(), false}, pokemon_cli.PokemonUser{0, rdm, GeneratePotionInventory(), true})
		}
	})
	pokemonsViewList.SetChangedFunc(func(index int, name string, second_name string, shortcut rune) {
		UpdatePokeText(pokemons[index])
	})
}

func PreparePokemonsSelectedViewList() {
	pokemonsSelectedViewList.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		pokemons = append(pokemons, pokemonsSelected[i])
		UpdatePokeList(pokemons, pokemonsViewList)
		pokemonsSelected = removeIndex(i, pokemonsSelected)
		UpdatePokeList(pokemonsSelected, pokemonsSelectedViewList)
		if len(pokemonsSelected) <= 0 {
			switchToSelectionPokemons()
			return
		}
		UpdateViewListIndex(i, *pokemonsSelectedViewList)
	})
	pokemonsSelectedViewList.SetChangedFunc(func(index int, name string, second_name string, shortcut rune) {
		UpdatePokeText(pokemonsSelected[index])
	})
}

func PrepareTrimedViewList() {
	pokemonsTrimedViewList.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		pokemonsSelected = append(pokemonsSelected, pokemonsTrimed[i])
		UpdatePokeList(pokemonsSelected, pokemonsSelectedViewList)
		for index, p := range pokemons {
			if p.Name == pokemonsTrimed[i].Name {
				pokemons = removeIndex(index, pokemons)
				UpdatePokeList(pokemons, pokemonsViewList)
				break
			}
			if len(pokemonsSelected) == 5 {
				rdm := getRandomPokemonList()
				startCombat(pokemon_cli.PokemonUser{0, pokemonsSelected, GeneratePotionInventory(), false}, pokemon_cli.PokemonUser{0, rdm, GeneratePotionInventory(), true})
			}
		}
		pokemonsTrimed = removeIndex(i, pokemonsTrimed)
		UpdatePokeList(pokemonsTrimed, pokemonsTrimedViewList)
		UpdateViewListIndex(i, *pokemonsTrimedViewList)
	})
	pokemonsTrimedViewList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		UpdatePokeText(pokemonsTrimed[index])
	})
}

func UpdatePokeText(pok pokemon_cli.Pokemon) {
	pokeText.Clear()
	pokeText = pokeText.SetText(GetPokeText(pok))
	pokemonAsciiArt.Clear()
	pokemonAsciiArt.SetText(getPokemonImage(pok))
}

func UpdatePokeList(list []pokemon_cli.Pokemon, tList *tview.List) {
	tList.Clear()
	for _, pok := range list {
		tList.AddItem(pok.Name, "", '*', nil)
	}
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

func BuildActionList() *tview.List {
	return tview.NewList().AddItem("Attaquer", "", '*', func() {
		userAttackUser(&puser1, &puser2)
	}).AddItem("Changer de pokemon", "", '*', func() {
		userWantChangePokemon(&puser1)
	}).AddItem("Utiliser un objets", "", '*', func() {
		userWantUsePotion(&puser1)
	})
}

func BuildCliLogoItem() *tview.TextArea {
	return tview.NewTextArea().SetText(GetPokeCliLogo()+"\n", false)
}

func BuildPokemonInfoFlexItem() *tview.Flex {
	return tview.NewFlex().AddItem(pokemonAsciiArt, 0, 1, false).AddItem(pokeText, 0, 1, false)
}

func BuildSelectionViewListItem(active bool) *tview.Flex {
	return tview.NewFlex().AddItem(pokemonsViewList.SetSelectedBackgroundColor(GetListBackgroundColor(active)), 0, 1, active)
}

func BuildSelectedViewListItem(active bool) *tview.Flex {
	return tview.NewFlex().AddItem(pokemonsSelectedViewList.SetSelectedBackgroundColor(GetListBackgroundColor(active)), 0, 1, active)
}

func BuildHelpText() *tview.TextView {
	return helpText
}

func BuildActivePokemon(user pokemon_cli.PokemonUser) *tview.Flex {
	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText(getPokemonImage(user.PokemonInventory[user.ActivePokemon])), 0, 1, false).
		AddItem(tview.NewTextView().SetText(GetPokeText(user.PokemonInventory[user.ActivePokemon])), 0, 1, false)
}

func BuildPokemonSelectionInCombatListItem(user *pokemon_cli.PokemonUser) *tview.List {
	tl := tview.NewList()
	for _, c := range user.PokemonInventory {
		tl.AddItem(c.Name, "", '*', nil)
	}
	tl.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		userChangePokemon(user, i, true)
	})
	return tl
}

func BuildPokemonSelectionAfterCombatListItem(user *pokemon_cli.PokemonUser) *tview.List {
	tl := tview.NewList()
	for _, c := range user.PokemonInventory {
		tl.AddItem(c.Name, "", '*', nil)
	}
	tl.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		userChangePokemon(user, i, false)
	})
	return tl
}

func BuildPotionSelectionInCombatListItem(user *pokemon_cli.PokemonUser) *tview.List {
	tl := tview.NewList()
	for _, c := range user.ObjectInventory {
		tl.AddItem(c.Name, "", '*', nil)
	}
	tl.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		userSelectedPotion(user, i)
	})
	return tl
}

func BuildPokemonSelectionForPotionInCombatListItem(user *pokemon_cli.PokemonUser, potionIndex int) *tview.List {
	tl := tview.NewList()
	for _, c := range user.PokemonInventory {
		tl.AddItem(c.Name, "", '*', nil)
	}
	tl.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		userUsePotionOnPokemon(user, potionIndex, i)
	})
	return tl
}

func GetListBackgroundColor(active bool) tcell.Color {
	return map[bool]tcell.Color{true: tcell.ColorOrange, false: tcell.ColorAntiqueWhite}[active]
}

func GetPokeText(pok pokemon_cli.Pokemon) string {
	return "\nName: " + pok.Name +
		"\nHp: " + strconv.Itoa(pok.Hp) + "/" + strconv.Itoa(pok.MaxHp) +
		"\nDégats: " + strconv.Itoa(pok.Dmg) +
		"\nTypes: " + parseTabe(pok.Type)
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


// func getReversePokemonImage(pokemon pokemon_cli.Pokemon) string {
// 	content, err := ioutil.ReadFile("assets/" + pokemon.Name + ".txt")
// 	if err != nil {
// 		return ""
// 	}
// 	// reverse each line
// 	lines := strings.Split(string(content), ")
// 	return string(content)
// }
