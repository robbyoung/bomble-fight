package players

type IPlayerApplication interface {
	CreatePlayer(string) (*Player, error)
}

type IPlayerStorage interface {
	LoadPlayer(string) *aggregate
	SavePlayer(*aggregate)
}

type createPlayerRequest struct {
	Name string
}
