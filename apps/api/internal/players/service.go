package players

type service struct {
	application IPlayerApplication
}

func newService(app *application) *service {
	return &service{
		application: app,
	}
}

func (s *service) CreatePlayer(name string) (*Player, error) {
	return s.application.CreatePlayer(name)
}

func (s *service) GetPlayer(id string) *Player {
	return s.application.GetPlayer(id)
}
