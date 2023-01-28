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
	spec.ExpectEqualInts(t, int(Pending), int(agg.status), "Unexpected fight status code")
}

func TestAggregate_AddBet_SetReady(t *testing.T) {
	players := []string{"aplayerid", "anotherplayerid"}
	combatants := []string{"acombatantid", "anothercombatantid"}
	agg := newAggregate(players, combatants)

	agg.AddBet("aplayerid", "acombatantid", 10)
	agg.AddBet("anotherplayerid", "acombatantid", 20)

	spec.ExpectEqualInts(t, int(Ready), int(agg.status), "Unexpected fight status code")
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

func TestAggregate_GetBet(t *testing.T) {
	players := []string{"aplayerid", "anotherplayerid"}
	combatants := []string{"acombatantid", "anothercombatantid"}
	agg := newAggregate(players, combatants)

	agg.AddBet("anotherplayerid", "anothercombatantid", 50)
	bet := agg.GetBet("anotherplayerid")

	spec.ExpectEqualStrings(t, "anotherplayerid", bet.PlayerId, "Unexpected player id in bet")
	spec.ExpectEqualStrings(t, "anothercombatantid", bet.CombatantId, "Unexpected combatant id in bet")
	spec.ExpectEqualInts(t, 50, bet.Amount, "Unexpected amount in bet")
}

func TestAggregate_GetBetsForCombatant(t *testing.T) {
	players := []string{"aplayerid", "anotherplayerid", "athirdplayerid"}
	combatants := []string{"acombatantid", "anothercombatantid"}
	agg := newAggregate(players, combatants)

	agg.AddBet("aplayerid", "anothercombatantid", 50)
	agg.AddBet("anotherplayerid", "acombatantid", 25)
	agg.AddBet("athirdplayerid", "acombatantid", 30)
	bets, err := agg.GetBetsForCombatant("acombatantid")

	if err != nil {
		t.Fatal("Unexpected error from GetBetsForCombatant()")
	}

	spec.ExpectEqualInts(t, 2, len(bets), "Unexpected length of bet array")
}

func TestAggregate_GetBetsForCombatant_Empty(t *testing.T) {
	players := []string{"aplayerid", "anotherplayerid", "athirdplayerid"}
	combatants := []string{"acombatantid", "anothercombatantid"}
	agg := newAggregate(players, combatants)

	agg.AddBet("aplayerid", "anothercombatantid", 50)
	bets, err := agg.GetBetsForCombatant("acombatantid")

	if err != nil {
		t.Fatal("Unexpected error from GetBetsForCombatant()")
	}

	spec.ExpectEqualInts(t, 0, len(bets), "Unexpected length of bet array")
}

func TestAggregate_GetBetsForCombatant_InvalidCombatant(t *testing.T) {
	players := []string{"aplayerid", "anotherplayerid", "athirdplayerid"}
	combatants := []string{"acombatantid", "anothercombatantid"}
	agg := newAggregate(players, combatants)

	bets, err := agg.GetBetsForCombatant("invalidcombatantid")

	if err == nil {
		t.Fatal("Expected error from GetBetsForCombatant(), received none")
	}

	if bets != nil {
		t.Fatal("Expected no bets array to be returned")
	}
}

func TestAggregate_GetCombatantIds(t *testing.T) {
	players := []string{"aplayerid", "anotherplayerid"}
	combatants := []string{"acombatantid", "anothercombatantid"}
	agg := newAggregate(players, combatants)

	c1, c2, err := agg.GetCombatantIds()

	if err != nil {
		t.Fatal("Received an unexpected error from GetCombatantIds()")
	}

	spec.ExpectEqualStrings(t, "acombatantid", c1, "Combatant 1 id mismatch")
	spec.ExpectEqualStrings(t, "anothercombatantid", c2, "Combatant 2 id mismatch")
}

func TestAggregate_GetCombatantIds_InvalidNumber(t *testing.T) {
	players := []string{"aplayerid", "anotherplayerid"}
	combatants := []string{"acombatantid", "anothercombatantid", "athirdcombatantid"}
	agg := newAggregate(players, combatants)

	_, _, err := agg.GetCombatantIds()

	if err == nil {
		t.Fatal("Expected an error from GetCombatantIds() but received none")
	}
}
