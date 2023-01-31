package fights

import (
	"bomble-fight/internal/combatants"
	"bomble-fight/internal/players"
	"bomble-fight/internal/spec"
	"testing"
)

func TestApplication_New(t *testing.T) {
	storage := newLocalStorage()
	ps := players.TestService()
	cs := combatants.TestService()
	app := newApplication(storage, cs, ps)

	if app == nil {
		t.Fatalf("NewApplication() returned nil")
	}
}

func TestApplication_CreateFight(t *testing.T) {
	storage := newLocalStorage()
	ps := players.Service()
	cs := combatants.Service()
	app := newApplication(storage, cs, ps)

	c := cs.GenerateCombatants(1)[0]
	p, _ := ps.CreatePlayer("aname")

	f := app.CreateFight([]string{p.Id}, []string{c.Id})

	if f == nil {
		t.Fatal("CreateFight() returned nil unexpectedly")
	}

	if app.storage.LoadFight(f.Id) == nil {
		t.Fatal("No fight was saved to storage")
	}
}

func TestApplication_AddBet(t *testing.T) {
	storage := newLocalStorage()
	ps := players.Service()
	cs := combatants.Service()
	app := newApplication(storage, cs, ps)

	c := cs.GenerateCombatants(1)[0]
	p, _ := ps.CreatePlayer("aname")

	f := app.CreateFight([]string{p.Id}, []string{c.Id})
	err := app.AddBet(f.Id, p.Id, c.Id, 50)

	if err != nil {
		t.Fatal("AddBet() returned an error unexpectedly")
	}
	spec.ExpectEqualInts(t, 1, len(storage.fights[f.Id].Bets), "Unexpected bet count")
	spec.ExpectEqualInts(t, int(Ready), int(storage.fights[f.Id].Status), "Unexpected fight status")
}

func TestApplication_AddBet_AmountTooHigh(t *testing.T) {
	storage := newLocalStorage()
	ps := players.Service()
	cs := combatants.Service()
	app := newApplication(storage, cs, ps)

	c := cs.GenerateCombatants(1)[0]
	p, _ := ps.CreatePlayer("aname")

	f := app.CreateFight([]string{p.Id}, []string{c.Id})
	err := app.AddBet(f.Id, p.Id, c.Id, 110)

	if err == nil {
		t.Fatal("Expected an error from AddBet() but received none")
	}
	spec.ExpectEqualInts(t, 0, len(storage.fights[f.Id].Bets), "Unexpected bet count")
}

func TestApplication_AddBet_BadPid(t *testing.T) {
	storage := newLocalStorage()
	ps := players.Service()
	cs := combatants.Service()
	app := newApplication(storage, cs, ps)

	c := cs.GenerateCombatants(1)[0]
	p, _ := ps.CreatePlayer("aname")

	f := app.CreateFight([]string{}, []string{c.Id})
	err := app.AddBet(f.Id, p.Id, c.Id, 50)

	if err == nil {
		t.Fatal("Expected an error from AddBet() but received none")
	}
	spec.ExpectEqualInts(t, 0, len(storage.fights[f.Id].Bets), "Unexpected bet count")
}

func TestApplication_AddBet_BadCid(t *testing.T) {
	storage := newLocalStorage()
	ps := players.Service()
	cs := combatants.Service()
	app := newApplication(storage, combatants.Service(), players.Service())

	c := cs.GenerateCombatants(1)[0]
	p, _ := ps.CreatePlayer("aname")

	f := app.CreateFight([]string{p.Id}, []string{})
	err := app.AddBet(f.Id, p.Id, c.Id, 50)

	if err == nil {
		t.Fatal("Expected an error from AddBet() but received none")
	}
	spec.ExpectEqualInts(t, 0, len(storage.fights[f.Id].Bets), "Unexpected bet count")
}

func TestApplication_AddBet_DoubleUp(t *testing.T) {
	storage := newLocalStorage()
	ps := players.Service()
	cs := combatants.Service()
	app := newApplication(storage, cs, ps)

	c := cs.GenerateCombatants(1)[0]
	p, _ := ps.CreatePlayer("aname")

	f := app.CreateFight([]string{p.Id}, []string{c.Id})
	app.AddBet(f.Id, p.Id, c.Id, 50)
	err := app.AddBet(f.Id, p.Id, c.Id, 50)

	if err == nil {
		t.Fatal("Expected an error from AddBet() but received none")
	}
	spec.ExpectEqualInts(t, 1, len(storage.fights[f.Id].Bets), "Unexpected bet count")
}

func TestApplication_ResolveFight(t *testing.T) {
	storage := newLocalStorage()
	ps := players.Service()
	cs := combatants.Service()
	app := newApplication(storage, cs, ps)

	c := cs.GenerateCombatants(2)
	p, _ := ps.CreatePlayer("aname")

	f := app.CreateFight([]string{p.Id}, []string{c[0].Id, c[1].Id})
	app.AddBet(f.Id, p.Id, c[0].Id, 50)

	actions, err := app.ResolveFight(f.Id)

	if err != nil {
		t.Fatalf("Unexpected error from ResolveFight(): %s", err)
	}

	lastCode := actions[len(actions)-1].Code
	secondLastCode := actions[len(actions)-2].Code

	if lastCode != combatants.Killed && secondLastCode != combatants.Killed {
		t.Fatal("Neither of the last fight codes are 'Killed'")
	}
	spec.ExpectEqualInts(t, int(Finished), int(storage.fights[f.Id].Status), "Unexpected fight status")
}

func TestApplication_ResolveFight_NotReady(t *testing.T) {
	storage := newLocalStorage()
	ps := players.Service()
	cs := combatants.Service()
	app := newApplication(storage, cs, ps)

	c := cs.GenerateCombatants(2)
	p, _ := ps.CreatePlayer("aname")

	f := app.CreateFight([]string{p.Id}, []string{c[0].Id, c[1].Id})

	_, err := app.ResolveFight(f.Id)

	if err == nil {
		t.Fatal("Expected an error from ResolveFight() but received none")
	}

	spec.ExpectEqualInts(t, int(Pending), int(storage.fights[f.Id].Status), "Unexpected fight status")
}

func TestApplication_ResolveFight_AlreadyOver(t *testing.T) {
	storage := newLocalStorage()
	ps := players.Service()
	cs := combatants.Service()
	app := newApplication(storage, cs, ps)

	c := cs.GenerateCombatants(2)
	p, _ := ps.CreatePlayer("aname")

	f := app.CreateFight([]string{p.Id}, []string{c[0].Id, c[1].Id})
	app.AddBet(f.Id, p.Id, c[0].Id, 50)

	app.ResolveFight(f.Id)
	_, err := app.ResolveFight(f.Id)

	if err == nil {
		t.Fatal("Expected an error from ResolveFight() but received none")
	}

	spec.ExpectEqualInts(t, int(Finished), int(storage.fights[f.Id].Status), "Unexpected fight status")
}

func TestApplication_PayoutFans(t *testing.T) {
	storage := newLocalStorage()
	ps := players.Service()
	cs := combatants.Service()
	app := newApplication(storage, cs, ps)

	c := cs.GenerateCombatants(1)[0]
	p, _ := ps.CreatePlayer("aname")

	f := app.CreateFight([]string{p.Id}, []string{c.Id})
	app.AddBet(f.Id, p.Id, c.Id, 50)

	err := app.PayoutFans(f.Id, c.Id)
	p = ps.GetPlayer(p.Id)

	if err != nil {
		t.Fatal("Unexpected error from PayoutFans()")
	}
	spec.ExpectEqualInts(t, 150, p.Money, "Unexpected money amount")
}
