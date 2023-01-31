package fights

import (
	"errors"

	"github.com/google/uuid"
)

type aggregate struct {
	id           string
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
		id:           uuid.New().String(),
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
	agg.players[playerId] = true
	agg.readyCheck()

	return nil
}

func (agg *aggregate) GetBet(playerId string) bet {
	return *agg.bets[playerId]
}

func (agg *aggregate) GetCombatantIds() (string, string, error) {
	if len(agg.combatantIds) != 2 {
		return "", "", errors.New("fight contains an unexpected number of combatant ids")
	}

	return agg.combatantIds[0], agg.combatantIds[1], nil
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

func (agg *aggregate) GetBetsForCombatant(cid string) ([]bet, error) {
	if !agg.ContainsCombatant(cid) {
		return nil, errors.New("no record of specified combatant")
	}

	bets := make([]bet, 0)
	for _, b := range agg.bets {
		if b.CombatantId == cid {
			bets = append(bets, *b)
		}
	}

	return bets, nil
}

func (agg *aggregate) readyCheck() {
	for _, ready := range agg.players {
		if !ready {
			return
		}
	}

	agg.status = Ready
}

func (agg *aggregate) toPersistence() *persistedModel {
	return &persistedModel{
		Id:           agg.id,
		Bets:         agg.bets,
		Players:      agg.players,
		CombatantIds: agg.combatantIds,
		Status:       agg.status,
	}
}

func fromPersistence(model *persistedModel) *aggregate {
	return &aggregate{
		id:           model.Id,
		bets:         model.Bets,
		players:      model.Players,
		combatantIds: model.CombatantIds,
		status:       model.Status,
	}
}
