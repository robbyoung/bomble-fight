package fights

type localStorage struct {
	fights map[string]*persistedModel
}

func newLocalStorage() *localStorage {
	return &localStorage{
		fights: make(map[string]*persistedModel),
	}
}

func (storage *localStorage) LoadFight(id string) *aggregate {
	return fromPersistence(storage.fights[id])
}

func (storage *localStorage) SaveFight(fight *aggregate) {
	storage.fights[fight.id] = fight.toPersistence()
}
