package players

type localStorage struct {
	players map[string]*persistedModel
}

func newLocalStorage() *localStorage {
	return &localStorage{
		players: make(map[string]*persistedModel),
	}
}

func (storage *localStorage) LoadPlayer(id string) *aggregate {
	p, ok := storage.players[id]

	if ok {
		return fromPersistence(p)
	}

	return nil
}

func (storage *localStorage) SavePlayer(player *aggregate) {
	storage.players[player.Id] = player.toPersistence()
}
