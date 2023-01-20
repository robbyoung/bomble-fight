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

	return nil
}

func convertToFightModel(a *aggregate) *Fight {
	return a.toPersistence()
}
