package players

type IPlayerApplication interface {
}

type IPlayerStorage interface {
	LoadCombatant(string) *aggregate
	SaveCombatant(*aggregate)
}
