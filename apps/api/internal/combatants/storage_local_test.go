package combatants

import (
	"bomble-fight/internal/common"
	"bomble-fight/internal/spec"
	"testing"
)

func TestNewLocalCombatantStorage(t *testing.T) {
	r := common.NewRandom()
	cs := NewLocalCombatantStorage(r)

	if cs == nil {
		t.Fatalf("NewLocalCombatantStorage() returned nil")
	}

	if len(cs.combatants) != 0 {
		t.Fatalf("NewLocalCombatantStorage() didn't return an empty store")
	}
}

func TestCombatantStorageSaveAndLoad(t *testing.T) {
	r := common.NewRandom()
	cs := NewLocalCombatantStorage(r)

	original := NewCombatantAggregate(r)
	cs.SaveCombatant(original)

	loaded := cs.LoadCombatant(original.Id)

	if original == loaded {
		t.Fatalf("Loaded combatant uses same pointer as original")
	}

	if loaded.rand == nil {
		t.Fatalf("Loaded combatant has a nil random generator")
	}

	spec.ExpectEqualStrings(t, original.Id, loaded.Id)
	spec.ExpectEqualStrings(t, original.Name, loaded.Name)
	spec.ExpectEqualInts(t, original.Health, loaded.Health)
	spec.ExpectEqualInts(t, original.Ferocity, loaded.Ferocity)
	spec.ExpectEqualInts(t, original.Endurance, loaded.Endurance)
	spec.ExpectEqualInts(t, original.Skill, loaded.Skill)
	spec.ExpectEqualInts(t, original.Agility, loaded.Agility)
	spec.ExpectEqualInts(t, original.Speed, loaded.Speed)
}
