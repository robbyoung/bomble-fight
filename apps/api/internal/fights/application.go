package fights

import (
	c "bomble-fight/internal/combatants"
	p "bomble-fight/internal/players"
	"errors"
)

type application struct {
	combatants c.ICombatantService
	players    p.IPlayerService
	storage    IFightStorage
}

func newApplication(storage IFightStorage, cs c.ICombatantService, ps p.IPlayerService) *application {
	return &application{storage: storage, players: ps, combatants: cs}
}

func (a *application) CreateFight(pids []string, cids []string) *Fight {
	agg := newAggregate(pids, cids)
	a.storage.SaveFight(agg)
	return convertToFightModel(agg)
}

func (a *application) AddBet(fid string, pid string, cid string, amount int) error {
	f := a.storage.LoadFight(fid)
	if f == nil {
		return errors.New("no fight found with specified ID")
	}

	p := a.players.GetPlayer(pid)
	if p == nil {
		return errors.New("no player found with specified ID")
	}

	if p.Money < amount {
		return errors.New("player doesn't have enough money")
	}

	err := f.AddBet(pid, cid, amount)
	if err != nil {
		return err
	}

	a.players.ChargePlayer(pid, amount)

	a.storage.SaveFight(f)

	return nil
}

func convertToFightModel(a *aggregate) *Fight {
	return a.toPersistence()
}

func (a *application) PayoutFans(fid string, cid string) error {
	f := a.storage.LoadFight(fid)
	if f == nil {
		return errors.New("no fight found with specified ID")
	}

	bets, err := f.GetBetsForCombatant(cid)
	if err != nil {
		return err
	}

	for _, b := range bets {
		a.players.PayoutPlayer(b.PlayerId, b.Amount*2)
	}

	a.storage.SaveFight(f)

	return nil
}

func (a *application) ResolveFight(fid string) ([]*c.Action, error) {
	f := a.storage.LoadFight(fid)
	if f == nil {
		return nil, errors.New("no fight found with specified ID")
	}

	if f.status == Pending {
		return nil, errors.New("fight isn't ready to start yet")
	}

	if f.status == Finished {
		return nil, errors.New("fight is already over")
	}

	c1, c2, _ := f.GetCombatantIds()
	actions := make([]*c.Action, 0)
	for {
		a1, a2 := a.combatants.Fight(c1, c2)
		actions = append(actions, a1, a2)

		if len(actions) > 1000 {
			return nil, errors.New("fight didn't resolve after 1000 steps")
		}

		if a1.Code == c.Killed || a2.Code == c.Killed {
			f.status = Finished
			a.storage.SaveFight(f)
			return actions, nil
		}
	}
}
