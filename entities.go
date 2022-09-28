package pokemon_cli

type PokemonUser struct {
	ActivePokemon    int
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
	Name        string
	Description string
	Heal        int
	DropRate    int
}

type Effect struct {
	Name          string
	RoundDuration int
	RoundStarted  int
	PerformChance int
}
