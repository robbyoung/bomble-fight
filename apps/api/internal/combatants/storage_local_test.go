package combatants

import (
	"bomble-fight/internal/common"
	"bomble-fight/internal/spec"
	"testing"
)

func TestStorage_New(t *testing.T) {
	r := common.NewRandom()
	cs := newLocalStorage(r)

	if cs == nil {
		t.Fatalf("NewLocalCombatantStorage() returned nil")
	}

	if len(cs.combatants) != 0 {
		t.Fatalf("NewLocalCombatantStorage() didn't return an empty store")
	}
}

func TestStorage_SaveAndLoad(t *testing.T) {
	r := common.NewRandom()
	cs := newLocalStorage(r)

	original := newAggregate(r)
	original.Streak = 22

	cs.SaveCombatant(original)

	loaded := cs.LoadCombatant(original.Id)

	if original == loaded {
		t.Fatalf("Loaded combatant uses same pointer as original")
	}

	if loaded.rand == nil {
		t.Fatalf("Loaded combatant has a nil random generator")
	}

	spec.ExpectEqualStrings(t, original.Id, loaded.Id, "Id mismatch")
	spec.ExpectEqualStrings(t, original.Name, loaded.Name, "Name mismatch")
	spec.ExpectEqualInts(t, original.MaxHealth, loaded.MaxHealth, "MaxHealth mismatch")
	spec.ExpectEqualInts(t, original.MaxHealth, loaded.CurrentHealth, "CurrentHealth mismatch")
	spec.ExpectEqualInts(t, original.Ferocity, loaded.Ferocity, "Ferocity mismatch")
	spec.ExpectEqualInts(t, original.Endurance, loaded.Endurance, "Endurance mismatch")
	spec.ExpectEqualInts(t, original.Skill, loaded.Skill, "Skill mismatch")
	spec.ExpectEqualInts(t, original.Agility, loaded.Agility, "Agility mismatch")
	spec.ExpectEqualInts(t, original.Speed, loaded.Speed, "Speed mismatch")
	spec.ExpectEqualInts(t, original.Streak, loaded.Streak, "Speed mismatch")
}
