package fights

import "errors"

type aggregate struct {
	bets         map[string]*bet
	players      map[string]bool
	combatantIds []string
	status       FightStatusCode
}

type bet struct {
	PlayerId    string
	CombatantId string
	Amount      int
}

func newAggregate(playerIds []string, combatantIds []string) *aggregate {
	players := make(map[string]bool)
	for _, id := range playerIds {
		players[id] = false
	}

	return &aggregate{
		bets:         make(map[string]*bet),
		players:      players,
		combatantIds: combatantIds,
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

	if agg.ContainsBet(playerId) {
		return errors.New("player already has placed a bet")
	}

	if amount <= 0 {
		return errors.New("bet amount must be greater than 0")
	}

	b := &bet{
		PlayerId:    playerId,
		CombatantId: combatantId,
		Amount:      amount,
	}
	agg.bets[playerId] = b
	return nil
}

func (agg *aggregate) ContainsBet(playerId string) bool {
	_, exists := agg.bets[playerId]
	return exists
}

func (agg *aggregate) ContainsPlayer(id string) bool {
	_, exists := agg.players[id]
	return exists
}

func (agg *aggregate) ContainsCombatant(id string) bool {
	for _, c := range agg.combatantIds {
		if c == id {
			return true
		}
	}
	return false
}
