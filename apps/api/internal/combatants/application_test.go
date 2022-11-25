package combatants

import (
	"bomble-fight/internal/common"
	"bomble-fight/internal/spec"
	"testing"
)

func TestNewCombatantApplication(t *testing.T) {
	r := spec.NewMockRandom([]int{1})
	storage := NewLocalCombatantStorage(r)
	app := NewCombatantApplication(storage, r)

	if app == nil {
		t.Fatalf("NewCombatantApplication() returned nil")
	}
}

func TestGenerateCombatants(t *testing.T) {
	r := common.NewRandom()
	storage := NewLocalCombatantStorage(r)

	storage.SaveCombatant(NewCombatantAggregate(r))
	storage.SaveCombatant(NewCombatantAggregate(r))

	app := NewCombatantApplication(storage, r)

	result := app.GenerateCombatants(3)

	if len(result) != 3 {
		t.Errorf("Expected 3 combatants to be generated, got %d", len(result))
	}

	if result[0] == nil || result[1] == nil || result[2] == nil {
		t.Errorf("One or more of the generated combatants was nil")
	}

	spec.ExpectEqualInts(t, 5, len(storage.combatants), "Unexpected stored combatants count")
}

func TestFight(t *testing.T) {
	r := spec.NewMockRandom([]int{4, 4})
	storage := NewLocalCombatantStorage(r)

	c1 := NewCombatantAggregate(spec.NewMockRandom([]int{5}))
	c2 := NewCombatantAggregate(spec.NewMockRandom([]int{3}))
	storage.SaveCombatant(c1)
	storage.SaveCombatant(c2)

	app := NewCombatantApplication(storage, r)
	a1, a2 := app.Fight(c1.Id, c2.Id)

	c2 = storage.LoadCombatant(c2.Id)

	if a1.Code != Critical {
		t.Errorf("Expected Critical from combatant 1, got %d", int(a1.Code))
	}

	if a2.Code != Hit {
		t.Errorf("Expected Hit from combatant 2, got %d", int(a2.Code))
	}

	spec.ExpectEqualInts(t, 40, c2.CurrentHealth, "Unexpected combatant 2 health")

}
