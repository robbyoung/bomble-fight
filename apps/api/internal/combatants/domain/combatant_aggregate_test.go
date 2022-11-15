package combatants

import (
	"bomble-fight/internal/spec"
	"testing"
)

func TestNewCombatantAggregate(t *testing.T) {
	r := spec.NewMockRandom(2)
	agg := NewCombatantAggregate(r)

	if agg == nil {
		t.Fatalf("NewCombatantAggregate() returned nil")
	}

	spec.ExpectEqualStrings(t, "Otto", agg.Name)

	spec.ExpectEqualInts(t, 2, agg.Ferocity)
	spec.ExpectEqualInts(t, 2, agg.Endurance)
	spec.ExpectEqualInts(t, 2, agg.Agility)
	spec.ExpectEqualInts(t, 2, agg.Skill)
	spec.ExpectEqualInts(t, 2, agg.Speed)
}
