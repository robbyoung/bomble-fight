package players

import (
	"bomble-fight/internal/spec"
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

func TestService_CreatePlayer(t *testing.T) {
	clearStorage()
	service := Service()

	p, err := service.application.CreatePlayer("aname")

	if err != nil {
		t.Fatalf("Unexpected error from CreatePlayer()")
	}

	spec.ExpectEqualStrings(t, "aname", p.Name, "Created player had unexpected name")
}

func TestService_CreatePlayer_EmptyName(t *testing.T) {
	clearStorage()
	service := Service()

	p, err := service.application.CreatePlayer("")

	if err == nil {
		t.Fatalf("Expected error from CreatePlayer(), received none")
	}

	if p != nil {
		t.Fatalf("Expected player object to be nil")
	}
}
