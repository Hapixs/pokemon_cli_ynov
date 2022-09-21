package pokemon_cli

func GetFireEffect(roundId int) Effect {
	return buildEffect(roundId, "brûlure", 3, 10)
}

func GetPoisonEffect(roundId int) Effect {
	return buildEffect(roundId, "poison", 2, 15)
}

func GetGlaceEffect(roundId int) Effect {
	return buildEffect(roundId, "glace", 1, 30)
}

func GetElectrikEffect(roundId int) Effect {
	return buildEffect(roundId, "electrique", 4, 13)
}

func GetPsyEffect(roundId int) Effect {
	return buildEffect(roundId, "psy", 2, 8)
}

func buildEffect(roundId int, name string, duration, chance int) Effect {
	return Effect{name, duration, roundId, chance}
}
