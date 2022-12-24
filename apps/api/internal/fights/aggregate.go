package fights

type aggregate struct {
	bets         []bet
	playerIds    []string
	combatantIds []string
	status       FightStatusCode
}

type bet struct {
	PlayerId    string
	CombatantId string
	Amount      int
}

func newAggregate(players []string, combatants []string) *aggregate {
	return &aggregate{
		bets:         make([]bet, 0),
		playerIds:    players,
		combatantIds: combatants,
		status:       Pending,
	}
}
