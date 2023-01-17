package fights

type IFightStorage interface {
	LoadFight(string) *aggregate
	SaveFight(*aggregate)
}
