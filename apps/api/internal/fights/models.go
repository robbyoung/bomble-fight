package fights

type FightStatusCode int

const (
	Pending  FightStatusCode = 0
	Ready    FightStatusCode = 1
	Finished FightStatusCode = 2
)

type persistedModel struct {
	Id           string
	Bets         map[string]*bet
	Players      map[string]bool
	CombatantIds []string
	Status       FightStatusCode
}

type Fight = persistedModel
