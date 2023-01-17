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

func (app *application) GetPlayer(id string) *Player {
	loaded := app.storage.LoadPlayer(id)

	if loaded != nil {
		return convertToPlayerModel(loaded)
	}

	return nil
}

func (app *application) SpendMoney(id string, amount int) error {
	player := app.storage.LoadPlayer(id)

	err := player.SpendMoney(amount)

	app.storage.SavePlayer(player)

	return err
}

func convertToPlayerModel(p *aggregate) *Player {
	return p.toPersistence()
}
