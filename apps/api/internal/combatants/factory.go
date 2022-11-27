package combatants

import "bomble-fight/internal/common"

var r = common.NewRandom()
var storage = newLocalStorage(r)

func Api() *api {
	return newApi(newApplication(storage, r))
}

func Service() *service {
	return newService(newApplication(storage, r))
}

// testing only
func clearStorage() {
	storage = newLocalStorage(r)
}
