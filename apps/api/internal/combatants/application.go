package combatants

type CombatantApplication struct {
	storage ICombatantStorage
}

func NewCombatantApplication(storage ICombatantStorage) *CombatantApplication {
	return &CombatantApplication{storage: storage}
}

func (app *CombatantApplication) Fight(id1 string, id2 string) (*CombatantAction, *CombatantAction) {
	c1 := app.storage.LoadCombatant(id1)
	c2 := app.storage.LoadCombatant(id2)

	a1 := c1.Initiate()
	a2 := c2.Respond(a1)

	return a1, a2
}
