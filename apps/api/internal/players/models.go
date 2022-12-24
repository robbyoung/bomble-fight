package players

type Player = persistedModel

type persistedModel struct {
	Id    string
	Name  string
	Money int
}
