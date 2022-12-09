package players

import (
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
