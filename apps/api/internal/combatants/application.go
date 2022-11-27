package combatants

import "bomble-fight/internal/common"

type application struct {
	storage ICombatantStorage
	rand    common.IRandom
}

func newApplication(storage ICombatantStorage, r common.IRandom) *application {
	return &application{storage: storage, rand: r}
}

func (app *application) GenerateCombatants(count int) []*Combatant {
	results := make([]*Combatant, count)

	for i := 0; i < count; i++ {
		c := newAggregate(app.rand)
		app.storage.SaveCombatant(c)
		results[i] = convertToCombatantModel(c)
	}

	return results
}

func (app *application) Fight(id1 string, id2 string) (*Action, *Action) {
	c1 := app.storage.LoadCombatant(id1)
	c2 := app.storage.LoadCombatant(id2)

	a1 := c1.Initiate()
	a2 := c2.Respond(a1)

	app.storage.SaveCombatant(c1)
	app.storage.SaveCombatant(c2)

	return a1, a2
}

func convertToCombatantModel(c *aggregate) *Combatant {
	return c.ToPersistence()
}
