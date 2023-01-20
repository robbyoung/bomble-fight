package players

import (
	"bomble-fight/internal/spec"
	"testing"
)

func TestApplication_New(t *testing.T) {
	storage := newLocalStorage()
	app := newApplication(storage)

	if app == nil {
		t.Fatalf("newApplication() returned nil")
	}
}

func TestApplication_CreatePlayer(t *testing.T) {
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

func TestApplication_CreatePlayer_EmptyName(t *testing.T) {
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

func TestApplication_GetPlayer(t *testing.T) {
	storage := newLocalStorage()
	app := newApplication(storage)

	created, _ := app.CreatePlayer("aplayername")

	p := app.GetPlayer(created.Id)

	if p == nil {
		t.Fatalf("Unexpected nil from GetPlayer()")
	}

	spec.ExpectEqualStrings(t, "aplayername", p.Name, "Expected player names to match")
}

func TestApplication_GetPlayer_Nonexistent(t *testing.T) {
	storage := newLocalStorage()
	app := newApplication(storage)

	app.CreatePlayer("aplayername")

	p := app.GetPlayer("aninvalidguid")

	if p != nil {
		t.Fatalf("Unexpected player reference from GetPlayer()")
	}
}

func TestApplication_SpendMoney(t *testing.T) {
	storage := newLocalStorage()
	app := newApplication(storage)

	p, _ := app.CreatePlayer("aplayername")

	err := app.ChargePlayer(p.Id, 10)

	if err != nil {
		t.Fatal("Unexpected error from ChargePlayer()")
	}

	loaded := app.GetPlayer(p.Id)

	spec.ExpectEqualInts(t, 90, loaded.Money, "Unexpected money value in stored player")
}

func TestApplication_EarnMoney(t *testing.T) {
	storage := newLocalStorage()
	app := newApplication(storage)

	p, _ := app.CreatePlayer("aplayername")

	err := app.PayoutPlayer(p.Id, 10)

	if err != nil {
		t.Fatal("Unexpected error from PayoutPlayer()")
	}

	loaded := app.GetPlayer(p.Id)

	spec.ExpectEqualInts(t, 110, loaded.Money, "Unexpected money value in stored player")
}
