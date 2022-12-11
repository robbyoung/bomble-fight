package players

type application struct {
	storage IPlayerStorage
}

func newApplication(storage IPlayerStorage) *application {
	return &application{storage: storage}
}

func (app *application) CreatePlayer(name string) *Player {
	p := newAggregate(name)
	app.storage.SavePlayer(p)

	return convertToPlayerModel(p)
}

func convertToPlayerModel(p *aggregate) *Player {
	return p.toPersistence()
}
