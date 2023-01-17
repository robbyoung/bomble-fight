package players

var storage = newLocalStorage()

func Api() IPlayerApi {
	return newApi(newApplication(storage))
}

func TestApi() IPlayerApi {
	return newApi(newApplication(newLocalStorage()))
}

func Service() IPlayerService {
	return newApplication(storage)
}

func TestService() IPlayerService {
	return newApplication(newLocalStorage())
}
