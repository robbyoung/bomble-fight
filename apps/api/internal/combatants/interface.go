package combatants

type ICombatantApplication interface {
	GenerateCombatants(int) []*Combatant
	Fight(string, string) (*CombatantAction, *CombatantAction)
}

type ICombatantStorage interface {
	LoadCombatant(string) *CombatantAggregate
	SaveCombatant(*CombatantAggregate)
}
