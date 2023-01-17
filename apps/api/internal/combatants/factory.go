package combatants

import "bomble-fight/internal/common"

var r = common.NewRandom()
var storage = newLocalStorage(r)

func Api() *api {
	return newApi(newApplication(storage, r))
}

func TestApi() *api {
	return newApi(newApplication(newLocalStorage(r), r))
}

func Service() *service {
	return newService(newApplication(storage, r))
}

func TestService() *service {
	return newService(newApplication(newLocalStorage(r), r))
}
