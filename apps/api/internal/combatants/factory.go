package combatants

import "bomble-fight/internal/common"

var r = common.NewRandom()
var storage = NewLocalCombatantStorage(r)

func Api() *CombatantApi {
	return NewCombatantApi(NewCombatantApplication(storage, r))
}

func Service() *CombatantService {
	return NewCombatantService(NewCombatantApplication(storage, r))
}

// testing only
func clearStorage() {
	storage = NewLocalCombatantStorage(r)
}
