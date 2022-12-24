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

	spec.ExpectEqualInts(t, len(players), len(agg.playerIds), "Unexpected number of players in Fight")
	spec.ExpectEqualInts(t, len(combatants), len(agg.combatantIds), "Unexpected number of combatants in Fight")
	spec.ExpectEqualInts(t, 0, len(agg.bets), "Unexpected number of starting bets in Fight")
}
