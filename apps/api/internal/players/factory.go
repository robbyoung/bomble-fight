package players

var storage = newLocalStorage()

func Api() *api {
	return newApi(newApplication(storage))
}

func TestApi() *api {
	return newApi(newApplication(newLocalStorage()))
}

func Service() *service {
	return newService(newApplication(storage))
}

func TestService() *service {
	return newService(newApplication(newLocalStorage()))
}
