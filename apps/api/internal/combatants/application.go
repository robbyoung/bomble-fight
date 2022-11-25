package combatants

import "bomble-fight/internal/common"

type CombatantApplication struct {
	storage ICombatantStorage
	rand    common.IRandom
}

func NewCombatantApplication(storage ICombatantStorage, r common.IRandom) *CombatantApplication {
	return &CombatantApplication{storage: storage, rand: r}
}

func (app *CombatantApplication) GenerateCombatants(count int) []*Combatant {
	results := make([]*Combatant, count)

	for i := 0; i < count; i++ {
		c := NewCombatantAggregate(app.rand)
		app.storage.SaveCombatant(c)
		results[i] = ConvertToCombatantModel(c)
	}

	return results
}

func (app *CombatantApplication) Fight(id1 string, id2 string) (*CombatantAction, *CombatantAction) {
	c1 := app.storage.LoadCombatant(id1)
	c2 := app.storage.LoadCombatant(id2)

	a1 := c1.Initiate()
	a2 := c2.Respond(a1)

	app.storage.SaveCombatant(c1)
	app.storage.SaveCombatant(c2)

	return a1, a2
}

func ConvertToCombatantModel(c *CombatantAggregate) *Combatant {
	return c.ToPersistence()
}
