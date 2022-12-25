package combatants

type ICombatantService interface {
	Fight(string, string) (*Action, *Action)
}

type ICombatantApplication interface {
	GenerateCombatants(int) []*Combatant
	Fight(string, string) (*Action, *Action)
}

type ICombatantStorage interface {
	LoadCombatant(string) *aggregate
	SaveCombatant(*aggregate)
}

type generateCombatantsRequest struct {
	Count int
}
