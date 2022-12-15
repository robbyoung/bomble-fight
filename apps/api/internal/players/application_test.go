package players

import (
	"bomble-fight/internal/spec"
	"testing"
)

func TestNewApplication(t *testing.T) {
	storage := newLocalStorage()
	app := newApplication(storage)

	if app == nil {
		t.Fatalf("newApplication() returned nil")
	}
}

func TestCreatePlayer(t *testing.T) {
	storage := newLocalStorage()
	app := newApplication(storage)

	p, err := app.CreatePlayer("aname")

	if err != nil {
		t.Fatalf("CreatePlayer() returned an unexpected error")
	}

	if storage.LoadPlayer(p.Id) == nil {
		t.Fatalf("CreatePlayer() didn't save player")
	}

	spec.ExpectEqualStrings(t, p.Name, "aname", "New player has unexpected name")
}

func TestCreatePlayerWithEmptyName(t *testing.T) {
	storage := newLocalStorage()
	app := newApplication(storage)

	p, err := app.CreatePlayer("")

	if err == nil {
		t.Fatalf("Expected an error but didn't get one")
	}

	if p != nil {
		t.Fatalf("CreatePlayer() retured an unexpected player")
	}
}
