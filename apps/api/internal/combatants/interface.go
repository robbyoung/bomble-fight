package combatants

type ICombatantService interface {
	GenerateCombatants(int) []*Combatant
	GetCombatant(string) *Combatant
	Fight(string, string) (*Action, *Action)
}

type ICombatantApplication interface {
	GenerateCombatants(int) []*Combatant
	GetCombatant(string) *Combatant
	Fight(string, string) (*Action, *Action)
}

type ICombatantStorage interface {
	LoadCombatant(string) *aggregate
	SaveCombatant(*aggregate)
}

type generateCombatantsRequest struct {
	Count int
}
