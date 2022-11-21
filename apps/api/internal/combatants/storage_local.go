package combatants

import "bomble-fight/internal/common"

type LocalCombatantStorage struct {
	combatants map[string]*CombatantPersistedModel
	r          common.IRandom
}

func NewLocalCombatantStorage(r common.IRandom) *LocalCombatantStorage {
	return &LocalCombatantStorage{
		combatants: make(map[string]*CombatantPersistedModel),
	}
}

func (storage *LocalCombatantStorage) LoadCombatant(id string) *CombatantAggregate {
	return FromPersistence(storage.combatants[id], storage.r)
}

func (storage *LocalCombatantStorage) SaveCombatant(combatant *CombatantAggregate) {
	storage.combatants[combatant.Id] = combatant.ToPersistence()
}
