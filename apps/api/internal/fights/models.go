package fights

type FightStatusCode int

const (
	Pending  FightStatusCode = 0
	Starting FightStatusCode = 1
	Active   FightStatusCode = 2
	Finished FightStatusCode = 3
)
