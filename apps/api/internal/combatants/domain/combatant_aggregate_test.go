package combatants

import (
	"testing"
)

func TestNewCombatantAggregate(t *testing.T) {
	agg := NewCombatantAggregate()

	if agg == nil {
		t.Fatalf("NewCombatantAggregate() returned nil")
	}

	if len(agg.Name) == 0 {
		t.Fatalf("NewCombatantAggregate() returned a combatant with an invalid name")
	}

	if agg.Ferocity < 1 || agg.Ferocity > 10 {
		t.Fatalf("NewCombatantAggregate() returned a combatant with an invalid Ferocity %d", agg.Ferocity)
	}

	if agg.Endurance < 1 || agg.Endurance > 10 {
		t.Fatalf("NewCombatantAggregate() returned a combatant with an invalid Endurance %d", agg.Endurance)
	}

	if agg.Skill < 1 || agg.Skill > 10 {
		t.Fatalf("NewCombatantAggregate() returned a combatant with an invalid Skill %d", agg.Skill)
	}

	if agg.Agility < 1 || agg.Agility > 10 {
		t.Fatalf("NewCombatantAggregate() returned a combatant with an invalid Agility %d", agg.Agility)
	}

	if agg.Speed < 1 || agg.Speed > 10 {
		t.Fatalf("NewCombatantAggregate() returned a combatant with an invalid Speed %d", agg.Speed)
	}
}
