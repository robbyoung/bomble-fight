package players

import (
	"errors"

	"github.com/google/uuid"
)

type aggregate struct {
	Id    string
	Name  string
	Money int
}

func newAggregate(name string) *aggregate {
	return &aggregate{
		Id:    uuid.New().String(),
		Name:  name,
		Money: 100,
	}
}

func (player *aggregate) SpendMoney(amount int) error {
	if amount < 0 {
		return errors.New("can't spend a negative amount of money")
	}

	if player.Money < amount {
		return errors.New("player doensn't have enough money to spend")
	}

	player.Money -= amount
	return nil
}

func (player *aggregate) EarnMoney(amount int) error {
	if amount < 0 {
		return errors.New("can't earn a negative amount of money")
	}

	player.Money += amount
	return nil
}

func fromPersistence(model *persistedModel) *aggregate {
	return &aggregate{
		Id:    model.Id,
		Name:  model.Name,
		Money: model.Money,
	}
}

func (player *aggregate) toPersistence() *persistedModel {
	return &persistedModel{
		Id:    player.Id,
		Name:  player.Name,
		Money: player.Money,
	}
}
