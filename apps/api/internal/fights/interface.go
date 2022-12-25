package fights

type IFightStorage interface {
	LoadPlayer(string) *aggregate
	SavePlayer(*aggregate)
}
