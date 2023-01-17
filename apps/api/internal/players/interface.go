package players

import (
	"bomble-fight/internal/common"
	"net/http"
)

type IPlayerApi interface {
	CreatePlayer(w http.ResponseWriter, req *http.Request, appEnv common.AppEnv)
}

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
