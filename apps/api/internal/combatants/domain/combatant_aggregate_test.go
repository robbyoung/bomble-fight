package combatants

import (
	"bomble-fight/internal/spec"
	"testing"
)

func TestNewCombatantAggregate(t *testing.T) {
	r := spec.NewMockRandom([]int{2, 3, 4})
	agg := NewCombatantAggregate(r)

	if agg == nil {
		t.Fatalf("NewCombatantAggregate() returned nil")
	}

	spec.ExpectEqualStrings(t, "Otto", agg.Name)

	spec.ExpectEqualInts(t, 4, agg.Ferocity)
	spec.ExpectEqualInts(t, 2, agg.Endurance)
	spec.ExpectEqualInts(t, 3, agg.Skill)
	spec.ExpectEqualInts(t, 4, agg.Agility)
	spec.ExpectEqualInts(t, 2, agg.Speed)
}

func TestInitiateWithAttack(t *testing.T) {
	r := spec.NewMockRandom([]int{1, 1, 2, 2, 2, 2, 2, 1, 4})
	agg := NewCombatantAggregate(r)

	response := agg.Initiate()

	spec.ExpectEqualInts(t, int(Attack), int(response.Code))
	spec.ExpectEqualInts(t, 4, response.Detail)
}

func TestInitiateWithCriticalAttack(t *testing.T) {
	r := spec.NewMockRandom([]int{1, 1, 2, 2, 2, 2, 2, 5, 4})
	agg := NewCombatantAggregate(r)

	response := agg.Initiate()

	spec.ExpectEqualInts(t, int(Critical), int(response.Code))
	spec.ExpectEqualInts(t, 8, response.Detail)
}

func TestRespondToAttackWithDodge(t *testing.T) {
	r := spec.NewMockRandom([]int{1, 1, 2, 2, 2, 2, 2, 1})
	agg := NewCombatantAggregate(r)

	action := CombatantAction{
		Code:   Attack,
		Detail: 10,
	}
	response := agg.Respond(action)

	spec.ExpectEqualInts(t, int(Dodge), int(response.Code))
	spec.ExpectEqualInts(t, 50, agg.Health)
}

func TestRespondToAttackWithHit(t *testing.T) {
	r := spec.NewMockRandom([]int{1, 1, 2, 2, 2, 2, 2, 8})
	agg := NewCombatantAggregate(r)

	action := CombatantAction{
		Code:   Attack,
		Detail: 10,
	}
	response := agg.Respond(action)

	spec.ExpectEqualInts(t, int(Hit), int(response.Code))
	spec.ExpectEqualInts(t, 40, agg.Health)
}
