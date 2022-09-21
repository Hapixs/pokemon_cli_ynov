package pokemon_cli

type AttackPokemonContext struct {
	Attacker  *Pokemon
	Reciever  *Pokemon
	Damages   int
	BadEffect []Effect
}

type PotionEffectContext struct {
	Reciever    Pokemon
	Potion      Potion
	worked      bool
	infoMessage string
}
