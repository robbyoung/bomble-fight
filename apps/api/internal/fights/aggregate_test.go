package fights

import (
	"bomble-fight/internal/spec"
	"testing"
)

func TestAggregate_New(t *testing.T) {
	players := []string{"aplayerid", "anotherplayerid"}
	combatants := []string{"acombatantid", "anothercombatantid"}
	agg := newAggregate(players, combatants)

	if agg == nil {
		t.Fatalf("NewAggregate() returned nil")
	}

	spec.ExpectEqualInts(t, len(players), len(agg.players), "Unexpected number of players in Fight")
	spec.ExpectEqualInts(t, len(combatants), len(agg.combatantIds), "Unexpected number of combatants in Fight")
	spec.ExpectEqualInts(t, 0, len(agg.bets), "Unexpected number of starting bets in Fight")
}

func TestAggregate_AddBet(t *testing.T) {
	players := []string{"aplayerid", "anotherplayerid"}
	combatants := []string{"acombatantid", "anothercombatantid"}
	agg := newAggregate(players, combatants)

	err := agg.AddBet("aplayerid", "acombatantid", 10)
	bet := agg.bets["aplayerid"]

	if err != nil {
		t.Fatalf("Unexpected error from AddBet()")
	}

	spec.ExpectEqualStrings(t, "aplayerid", bet.PlayerId, "Unexpected player id in bet")
	spec.ExpectEqualStrings(t, "acombatantid", bet.CombatantId, "Unexpected combatant id in bet")
	spec.ExpectEqualInts(t, 10, bet.Amount, "Unexpected amount in bet")
}

func TestAggregate_AddBet_InvalidPlayerId(t *testing.T) {
	players := []string{"aplayerid", "anotherplayerid"}
	combatants := []string{"acombatantid", "anothercombatantid"}
	agg := newAggregate(players, combatants)

	err := agg.AddBet("invalidplayerid", "acombatantid", 10)
	bet := agg.bets["invalidplayerid"]

	if err == nil {
		t.Fatalf("Expected error from AddBet() but received none")
	}

	if bet != nil {
		t.Fatalf("Expected no bet to be created")
	}
}

func TestAggregate_AddBet_InvalidCombatantId(t *testing.T) {
	players := []string{"aplayerid", "anotherplayerid"}
	combatants := []string{"acombatantid", "anothercombatantid"}
	agg := newAggregate(players, combatants)

	err := agg.AddBet("aplayerid", "aninvalidcombatantid", 10)
	bet := agg.bets["aplayerid"]

	if err == nil {
		t.Fatalf("Expected error from AddBet() but received none")
	}

	if bet != nil {
		t.Fatalf("Expected no bet to be created")
	}
}

func TestAggregate_AddBet_InvalidAmount(t *testing.T) {
	players := []string{"aplayerid", "anotherplayerid"}
	combatants := []string{"acombatantid", "anothercombatantid"}
	agg := newAggregate(players, combatants)

	err := agg.AddBet("aplayerid", "acombatantid", 0)
	bet := agg.bets["aplayerid"]

	if err == nil {
		t.Fatalf("Expected error from AddBet() but received none")
	}

	if bet != nil {
		t.Fatalf("Expected no bet to be created")
	}
}

func TestAggregate_AddBet_BetAlreadyPlaced(t *testing.T) {
	players := []string{"aplayerid", "anotherplayerid"}
	combatants := []string{"acombatantid", "anothercombatantid"}
	agg := newAggregate(players, combatants)

	agg.AddBet("aplayerid", "acombatantid", 20)
	err := agg.AddBet("aplayerid", "anothercombatantid", 50)
	bet := agg.bets["aplayerid"]

	if err == nil {
		t.Fatalf("Expected error from AddBet() but received none")
	}

	spec.ExpectEqualStrings(t, "acombatantid", bet.CombatantId, "Unexpected combatant id in bet")
	spec.ExpectEqualInts(t, 20, bet.Amount, "Unexpected amount in bet")
}
