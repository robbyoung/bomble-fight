package combatants

type ICombatantApplication interface {
	GenerateCombatants(string) []*Combatant
	Fight(string, string) (*CombatantAction, *CombatantAction)
}

type ICombatantStorage interface {
	LoadCombatant(string) *CombatantAggregate
	SaveCombatant(*CombatantAggregate)
}
