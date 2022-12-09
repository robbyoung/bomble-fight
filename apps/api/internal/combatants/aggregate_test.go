package combatants

import (
	"bomble-fight/internal/spec"
	"testing"
)

func TestNewAggregate(t *testing.T) {
	r := spec.NewMockRandom([]int{2, 3, 4})
	agg := newAggregate(r)

	if agg == nil {
		t.Fatalf("NewAggregate() returned nil")
	}

	if len(agg.Id) != 36 {
		t.Fatalf("Aggregate ID has unexpected value %s", agg.Id)
	}

	spec.ExpectEqualStrings(t, "Otto", agg.Name, "Unexpected name value")

	spec.ExpectEqualInts(t, 3, agg.Ferocity, "Unexpected ferocity value")
	spec.ExpectEqualInts(t, 4, agg.Endurance, "Unexpected endurance value")
	spec.ExpectEqualInts(t, 2, agg.Skill, "Unexpected skill value")
	spec.ExpectEqualInts(t, 3, agg.Agility, "Unexpected agility value")
	spec.ExpectEqualInts(t, 4, agg.Speed, "Unexpected speed value")
}

func TestInitiateWithAttack(t *testing.T) {
	r := spec.NewMockRandom([]int{1, 2, 2, 2, 2, 2, 5, 4})
	agg := newAggregate(r)

	response := agg.Initiate()

	spec.ExpectEqualInts(t, int(Attack), int(response.Code), "Unexpected action code")
	spec.ExpectEqualInts(t, 2, response.Detail, "Unexpected action detail")
}

func TestInitiateWithCriticalAttack(t *testing.T) {
	r := spec.NewMockRandom([]int{1, 3, 2, 2, 2, 2, 1, 4})
	agg := newAggregate(r)

	response := agg.Initiate()

	spec.ExpectEqualInts(t, int(Critical), int(response.Code), "Unexpected action code")
	spec.ExpectEqualInts(t, 6, response.Detail, "Unexpected action detail")
}

func TestRespondToAttackWithDodge(t *testing.T) {
	r := spec.NewMockRandom([]int{1, 2, 2, 2, 2, 2, 1, 8})
	agg := newAggregate(r)

	action := &Action{
		Code:   Attack,
		Detail: 10,
	}
	response := agg.Respond(action)

	spec.ExpectEqualInts(t, int(Dodge), int(response.Code), "Unexpected action code")
	spec.ExpectEqualInts(t, 50, agg.CurrentHealth, "Unexpected health value")
}

func TestRespondToAttackWithBlock(t *testing.T) {
	r := spec.NewMockRandom([]int{1, 2, 2, 2, 2, 2, 8, 1})
	agg := newAggregate(r)

	action := &Action{
		Code:   Attack,
		Detail: 10,
	}
	response := agg.Respond(action)

	spec.ExpectEqualInts(t, int(Block), int(response.Code), "Unexpected action code")
	spec.ExpectEqualInts(t, 50, agg.CurrentHealth, "Unexpected health value")
}

func TestRespondToAttackWithHit(t *testing.T) {
	r := spec.NewMockRandom([]int{1, 2, 2, 2, 2, 2, 8, 8})
	agg := newAggregate(r)

	action := &Action{
		Code:   Attack,
		Detail: 10,
	}
	response := agg.Respond(action)

	spec.ExpectEqualInts(t, int(Hit), int(response.Code), "Unexpected action code")
	spec.ExpectEqualInts(t, 40, agg.CurrentHealth, "Unexpected health value")
}

func TestRespondToCriticalAttackWithHit(t *testing.T) {
	r := spec.NewMockRandom([]int{1, 2, 2, 2, 2, 2, 8, 8})
	agg := newAggregate(r)

	action := &Action{
		Code:   Critical,
		Detail: 20,
	}
	response := agg.Respond(action)

	spec.ExpectEqualInts(t, int(Hit), int(response.Code), "Unexpected action code")
	spec.ExpectEqualInts(t, 30, agg.CurrentHealth, "Unexpected health value")
}

func TestRespondToAttackWithKilled(t *testing.T) {
	r := spec.NewMockRandom([]int{1, 2, 2, 2, 2, 2, 8, 8})
	agg := newAggregate(r)

	action := &Action{
		Code:   Attack,
		Detail: 50,
	}
	response := agg.Respond(action)

	spec.ExpectEqualInts(t, int(Killed), int(response.Code), "Unexpected action code")
	spec.ExpectEqualInts(t, 0, agg.CurrentHealth, "Unexpected health value")
}

func TestRespondToCriticalWithKilled(t *testing.T) {
	r := spec.NewMockRandom([]int{1, 2, 2, 2, 2, 2, 8, 8})
	agg := newAggregate(r)

	action := &Action{
		Code:   Critical,
		Detail: 100,
	}
	response := agg.Respond(action)

	spec.ExpectEqualInts(t, int(Killed), int(response.Code), "Unexpected action code")
	spec.ExpectEqualInts(t, 0, agg.CurrentHealth, "Unexpected health value")
}

func TestVictory(t *testing.T) {
	r := spec.NewMockRandom([]int{5})
	agg := newAggregate(r)

	agg.Streak = 1
	agg.CurrentHealth = 1

	agg.Victory()

	spec.ExpectEqualInts(t, 2, agg.Streak, "Unexpected streak value")
	spec.ExpectEqualInts(t, 50, agg.CurrentHealth, "Unexpected health value")
}
