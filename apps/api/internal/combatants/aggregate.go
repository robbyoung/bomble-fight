package combatants

import (
	"bomble-fight/internal/common"

	"github.com/google/uuid"
)

type aggregate struct {
	rand common.IRandom

	Id            string
	Name          string
	MaxHealth     int
	CurrentHealth int
	Streak        int

	Ferocity  int
	Agility   int
	Endurance int
	Skill     int
	Speed     int
}

var names = []string{
	"Winslow",
	"Eberhardt",
	"Otto",
	"Alured",
	"Crawford",
	"Willowby",
	"Rollo",
	"Augustus",
	"Forster",
	"Guy",
}

func newAggregate(r common.IRandom) *aggregate {
	return &aggregate{
		rand: r,

		Id:            uuid.New().String(),
		Name:          r.RandArrayEntry(names),
		MaxHealth:     50,
		CurrentHealth: 50,
		Streak:        0,

		Ferocity:  r.RandInt(1, 11),
		Endurance: r.RandInt(1, 11),
		Skill:     r.RandInt(1, 11),
		Agility:   r.RandInt(1, 11),
		Speed:     r.RandInt(1, 11),
	}
}

func (combatant *aggregate) Initiate() *Action {
	damage := combatant.Ferocity

	if combatant.rollForAbility(combatant.Skill) {
		return &Action{
			Code:   Critical,
			Detail: damage * 2,
		}
	}

	return &Action{
		Code:   Attack,
		Detail: damage,
	}
}

func (combatant *aggregate) Respond(action *Action) *Action {
	switch action.Code {
	case Attack:
		return combatant.respondToAttack(action.Detail)
	case Critical:
		return combatant.respondToAttack(action.Detail)
	default:
		return &Action{
			Code: Idle,
		}
	}
}

func (combatant *aggregate) Victory() {
	combatant.CurrentHealth = combatant.MaxHealth
	combatant.Streak++
}

func (combatant *aggregate) respondToAttack(damage int) *Action {
	if combatant.rollForAbility(combatant.Agility) {
		return &Action{
			Code: Dodge,
		}
	}

	if combatant.rollForAbility(combatant.Endurance) {
		return &Action{
			Code: Block,
		}
	}

	combatant.CurrentHealth -= damage

	if combatant.CurrentHealth <= 0 {
		combatant.CurrentHealth = 0
		return &Action{
			Code: Killed,
		}
	}

	return &Action{
		Code: Hit,
	}
}

func (combatant *aggregate) rollForAbility(threshold int) bool {
	roll := combatant.rand.RandInt(0, 10)

	return roll <= threshold
}

func fromPersistence(model *persistedModel, r common.IRandom) *aggregate {
	return &aggregate{
		rand:          r,
		Id:            model.Id,
		Name:          model.Name,
		MaxHealth:     model.MaxHealth,
		CurrentHealth: model.CurrentHealth,
		Streak:        model.Streak,
		Ferocity:      model.Ferocity,
		Endurance:     model.Endurance,
		Skill:         model.Skill,
		Agility:       model.Agility,
		Speed:         model.Speed,
	}
}

func (combatant *aggregate) ToPersistence() *persistedModel {
	return &persistedModel{
		Id:            combatant.Id,
		Name:          combatant.Name,
		MaxHealth:     combatant.MaxHealth,
		CurrentHealth: combatant.CurrentHealth,
		Streak:        combatant.Streak,
		Ferocity:      combatant.Ferocity,
		Endurance:     combatant.Endurance,
		Skill:         combatant.Skill,
		Agility:       combatant.Agility,
		Speed:         combatant.Speed,
	}
}
