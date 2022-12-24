package fights

import (
	"bomble-fight/internal/spec"
	"testing"
)

func TestStorage_New(t *testing.T) {
	fs := newLocalStorage()

	if fs == nil {
		t.Fatalf("NewLocalStorage() returned nil")
	}

	if len(fs.fights) != 0 {
		t.Fatalf("NewLocalStorage() didn't return an empty store")
	}
}

func TestStorage_SaveAndLoad(t *testing.T) {
	players := []string{"aplayerid", "anotherplayerid"}
	combatants := []string{"acombatantid", "anothercombatantid"}

	original := newAggregate(players, combatants)
	original.AddBet("aplayerid", "acombatantid", 50)
	original.AddBet("anotherplayerid", "anothercombatantid", 50)

	fs := newLocalStorage()

	fs.SaveFight(original)

	loaded := fs.LoadFight(original.id)

	if original == loaded {
		t.Fatalf("Loaded combatant uses same pointer as original")
	}

	spec.ExpectEqualInts(t, 2, len(loaded.players), "Unexpected player count in loaded fight")
	spec.ExpectEqualInts(t, 2, len(loaded.combatantIds), "Unexpected combatant count in loaded fight")
	spec.ExpectEqualInts(t, 2, len(loaded.bets), "Unexpected bet count in loaded fight")
}
