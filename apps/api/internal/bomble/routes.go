package bomble

import (
	"bomble-fight/internal/combatants"
	"bomble-fight/internal/common"
)

// Route is the model for the router setup
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc HandlerFunc
}

// Routes are the main setup for our Router
type Routes []Route

var random = common.NewRandom()

var combatantStorage = combatants.NewLocalCombatantStorage(random)
var combatantApplication = combatants.NewCombatantApplication(combatantStorage, random)
var combatantApi = combatants.NewCombatantApi(combatantApplication)

var routes = Routes{
	// Route{"Healthcheck", "GET", "/healthcheck", HealthcheckHandler},
	// Route{"GetUserState", "GET", "/state/{id:.+}", GetUserStateHandler},
	// Route{"AddPlayer", "POST", "/player", AddPlayerHandler},
	// Route{"AddBet", "POST", "/bet", AddBetHandler},
	// Route{"ListPlayers", "GET", "/players", ListPlayersHandler},
	// Route{"ListCombatants", "GET", "/combatants", ListCombatantsHandler},
	// Route{"GetFightStep", "POST", "/fight", FightStepHandler},
	// Route{"ResetFight", "POST", "/reset", ResetFightHandler},

	Route{"GenerateCombatants", "POST", "/combatants", combatantApi.GenerateCombatants},
}
