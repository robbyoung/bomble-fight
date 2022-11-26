package combatants

type CombatantService struct {
	application ICombatantApplication
}

func NewCombatantService(app *CombatantApplication) *CombatantService {
	return &CombatantService{
		application: app,
	}
}

func (s *CombatantService) Fight(id1 string, id2 string) (*CombatantAction, *CombatantAction) {
	return s.application.Fight(id1, id2)
}
