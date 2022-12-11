package players

var storage = newLocalStorage()

func Api() *api {
	return newApi(newApplication(storage))
}

// testing only
func clearStorage() {
	storage = newLocalStorage()
}
