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
	return fromPersistence(storage.players[id])
}

func (storage *localStorage) SavePlayer(player *aggregate) {
	storage.players[player.Id] = player.toPersistence()
}
