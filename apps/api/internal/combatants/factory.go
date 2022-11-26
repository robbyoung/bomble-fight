package combatants

import "bomble-fight/internal/common"

var r = common.NewRandom()
var storage = NewLocalCombatantStorage(r)

func createApplication() *CombatantApplication {
	return NewCombatantApplication(storage, r)
}

func clearStorage() {
	storage = NewLocalCombatantStorage(r)
}
