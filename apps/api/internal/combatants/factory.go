package combatants

import "bomble-fight/internal/common"

var r = common.NewRandom()
var storage = newLocalStorage(r)

func Api() ICombatantApi {
	return newApi(newApplication(storage, r))
}

func TestApi() ICombatantApi {
	return newApi(newApplication(newLocalStorage(r), r))
}

func Service() ICombatantService {
	return newApplication(storage, r)
}

func TestService() ICombatantService {
	return newApplication(newLocalStorage(r), r)
}
