package players

import "errors"

type application struct {
	storage IPlayerStorage
}

func newApplication(storage IPlayerStorage) *application {
	return &application{storage: storage}
}

func (app *application) CreatePlayer(name string) (*Player, error) {
	if len(name) == 0 {
		return nil, errors.New("Player name can't be empty")
	}

	p := newAggregate(name)
	app.storage.SavePlayer(p)

	return convertToPlayerModel(p), nil
}

func convertToPlayerModel(p *aggregate) *Player {
	return p.toPersistence()
}
