package fights

type FightStatusCode int

const (
	Pending  FightStatusCode = 0
	Starting FightStatusCode = 1
	Active   FightStatusCode = 2
	Finished FightStatusCode = 3
)

type persistedModel struct {
	Id           string
	Bets         map[string]*bet
	Players      map[string]bool
	CombatantIds []string
	Status       FightStatusCode
}

type Fight = persistedModel
