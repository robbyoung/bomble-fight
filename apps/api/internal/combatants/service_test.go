package combatants

import (
	"testing"
)

func TestService_New(t *testing.T) {
	clearStorage()
	service := Service()

	if service == nil {
		t.Fatalf("Service() returned nil")
	}

	if service.application == nil {
		t.Fatalf("Service object has no application")
	}
}

func TestService_Fight(t *testing.T) {
	clearStorage()
	service := Service()

	combatants := service.application.GenerateCombatants(2)

	a1, a2 := service.Fight(combatants[0].Id, combatants[1].Id)

	if a1 == nil || a2 == nil {
		t.Fatalf("One or both of the combatant actions was nil")
	}
}
