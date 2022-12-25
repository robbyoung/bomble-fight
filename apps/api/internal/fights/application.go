package fights

import (
	c "bomble-fight/internal/combatants"
	p "bomble-fight/internal/players"
)

type application struct {
	combatants c.ICombatantService
	storage    IFightStorage
}

func newApplication(storage IFightStorage, cs c.ICombatantService, ps p.IPlayerApplication) *application {
	return &application{storage: storage}
}
