package combatants

type LocalCombatantStorage struct {
	combatants map[string]*CombatantAggregate
}

func CreateLocalCombatantStorage() *LocalCombatantStorage {
	return &LocalCombatantStorage{
		combatants: make(map[string]*CombatantAggregate),
	}
}

func (storage *LocalCombatantStorage) LoadCombatant(id string) *CombatantAggregate {
	return storage.combatants[id]
}

func (storage *LocalCombatantStorage) SaveCombatant(combatant *CombatantAggregate) {
	storage.combatants[combatant.Id] = combatant
}
