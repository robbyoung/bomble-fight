package fights

import "errors"

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

func (agg *aggregate) AddBet(playerId string, combatantId string, amount int) error {
	if !agg.ContainsPlayer(playerId) {
		return errors.New("no record of specified player")
	}

	if !agg.ContainsCombatant(combatantId) {
		return errors.New("no record of specified combatant")
	}

	if amount <= 0 {
		return errors.New("bet amount must be greater than 0")
	}

	b := bet{
		PlayerId:    playerId,
		CombatantId: combatantId,
		Amount:      amount,
	}
	agg.bets = append(agg.bets, b)
	return nil
}

func (agg *aggregate) ContainsPlayer(id string) bool {
	for _, p := range agg.playerIds {
		if p == id {
			return true
		}
	}
	return false
}

func (agg *aggregate) ContainsCombatant(id string) bool {
	for _, c := range agg.combatantIds {
		if c == id {
			return true
		}
	}
	return false
}
