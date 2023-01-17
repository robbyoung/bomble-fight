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
	c, ok := storage.combatants[id]

	if ok {
		return fromPersistence(c, storage.r)
	}

	return nil
}

func (storage *localStorage) SaveCombatant(combatant *aggregate) {
	storage.combatants[combatant.Id] = combatant.ToPersistence()
}
