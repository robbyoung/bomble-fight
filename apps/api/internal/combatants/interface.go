package combatants

type ICombatantApplication interface {
	Fight(string, string) (*CombatantAction, *CombatantAction)
}

type ICombatantStorage interface {
	LoadCombatant(string) *CombatantAggregate
	SaveCombatant(combatant *CombatantAggregate)
}
