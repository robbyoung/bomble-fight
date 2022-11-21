package combatants

type ICombatantStorage interface {
	LoadCombatant(string) *CombatantAggregate
	SaveCombatant(combatant *CombatantAggregate) string
}
