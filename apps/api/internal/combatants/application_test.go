package combatants

import (
	"bomble-fight/internal/common"
	"bomble-fight/internal/spec"
	"testing"
)

func TestApplication_New(t *testing.T) {
	r := spec.NewMockRandom([]int{1})
	storage := newLocalStorage(r)
	app := newApplication(storage, r)

	if app == nil {
		t.Fatalf("NewApplication() returned nil")
	}
}

func TestApplication_GenerateCombatants(t *testing.T) {
	r := common.NewRandom()
	storage := newLocalStorage(r)

	storage.SaveCombatant(newAggregate(r))
	storage.SaveCombatant(newAggregate(r))

	app := newApplication(storage, r)

	result := app.GenerateCombatants(3)

	if len(result) != 3 {
		t.Errorf("Expected 3 combatants to be generated, got %d", len(result))
	}

	if result[0] == nil || result[1] == nil || result[2] == nil {
		t.Errorf("One or more of the generated combatants was nil")
	}

	spec.ExpectEqualInts(t, 5, len(storage.combatants), "Unexpected stored combatants count")
}

func TestApplication_Fight(t *testing.T) {
	r := spec.NewMockRandom([]int{4, 4})
	storage := newLocalStorage(r)

	c1 := newAggregate(spec.NewMockRandom([]int{5}))
	c2 := newAggregate(spec.NewMockRandom([]int{3}))
	storage.SaveCombatant(c1)
	storage.SaveCombatant(c2)

	app := newApplication(storage, r)
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

func TestApplication_GetCombatant(t *testing.T) {
	r := spec.NewMockRandom([]int{4})
	storage := newLocalStorage(r)

	app := newApplication(storage, r)

	original := app.GenerateCombatants(1)[0]
	loaded := app.GetCombatant(original.Id)

	if loaded == nil {
		t.Error("GetCombatant returned nil unexpectedly")
		return
	}

	if original == loaded {
		t.Error("Original and loaded combatant have the same reference")
	}

	spec.ExpectEqualStrings(t, original.Name, loaded.Name, "Loaded combatant has a different name")
}

func TestApplication_GetCombatant_NotFound(t *testing.T) {
	r := spec.NewMockRandom([]int{4})
	storage := newLocalStorage(r)

	app := newApplication(storage, r)

	loaded := app.GetCombatant("aninvalidid")

	if loaded != nil {
		t.Error("Expected nil from GetCombatant()")
	}
}
