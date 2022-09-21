package pokemon_cli

type PokemonUser struct {
	ActivePokemon    *Pokemon
	PokemonInventory []Pokemon
	ObjectInventory  []Potion
	Ai               bool
}

type Pokemon struct {
	Name          string
	Type          []string
	Hp            int
	MaxHp         int
	Dmg           int
	ActiveEffects []Effect
	AttackEffects []Effect
}

type Potion struct {
	Name     string
	Heal     int
	DropRate int
}

type Effect struct {
	Name          string
	RoundDuration int
	RoundStarted  int
	PerformChance int
}
