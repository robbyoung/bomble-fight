package players

type IPlayerService interface {
	CreatePlayer(string) (*Player, error)
	GetPlayer(string) *Player
	SpendMoney(string, int) error
}

type IPlayerApplication interface {
	CreatePlayer(string) (*Player, error)
	GetPlayer(string) *Player
	SpendMoney(string, int) error
}

type IPlayerStorage interface {
	LoadPlayer(string) *aggregate
	SavePlayer(*aggregate)
}

type createPlayerRequest struct {
	Name string
}
