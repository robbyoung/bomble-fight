package combatants

import "bomble-fight/internal/common"

type localStorage struct {
	combatants map[string]*persistedModel
	r          common.IRandom
}

func newLocalStorage(r common.IRandom) *localStorage {
	return &localStorage{
		combatants: make(map[string]*persistedModel),
		r:          r,
	}
}

func (storage *localStorage) LoadCombatant(id string) *aggregate {
	return fromPersistence(storage.combatants[id], storage.r)
}

func (storage *localStorage) SaveCombatant(combatant *aggregate) {
	storage.combatants[combatant.Id] = combatant.ToPersistence()
}
