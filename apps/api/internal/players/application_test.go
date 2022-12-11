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

	p := app.CreatePlayer("aname")

	if storage.LoadPlayer(p.Id) == nil {
		t.Fatalf("CreatePlayer() didn't save player")
	}

	spec.ExpectEqualStrings(t, p.Name, "aname", "New player has unexpected name")
}
