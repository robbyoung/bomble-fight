package players

type IPlayerApplication interface {
}

type IPlayerStorage interface {
	LoadPlayer(string) *aggregate
	SavePlayer(*aggregate)
}
