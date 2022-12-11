package players

type IPlayerApplication interface {
	CreatePlayer(string) *Player
}

type IPlayerStorage interface {
	LoadPlayer(string) *aggregate
	SavePlayer(*aggregate)
}

type createPlayerRequest struct {
	Name string
}
