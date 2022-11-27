package combatants

type service struct {
	application ICombatantApplication
}

func newService(app *application) *service {
	return &service{
		application: app,
	}
}

func (s *service) Fight(id1 string, id2 string) (*Action, *Action) {
	return s.application.Fight(id1, id2)
}
