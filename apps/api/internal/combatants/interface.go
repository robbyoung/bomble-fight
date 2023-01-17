package combatants

import (
	"bomble-fight/internal/common"
	"net/http"
)

type ICombatantApi interface {
	GenerateCombatants(w http.ResponseWriter, req *http.Request, appEnv common.AppEnv)
}

type ICombatantService interface {
	GenerateCombatants(int) []*Combatant
	GetCombatant(string) *Combatant
	Fight(string, string) (*Action, *Action)
}

type ICombatantApplication interface {
	GenerateCombatants(int) []*Combatant
	GetCombatant(string) *Combatant
	Fight(string, string) (*Action, *Action)
}

type ICombatantStorage interface {
	LoadCombatant(string) *aggregate
	SaveCombatant(*aggregate)
}

type generateCombatantsRequest struct {
	Count int
}
